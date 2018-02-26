// Package local provides a basic environment for web tests locally.
package local

import (
	"context"
	"fmt"

	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment"
	"github.com/bazelbuild/rules_webtesting/go/wtl/service"
	"github.com/bazelbuild/rules_webtesting/go/wtl/service/wsl"
)

const (
	compName = "local environment"
)

type local struct {
	*environment.Base
	wsl *service.Server
}

// NewEnv creates a new environment for launching a local browser using WSL.
func NewEnv(m *metadata.Metadata, d diagnostics.Diagnostics) (environment.Env, error) {
	base, err := environment.NewBase(compName, m, d)
	if err != nil {
		return nil, err
	}
	wsl, err := wsl.New(d, m)
	if err != nil {
		return nil, err
	}

	return &local{
		Base: base,
		wsl:  wsl,
	}, nil
}

func (l *local) SetUp(ctx context.Context) error {
	if err := l.Base.SetUp(ctx); err != nil {
		return err
	}
	return l.wsl.Start(ctx)
}

func (l *local) TearDown(ctx context.Context) error {
	if err := l.Base.TearDown(ctx); err != nil {
		return err
	}
	return l.wsl.Stop(ctx)
}

func (l *local) WDAddress(context.Context) string {
	return fmt.Sprintf("http://%s/", l.wsl.Address())
}

func (l *local) Healthy(ctx context.Context) error {
	if err := l.Base.Healthy(ctx); err != nil {
		return err
	}
	return l.wsl.Healthy(ctx)
}
