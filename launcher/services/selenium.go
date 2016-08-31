package selenium

import (
	"fmt"
	"log"
	"time"

	"github.com/bazelbuild/rules_web/launcher/errors"
	"github.com/bazelbuild/rules_web/launcher/services/service"
	"github.com/bazelbuild/rules_web/metadata/metadata"
)

func NewSelenium(m metadata.Metadata, xvfb bool) (*service.Server, error) {
	seleniumPath, err := m.GetExecutablePath("SELENIUM_SERVER")
	if err != nil {
		return nil, errors.New("SeleniumServer", err)
	}

	args := []string{}

	if chromedriverPath, err := m.GetExecutablePath("CHROMEDRIVER"); err == nil {
		log.Printf("ChromeDriver found at: %q", chromedriverPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.chrome.driver=%s", chromedriverPath))
	}

	if geckodriverPath, err := m.GetExecutablePath("GECKODRIVER"); err == nil {
		log.Printf("GeckoDriver found at: %q", geckodriverPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.gecko.driver=%s", geckodriverPath))
	}

	args = append(args, "-port", "{port}")
	return service.NewServer(
		"SeleniumServer",
		seleniumPath,
		"http://%s/wd/hub/status",
		xvfb,
		60*time.Second,
		nil, args...)
}
