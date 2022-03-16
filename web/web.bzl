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
    _browser = "browser",
)
load("//web/internal:constants.bzl", "DEFAULT_TEST_SUITE_TAGS", "DEFAULT_WEB_TEST_SUITE_TAGS")
load(
    "//web/internal:custom_browser.bzl",
    _custom_browser = "custom_browser",
)
load(
    "//web/internal:web_test.bzl",
    _web_test = "web_test",
)
load(
    "//web/internal:web_test_config.bzl",
    _web_test_config = "web_test_config",
)
load(
    "//web/internal:web_test_files.bzl",
    _web_test_files = "web_test_files",
)
load(
    "//web/internal:web_test_named_executable.bzl",
    _web_test_named_executable = "web_test_named_executable",
)
load(
    "//web/internal:web_test_named_file.bzl",
    _web_test_named_file = "web_test_named_file",
)
load("@bazel_skylib//lib:types.bzl", "types")

load("//web/internal:platform_archive.bzl", _platform_archive = "platform_archive")
load("//web/internal:platform_metadata.bzl", _platform_metadata = "platform_metadata")

# Expose the following rules publicly.
platform_archive = _platform_archive
platform_metadata = _platform_metadata

def web_test_suite(
        name,
        browsers,
        test_suite_tags = DEFAULT_TEST_SUITE_TAGS,
        visibility = None,
        **kwargs):
    """Defines a test_suite of web_test targets to be run.

    Args:
        name: Name; required. A unique name for this rule.
        browsers: List of labels; required. The browsers with which to run the test.
        test_suite_tags: List of strings; tags for the generated test_suite rule.
        visibility: List of labels; optional.
        **kwargs: Additional arguments for web_test rule.
    """
    if not types.is_list(browsers):
        fail("expected a list for attribute 'browsers' but got '%s'" %
             type(browsers))
    if not browsers:
        fail("expected non-empty value for attribute 'browsers'")

    tests = []

    if "tags" not in kwargs:
        kwargs["tags"] = DEFAULT_WEB_TEST_SUITE_TAGS

    for browser in browsers:  # pylint: disable=redefined-outer-name
        unqualified_browser = browser.split(":", 2)[1]
        test_name = name + "_" + unqualified_browser

        overridden_kwargs = _get_kwargs(unqualified_browser, kwargs)

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

def _get_kwargs(browser, in_kwargs):
    """Returns the arguments that should be used for a particular browser."""
    out_kwargs = {}

    for k, v in in_kwargs.items():
        if types.is_dict(v):
            if browser in v:
                out_kwargs[k] = v[browser]
            elif "default" in v:
                out_kwargs[k] = v["default"]
        else:
            out_kwargs[k] = v

    if not out_kwargs["tags"]:
        out_kwargs["tags"] = ["browser:" + browser]
    else:
        out_kwargs["tags"] = ["browser:" + browser] + out_kwargs["tags"]

    return out_kwargs

def browser(testonly = True, **kwargs):
    """Wrapper around browser to correctly set defaults."""
    _browser(testonly = testonly, **kwargs)

def custom_browser(testonly = True, **kwargs):
    """Wrapper around custom_browser to correctly set defaults."""
    _custom_browser(testonly = testonly, **kwargs)

def web_test(browser, config = None, launcher = None, size = None, **kwargs):
    """Wrapper around web_test to correctly set defaults."""
    config = config or str(Label("//web:default_config"))
    launcher = launcher or str(Label("//go/wtl/main"))
    size = size or "large"

    unqualified_browser = browser.split(":", 2)[1]

    if "tags" not in kwargs:
        kwargs["tags"] = DEFAULT_WEB_TEST_SUITE_TAGS.get(unqualified_browser, [])

    kwargs["tags"] = [unqualified_browser] + kwargs["tags"]

    _web_test(browser = browser, config = config, launcher = launcher, size = size, **kwargs)

def web_test_config(testonly = True, **kwargs):
    """Wrapper around web_test_config to correctly set defaults."""
    _web_test_config(testonly = testonly, **kwargs)

def web_test_named_executable(testonly = True, **kwargs):
    """Wrapper around web_test_named_executable to correctly set defaults."""
    _web_test_named_executable(testonly = testonly, **kwargs)

def web_test_named_file(testonly = True, **kwargs):
    """Wrapper around web_test_named_file to correctly set defaults."""
    _web_test_named_file(testonly = testonly, **kwargs)

def web_test_files(testonly = True, **kwargs):
    """Wrapper around web_test_files to correctly set defaults."""
    _web_test_files(testonly = testonly, **kwargs)
