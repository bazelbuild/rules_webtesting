# Copyright 2016 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Public definitions for web_test related build rules."""

load(
    "//web/internal:browser.bzl",
    browser_alias = "browser",
)
load("//web/internal:collections.bzl", "lists", "maps")
load("//web/internal:constants.bzl", "DEFAULT_TEST_SUITE_TAGS")
load(
    "//web/internal:web_test.bzl",
    web_test_alias = "web_test",
)
load(
    "//web/internal:web_test_archive.bzl",
    web_test_archive_alias = "web_test_archive",
)
load(
    "//web/internal:web_test_config.bzl",
    web_test_config_alias = "web_test_config",
)
load(
    "//web/internal:web_test_files.bzl",
    web_test_files_alias = "web_test_files",
)
load(
    "//web/internal:web_test_named_executable.bzl",
    web_test_named_executable_alias = "web_test_named_executable",
)
load(
    "//web/internal:web_test_named_file.bzl",
    web_test_named_file_alias = "web_test_named_file",
)

def web_test_suite(
        name,
        browsers,
        browser_overrides = None,
        test_suite_tags = None,
        visibility = None,
        **kwargs):
    """Defines a test_suite of web_test targets to be run.

    Args:
        name: Name; required. A unique name for this rule.
        browsers: List of labels; required. The browsers with which to run the test.
        browser_overrides: Dictionary; optional; default is an empty dictionary. A
          dictionary mapping from browser names to browser-specific web_test
          attributes, such as shard_count, flakiness, timeout, etc. For example:
          {'//browsers:chrome-native': {'shard_count': 3, 'flaky': 1}
           '//browsers:firefox-native': {'shard_count': 1, 'timeout': 100}}.
        test_suite_tags: List of strings; tags for the generated test_suite rule.
        visibility: List of labels; optional.
        **kwargs: Additional arguments for web_test rule.
    """
    if not lists.is_list_like(browsers):
        fail("expected a sequence type for attribute 'browsers' but got '%s'" %
             type(browsers))
    if not browsers:
        fail("expected non-empty value for attribute 'browsers'")

        # Check explicitly for None so that users can set this to the empty list.
    if test_suite_tags == None:
        test_suite_tags = DEFAULT_TEST_SUITE_TAGS

    tests = []
    browser_overrides = browser_overrides or {}

    for browser in browsers:  # pylint: disable=redefined-outer-name
        unqualified_browser = browser.split(":", 2)[1]
        test_name = name + "_" + unqualified_browser

        # Replace current browser attributes with those specified in the browser
        # overrides.
        overrides = browser_overrides.get(browser) or browser_overrides.get(
            unqualified_browser,
        ) or {}
        overridden_kwargs = _apply_browser_overrides(kwargs, overrides)
        if not "tags" in overridden_kwargs:
            overridden_kwargs["tags"] = []
        overridden_kwargs["tags"] = lists.clone(overridden_kwargs["tags"]) + ["browser:" + unqualified_browser]

        web_test(
            name = test_name,
            browser = browser,
            visibility = visibility,
            **overridden_kwargs
        )
        tests += [test_name]

    native.test_suite(
        name = name,
        tests = tests,
        tags = test_suite_tags,
        visibility = visibility,
    )

def _apply_browser_overrides(kwargs, overrides):
    """Handles browser-specific options that override the top-level definitions.

    Args:
        kwargs: A dictionary of arguments that will be overridden.
        overrides: A dictionary of attributes with the new attributes that should
          replace the top-level definitions.

    Returns:
        A dictionary of updated attributes.  For example:
        {'shard_count': 4, 'size': 'medium', 'timeout': 100, 'flaky': 1}
    """
    overridden_kwargs = maps.clone(kwargs)
    overridden_kwargs.update(overrides)

    return overridden_kwargs

def browser(testonly = True, **kwargs):
    """Wrapper around browser to correctly set defaults."""
    browser_alias(testonly = testonly, **kwargs)

def web_test(config = None, launcher = None, size = None, **kwargs):
    """Wrapper around web_test to correctly set defaults."""
    config = config or str(Label("//web:default_config"))
    launcher = launcher or str(Label("//go/wtl/main"))
    size = size or "large"
    web_test_alias(config = config, launcher = launcher, size = size, **kwargs)

def web_test_config(testonly = True, **kwargs):
    """Wrapper around web_test_config to correctly set defaults."""
    web_test_config_alias(testonly = testonly, **kwargs)

def web_test_named_executable(testonly = True, **kwargs):
    """Wrapper around web_test_named_executable to correctly set defaults."""
    web_test_named_executable_alias(testonly = testonly, **kwargs)

def web_test_named_file(testonly = True, **kwargs):
    """Wrapper around web_test_named_file to correctly set defaults."""
    web_test_named_file_alias(testonly = testonly, **kwargs)

def web_test_archive(testonly = True, **kwargs):
    """Wrapper around web_test_archive to correctly set defaults."""
    web_test_archive_alias(testonly = testonly, **kwargs)

def web_test_files(testonly = True, **kwargs):
    """Wrapper around web_test_files to correctly set defaults."""
    web_test_files_alias(testonly = testonly, **kwargs)
