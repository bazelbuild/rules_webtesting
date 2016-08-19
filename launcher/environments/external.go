package external

import (
	"context"
	"fmt"
	"os"

	"github.com/web_test_launcher/launcher/environments/environment"
	"github.com/web_test_launcher/launcher/errors"
	"github.com/web_test_launcher/metadata/metadata"
)

const (
	name            = "External WebDriver Environment"
	address_env_var = "EXTERNAL_WEBDRIVER_SERVER_ADDRESS"
)

type external struct {
	*environment.Base
	address string
}

// NewEnv creates a new environment that uses an externally started Selenium Server.
func NewEnv(m metadata.Metadata) (environment.Env, error) {
	address, ok := os.LookupEnv(address_env_var)
	if !ok {
		return nil, errors.New(name, fmt.Errorf("environment variable %q not set", address_env_var))
	}

	base, err := environment.NewBase(name, m)
	if err != nil {
		return nil, err
	}

	return &external{
		Base:    base,
		address: address,
	}, nil
}

// WDAddress returns the user-provided selenium address.
func (e *external) WDAddress(context.Context) string {
	return e.address
}
