# Change Log

## Version Next

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
