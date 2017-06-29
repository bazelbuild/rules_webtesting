# Bazel Web Testing Rules

[![Build
Status](http://ci.bazel.io/buildStatus/icon?job=rules_web)](http://ci.bazel.io/job/rules_web)

Bazel rules and supporting code to allow testing against a browser with
WebDriver.

## Configure your Bazel project

Add the following to your WORKSPACE file:

```bzl
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.4",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

git_repository(
    name = "io_bazel_rules_webtesting",
    remote = "https://github.com/bazelbuild/rules_webtesting.git",
    tag = "HEAD",
)

load("@io_bazel_rules_webtesting//web:repositories.bzl",
    "browser_repositories",
    "web_test_repositories")

web_test_repositories()

# Load repositories for example browser definitions.
# You should create your own browser definitions and link
# to the specific browser versions you are interested in
# testing with.
browser_repositories(
    chromium = True,
    firefox = True,
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
