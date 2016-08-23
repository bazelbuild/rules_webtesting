package browser

import (
	"strings"
	"testing"

	"github.com/tebeka/selenium/selenium"
)

func TestProvisionBrowser_NoCaps(t *testing.T) {
	wd, err := NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := wd.Get("about:"); err != nil {
		t.Error(err)
	}

	url, err := wd.CurrentURL()
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("Got empty url")
	}

	if err := wd.Quit(); err != nil {
		t.Error(err)
	}
}

func TestProvisionBrowser_WithCaps(t *testing.T) {
	wd, err := NewSession(selenium.Capabilities{
		"browserName":              "chrome",
		"unexpectedAlertBehaviour": "dismiss",
		"elementScrollBehavior":    1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := wd.Get("about:"); err != nil {
		t.Error(err)
	}

	url, err := wd.CurrentURL()
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("Got empty url")
	}

	if err := wd.Quit(); err != nil {
		t.Error(err)
	}
}

func TestGetInfo(t *testing.T) {
	i, err := GetInfo()

	if err != nil {
		t.Fatal(err)
	}

	if strings.HasSuffix(i.TargetBrowserName, "linux") {
		if i.Environment != "native" {
			t.Errorf(`got Environment = %q, expected "native"`, i.Environment)
		}
		if i.FormFactor != "DESKTOP" {
			t.Errorf(`got FormFactor = %q, expected "DESKTOP"`, i.FormFactor)
		}
	}
}
