# Bazel Web Testing Rules

[![Build Status](https://travis-ci.org/bazelbuild/rules_webtesting.svg?branch=master)](https://travis-ci.org/bazelbuild/rules_webtesting)

Bazel rules and supporting code to allow testing against a browser with
WebDriver.

## Configure your Bazel project

Add the following to your WORKSPACE file:

```bzl
# Load rules_go at master for example purposes only. You should specify
# a specific version in your project.
http_archive(
    name = "io_bazel_rules_go",
    strip_prefix = "rules_go-master",
    urls = [
        "https://github.com/bazelbuild/rules_go/archive/master.tar.gz",
    ],
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

# Load rules_webtesting at master for example purposes only. You should specify
# a specific version in your project.
http_archive(
    name = "io_bazel_rules_webtesting",
    strip_prefix = "rules_webtesting-master",
    urls = [
        "https://github.com/bazelbuild/rules_webtesting/archive/master.tar.gz",
    ],
)

load("@io_bazel_rules_webtesting//web:repositories.bzl", "web_test_repositories")
load("@io_bazel_rules_webtesting//web/internal:platform_http_file.bzl", "platform_http_file")

web_test_repositories()

platform_http_file(
    name = "org_chromium_chromium",
    licenses = ["notice"],  # BSD 3-clause (maybe more?)
    amd64_sha256 =
        "6933d0afce6e17304b62029fbbd246cbe9e130eb0d90d7682d3765d3dbc8e1c8",
    amd64_urls = [
        "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/561732/chrome-linux.zip",
    ],
    macos_sha256 =
        "084884e91841a923d7b6e81101f0105bbc3b0026f9f6f7a3477f5b313ee89e32",
    macos_urls = [
        "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/561733/chrome-mac.zip",
    ],
    windows_sha256 =
        "d1bb728118c12ea436d8ea07dba980789e7d860aa664dd1fad78bc20e8d9391c",
    windows_urls = [
        "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win_x64/540270/chrome-win32.zip",
    ],
)

platform_http_file(
    name = "org_chromium_chromedriver",
    licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
    amd64_sha256 =
        "71eafe087900dbca4bc0b354a1d172df48b31a4a502e21f7c7b156d7e76c95c7",
    amd64_urls = [
        "https://chromedriver.storage.googleapis.com/2.41/chromedriver_linux64.zip",
    ],
    macos_sha256 =
        "fd32a27148f44796a55f5ce3397015c89ebd9f600d9dda2bcaca54575e2497ae",
    macos_urls = [
        "https://chromedriver.storage.googleapis.com/2.41/chromedriver_mac64.zip",
    ],
    windows_sha256 =
        "a8fa028acebef7b931ef9cb093f02865f9f7495e49351f556e919f7be77f072e",
    windows_urls = [
        "https://chromedriver.storage.googleapis.com/2.38/chromedriver_win32.zip",
    ],
)

platform_http_file(
    name = "org_mozilla_firefox",
    licenses = ["reciprocal"],  # MPL 2.0
    amd64_sha256 =
        "3a729ddcb1e0f5d63933177a35177ac6172f12edbf9fbbbf45305f49333608de",
    amd64_urls = [
        "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/61.0.2/linux-x86_64/en-US/firefox-61.0.2.tar.bz2",
        "https://ftp.mozilla.org/pub/firefox/releases/61.0.2/linux-x86_64/en-US/firefox-61.0.2.tar.bz2",
    ],
    macos_sha256 =
        "bf23f659ae34832605dd0576affcca060d1077b7bf7395bc9874f62b84936dc5",
    macos_urls = [
        "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/61.0.2/mac/en-US/Firefox%2061.0.2.dmg",
        "https://ftp.mozilla.org/pub/firefox/releases/61.0.2/mac/en-US/Firefox%2061.0.2.dmg",
    ],
)

platform_http_file(
    name = "org_mozilla_geckodriver",
    licenses = ["reciprocal"],  # MPL 2.0
    amd64_sha256 =
        "c9ae92348cf00aa719be6337a608fae8304691a95668e8e338d92623ba9e0ec6",
    amd64_urls = [
        "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-linux64.tar.gz",
        "https://github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-linux64.tar.gz",
    ],
    macos_sha256 =
        "ce4a3e9d706db94e8760988de1ad562630412fa8cf898819572522be584f01ce",
    macos_urls = [
        "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-macos.tar.gz",
        "https://github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-macos.tar.gz",
    ],
)
```

## Write your tests

Write your test in the language of your choice, but use our provided Browser API
to get an instance of WebDriver.

Example Test (Java):

```java
import com.google.testing.web.WebTest;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;
import org.openqa.selenium.WebDriver;

@RunWith(JUnit4.class)
public class BrowserTest {
  private WebDriver driver;

  @Before public void createDriver() {
    driver = new WebTest().newWebDriverSession();
  }

  @After public void quitDriver() {
    try {
      driver.quit();
     } finally {
      driver = null;
     }
   }

  // your tests here
}
```

Example Test (Go):

```go
import (
    "testing"

    "github.com/tebeka/selenium"
    "github.com/bazelbuild/rules_webtesting/go/webtest"
)

func TestWebApp(t *testing.T) {
    wd, err := webtest.NewWebDriverSession(selenium.Capabilities{})
    if err != nil {
        t.Fatal(err)
    }

    // your test here

    if err := wd.Quit(); err != nil {
        t.Logf("Error quitting webdriver: %v", err)
    }
}
```

Example Test (Python):

```python
import unittest
from testing.web import webtest


class BrowserTest(unittest.TestCase):
  def setUp(self):
    self.driver = webtest.new_webdriver_session()

  def tearDown(self):
    try:
      self.driver.quit()
    finally:
      self.driver = None

  # Your tests here

if __name__ == "__main__":
  unittest.main()
```

In your BUILD files, load the correct language specific build rule and create a
test target using it:

```bzl
load("@io_bazel_rules_webtesting//web:py.bzl", "py_web_test_suite")

py_web_test_suite(
    name = "browser_test",
    srcs = ["browser_test.py"],
    browsers = [
        # For experimental purposes only. Eventually you should
        # create your own browser definitions.
        "@io_bazel_rules_webtesting//browsers:chromium-native",
    ],
    local = True,
    deps = ["@io_bazel_rules_webtesting//testing/web"],
)
```
