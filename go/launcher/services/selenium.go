package selenium

import (
	"fmt"
	"log"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/services/service"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

type Selenium struct {
	*service.Server
}

func NewSelenium(m *metadata.Metadata, xvfb bool) (*Selenium, error) {
	seleniumPath, err := m.GetFilePath("SELENIUM_SERVER")
	if err != nil {
		return nil, errors.New("SeleniumServer", err)
	}

	args := []string{}

	if chromedriverPath, err := m.GetFilePath("CHROMEDRIVER"); err == nil {
		log.Printf("ChromeDriver found at: %q", chromedriverPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.chrome.driver=%s", chromedriverPath))
	}
	if geckodriverPath, err := m.GetFilePath("GECKODRIVER"); err == nil {
		log.Printf("GeckoDriver found at: %q", geckodriverPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.gecko.driver=%s", geckodriverPath))
	}
	if firefoxPath, err := m.GetFilePath("FIREFOX"); err == nil {
		log.Printf("Firefox found at: %q", firefoxPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.firefox.bin=%s", firefoxPath))
	}
	args = append(args, "-port", "{port}")
	server, err := service.NewServer(
		"SeleniumServer",
		seleniumPath,
		"http://%s/wd/hub/status",
		xvfb,
		60*time.Second,
		nil, args...)
	if err != nil {
		return nil, err
	}
	return &Selenium{server}, nil
}

func (s *Selenium) Address() string {
	return fmt.Sprintf("http://%s/wd/hub/", s.Server.Address())
}
