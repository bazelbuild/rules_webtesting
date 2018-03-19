// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"

	"github.com/bazelbuild/rules_webtesting/go/cmdhelper"
	"github.com/bazelbuild/rules_webtesting/go/errors"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
)

// Cmd is a service that starts an external executable.
type Cmd struct {
	*Base
	cmd *exec.Cmd

	mu             sync.RWMutex
	stopMonitoring bool

	done chan interface{} // this channel is closed when the process stops.
}

// NewCmd creates a new service for starting an external server on the host machine.
func NewCmd(name string, d diagnostics.Diagnostics, exe string, xvfb bool, env map[string]string, args ...string) (*Cmd, error) {
	if xvfb {
		args = append([]string{"-a", exe}, args...)
		exe = "/usr/bin/xvfb-run"
	}
	cmd := exec.Command(exe, args...)
	if env != nil {
		cmd.Env = cmdhelper.BulkUpdateEnv(os.Environ(), env)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return &Cmd{
		Base: NewBase(name, d),
		cmd:  cmd,
		done: make(chan interface{}),
	}, nil
}

// Start starts the executable, waits for it to become healhy, and monitors it to ensure that it
// stays healthy.
func (c *Cmd) Start(ctx context.Context) error {
	if err := c.Base.Start(ctx); err != nil {
		return err
	}

	if err := c.cmd.Start(); err != nil {
		return errors.New(c.Name(), err)
	}

	go c.Monitor()
	return nil
}

// Stop stops the executable.
func (c *Cmd) Stop(ctx context.Context) error {
	if err := c.Base.Stop(ctx); err != nil {
		return err
	}
	c.StopMonitoring()
	c.Kill()
	return nil
}

// Kill kills the process.
func (c *Cmd) Kill() {
	if c.cmd.Process == nil {
		c.Warning(errors.New(c.Name(), "unable to kill; process is nil"))
		return
	}
	if err := c.cmd.Process.Kill(); err != nil {
		c.Warning(errors.New(c.Name(), fmt.Errorf("unable to kill: %v", err)))
	}
}

// Wait waits for the command to exit or ctx to be done. If ctx is done
// before the command exits, then an error is returned.
func (c *Cmd) Wait(ctx context.Context) error {
	select {
	case <-c.done:
		return nil
	case <-ctx.Done():
		select {
		case <-c.done:
			return nil
		default:
			return errors.New(c.Name(), ctx.Err())
		}
	}
}

// Monitor waits for cmd to exit, and when it does, logs an infrastructure failure
// if it exited abnormally.
func (c *Cmd) Monitor() {
	err := c.cmd.Wait()
	close(c.done)

	c.mu.RLock()
	defer c.mu.RUnlock()
	if err == nil || c.stopMonitoring {
		return
	}

	ee, ok := err.(*exec.ExitError)
	if !ok {
		return
	}
	signal := ee.Sys().(syscall.WaitStatus).Signal()
	exitCode := ee.Sys().(syscall.WaitStatus).ExitStatus()
	// KILL (0x9) and TERM (0xf) are normal when Forge is shutting down the
	// test (e.g., in the event of a timeout). In some cases, the shell sets
	// the exit code to 0x80 + the signal number.
	if signal == syscall.SIGKILL || signal == syscall.SIGTERM || exitCode == 0x80|0x09 || exitCode == 0x80|0x0f {
		return
	}
	c.Warning(errors.New(c.Name(), fmt.Errorf("exited prematurely with status: %v", err)))
}

// StdinPipe returns a pipe that will be connected to the command's standard input when the command starts.
func (c *Cmd) StdinPipe() (io.WriteCloser, error) {
	pipe, err := c.cmd.StdinPipe()
	if err != nil {
		return nil, errors.New(c.Name(), err)
	}
	return pipe, nil
}

// StopMonitoring turns off reporting of infrastructure failures should this process exit.
func (c *Cmd) StopMonitoring() {
	c.mu.Lock()
	c.stopMonitoring = true
	c.mu.Unlock()
}

// Healthy returns nil if c has been started and the process it started is still running.
func (c *Cmd) Healthy(ctx context.Context) error {
	if err := c.Base.Healthy(ctx); err != nil {
		return err
	}
	select {
	case <-c.done:
		return errors.NewPermanent(c.Name(), "executable has exited.")
	default:
	}
	return nil
}
