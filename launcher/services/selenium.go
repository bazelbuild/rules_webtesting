// Package selenium provides a service for launching GoogleSeleniumServer on the local
// host.
package selenium

import (
	"time"

	"github.com/web_test_launcher/launcher/services/service"
)

// Selenium is a service that starts GoogleSeleniumServer.
type Selenium struct {
	*service.Server
}

// New creates a new service for starting GoogleSeleniumServer on the host machine.
func New(env map[string]string) (*Selenium, error) {
	server, err := service.NewServer(
		"SeleniumServer",
		"web_test_rules/java/SeleniumServer",
		"http://%s/wd/hub/status",
		60*time.Second,
		env, "-port", "{port}")
	if err != nil {
		return nil, err
	}
	return &Selenium{server}, nil
}
