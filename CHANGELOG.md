# Change Log

## Version 0.1.0

*   Change browser-provisioning API from browser... to webtest....
*   Add Chrome environment that uses ChromeDriver directly instead of Selenium
    Server Standalone.
*   Add Firefox environment that uses GeckoDriver directly instead of Selenium
    Server Standalone.
*   Support using WebDriver remote ends other than Selenium Server Standalone.

## Version 0.0.4

*   Added browser_repositories rule so users can get the browsers defined
    //browsers/... working for experimenting with rules_webtesting.
*   Rename to rules_webtesting.
*   No longer use bind for configuraion. Instead make the rule attributes that
    depended on these bind statements public.
*   Remove usages of git_repository.
*   Make the set of repositories loaded by web_test_repositories fully
    configurable.
*   Switch from platform-specific repositories to repositories that get
    configured based on the current platform.

## Version 0.0.3

*   Change Python rules to srcs_version = "PY2AND3"

## Version 0.0.2

*   Add PhantomJS browser for Linux.
*   Add support for named files in archives.
*   Add MacOS support for Chrome, Firefox, and PhantomJS browser.
*   Support placeholders for named file paths in capabilities.
*   Use template files to generate test scripts.
*   Add Python Library for use with web_test.
