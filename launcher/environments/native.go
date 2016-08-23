// Package native provides an Env for launching a browser locally using
// GoogleSeleniumServer.
package native

import (
	"context"

	"github.com/web_test_launcher/launcher/environments/environment"
	"github.com/web_test_launcher/launcher/services/selenium"
	"github.com/web_test_launcher/metadata/metadata"
)

const compName = "native environment"

type native struct {
	*environment.Base
	selenium *selenium.Selenium
}

// NewEnv creates a new environment for launching a browser locally using
// GoogleSeleniumServer.
func NewEnv(m metadata.Metadata) (environment.Env, error) {
	base, err := environment.NewBase(compName, m)
	if err != nil {
		return nil, err
	}

	s, err := selenium.New(nil)
	if err != nil {
		return nil, err
	}

	return &native{
		Base:     base,
		selenium: s,
	}, nil
}

func (n *native) SetUp(ctx context.Context) error {
	if err := n.Base.SetUp(ctx); err != nil {
		return err
	}
	return n.selenium.Start(ctx)
}

func (n *native) TearDown(ctx context.Context) error {
	if err := n.Base.TearDown(ctx); err != nil {
		return err
	}
	return n.selenium.Stop(ctx)
}

func (n *native) WDAddress(context.Context) string {
	return n.selenium.Address
}

func (n *native) Healthy(ctx context.Context) error {
	if err := n.Base.Healthy(ctx); err != nil {
		return err
	}
	return n.selenium.Healthy(ctx)
}
