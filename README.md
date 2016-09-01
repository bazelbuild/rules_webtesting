# Bazel Web Testing Rules

[TOC]

## Configure your Bazel project

Add the following to your WORKSPACE file:

```bzl
git_repository(
    name = "io_bazel_rules_go",
    # need a version with Go1.7 support
    commit = "ae8ea32be1af991eef77d6347591dc8ba56c40a2",
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

git_repository(
    name = "io_bazel_rules_web",
    remote = "https://github.com/bazelbuild/rules_web.git",
    tag = "",
)

load("@io_bazel_rules_web//web:repositories.bzl", "web_test_repositories")

web_test_repositories(
    # specify test languages your project is using.
    go = True,
    java = True,
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

## Write your tests

Write your test in the language of your choice, but use our provided Browser API
to get an instance of WebDriver.

Example Test (Java):

```java
import com.google.testing.web.Browser;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;
import org.openqa.selenium.WebDriver;

@RunWith(JUnit4.class)
public class BrowserTest {
  private WebDriver driver;

  @Before public void createDriver() {
    driver = new Browser().newSession();
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
    "github.com/bazelbuild/rules_web/go/browser"
)

func TestWebApp(t *testing.T) {
    wd, err := NewSession(selenium.Capabilities{})
    if err != nil {
        t.Fatal(err)
    }

    // your test here

    if err := wd.Quit(); err != nil {
        t.Logf("Error quitting webdriver: %v", err)
    }
}
```

In your BUILD files, create your test target as normal, but tag it "manual".
Then create a web_test_suite that depends on your test target:

```bzl
load("@io_bazel_rules_go//go:def.bzl", "go_test")
load("@io_bazel_rules_web//web:web.bzl", "web_test_suite")

go_test(
    name = "browser_test_wrapped",
    srcs = ["browser_test.go"],
    tags = ["manual"],
    deps = [
        "@com_github_tebeka_selenium//:selenium",
        "@io_bazel_rules_web//go:browser",
    ],
)

web_test_suite(
    name = "browser_test",
    browsers = [
        "@io_bazel_rules_web//browsers:chrome-native",
    ],
    local = 1,
    test = ":browser_test_wrapped",
)
```

## [Rules reference](RULES.md)