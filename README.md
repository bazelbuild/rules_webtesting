# Bazel Web Testing Rules

Bazel rules and supporting code to allow testing against a browser with WebDriver.

## Configure your Bazel project

Add the following to your WORKSPACE file:

```bzl
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.2.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

git_repository(
    name = "io_bazel_rules_webtesting",
    remote = "https://github.com/bazelbuild/rules_webtesting.git",
    tag = "0.1.0",
)

load("@io_bazel_rules_webtesting//web:repositories.bzl", 
    "browser_repositories",
    "web_test_repositories")

web_test_repositories(
    # specify test languages your project is using.
    go = True,
    java = True,
    python = True,
)

# Load repositories for example browser definitions.
# You should create your own browser definitions and link
# to the specific browser versions you are interested in
# testing with.
browser_repositories(
    chrome = True,
    firefox = True,
    phantomjs = True,
)
```

This will configure the following repositories required to get web_test_suite
working:

*   [com_github_gorilla_mux](https://github.com/gorilla/mux)
*   [org_seleniumhq_server](http://www.seleniumhq.org/download/) -- Selenium
    Standalone Server
*   [org_seleniumhq_java](http://www.seleniumhq.org/download/) -- Java Client
    Binding (only if java = True)
*   [org_json_json](https://mvnrepository.com/artifact/org.json/json) (only if
    java = True)
*   [com_google_code_findbugs_jsr305](https://mvnrepository.com/artifact/com.google.code.findbugs/jsr305)
    (only if java = True)
*   [com_google_guava_guava](https://mvnrepository.com/artifact/com.google.guava/guava)
    (only if java = True)
*   [com_github_tebeka_selenium](https://github.com/tebeka/selenium) (only if
    go = True)
*   [org_seleniumhq_py](http://www.seleniumhq.org/download/) -- Python client
    binding (only if python = True)

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

    "github.com/tebeka/selenium/selenium"
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

In your BUILD files, create your test target as normal, but tag it "manual".
Then create a web_test_suite that depends on your test target:

```bzl
load("@io_bazel_rules_webtesting//web:web.bzl", "web_test_suite")

py_test(
    name = "browser_test_wrapped",
    srcs = ["browser_test.py"],
    tags = ["manual"],
    deps = [
        "@io_bazel_rules_webtesting//testing/web",
    ],
)

web_test_suite(
    name = "browser_test",
    browsers = [
        # For experimental purposes only. Eventually you should
        # create your own browser definitions.
        "@io_bazel_rules_webtesting//browsers:chrome-native",
    ],
    local = 1,
    test = ":browser_test_wrapped",
)
```

## [Rules reference](RULES.md)
