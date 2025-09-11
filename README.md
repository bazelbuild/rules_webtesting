# Bazel Web Testing Rules

NOTE: This project is no longer being updated.

[![Build status](https://badge.buildkite.com/d9c3974f925876394ca9d3e00670c0950b6f7ebf325412def7.svg?branch=master)](https://buildkite.com/bazel/rules-webtesting-saucelabs)

Bazel rules and supporting code to allow testing against a browser with
WebDriver.

## Configure your Bazel project

For each language that is used by your project, you need to add the corresponding sub module to your MODULE.bazel file:

```bzl
bazel_dep(name = "rules_webtesting", version = "0.4.0")
```
```bzl
bazel_dep(name = "rules_webtesting_${language}", version = "0.4.0")
```
For example for Java:
```bzl
bazel_dep(name = "rules_webtesting_java", version = "0.4.0")
```

## Write your tests

Write your test in the language of your choice, but use our provided Browser API
to get an instance of WebDriver.

### Example Java Test

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

### Example Python Test

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

### Example Go Test

```go
import (
    "testing"

    "github.com/tebeka/selenium"
    "github.com/bazelbuild/rules_web_testing_go/webtest"
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

### BUILD file

In your BUILD files, load the correct language specific build rule and create a
test target using it:

```bzl
load("@rules_web_testing_python//web:py.bzl", "py_web_test_suite")

py_web_test_suite(
    name = "browser_test",
    srcs = ["browser_test.py"],
    browsers = [
        "@rules_webtesting//browsers:chromium-local",
    ],
    local = True,
    deps = ["@rules_webtesting//testing/web"],
)
```
