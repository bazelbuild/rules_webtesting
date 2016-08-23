package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/web_test_launcher/launcher/errors"
	"github.com/web_test_launcher/launcher/healthreporter"
	"github.com/web_test_launcher/util/httphelper"
	"github.com/web_test_launcher/util/portpicker"
)

// Server is a service that starts an external server.
type Server struct {
	*Cmd
	Address       string
	port          int
	healthPattern string
	timeout       time.Duration
}

// NewServer creates a new service for starting an external server on the host machine.
// Args should contain {port}, which will be replaced with the selected port number.
func NewServer(name, exe, healthPattern string, timeout time.Duration, env map[string]string, args ...string) (*Server, error) {
	port, err := portpicker.PickUnusedPort()
	if err != nil {
		return nil, errors.New(name, err)
	}
	updatedArgs := []string{}
	for _, arg := range args {
		updatedArgs = append(updatedArgs, strings.Replace(arg, "{port}", strconv.Itoa(port), -1))
	}

	cmd, err := NewCmd(name, exe, env, updatedArgs...)
	if err != nil {
		return nil, err
	}
	return &Server{
		Cmd:           cmd,
		Address:       net.JoinHostPort("localhost", strconv.Itoa(port)),
		port:          port,
		healthPattern: healthPattern,
		timeout:       timeout,
	}, nil
}

// Start starts the server, waits for it to become healhy, and monitors it to ensure that it
// stays healthy.
func (s *Server) Start(ctx context.Context) error {
	if err := s.Cmd.Start(ctx); err != nil {
		return err
	}

	healthyCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	if err := healthreporter.WaitForHealthy(healthyCtx, s); err != nil {
		return errors.New(s.Name(), err)
	}
	return nil
}

// Stop stops the server.
func (s *Server) Stop(ctx context.Context) error {
	if err := s.Cmd.Stop(ctx); err != nil {
		return err
	}
	return nil
}

// Healthy returns nil if the server responds OK to requests to health page.
func (s *Server) Healthy(ctx context.Context) error {
	if err := s.Cmd.Healthy(ctx); err != nil {
		return err
	}

	url := fmt.Sprintf(s.healthPattern, s.Address)
	resp, err := httphelper.Get(ctx, url)
	if err != nil {
		return errors.New(s.Name(), fmt.Errorf("error getting %s: %v", url, err))
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(s.Name(), fmt.Errorf("request to %s returned status %v", url, resp.StatusCode))
	}
	return nil
}

// Port returns the port this server is running on as a string.
func (s *Server) Port() string {
	return fmt.Sprintf("%d", s.port)
}
