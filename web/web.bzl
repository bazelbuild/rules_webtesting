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

load("//web/internal:collections.bzl", "lists", "maps")
load(
    "//web/internal:browser.bzl",
    browser_alias="browser",
)
load("//web/internal:constants.bzl", "DEFAULT_TEST_SUITE_TAGS")
load(
    "//web/internal:web_test.bzl",
    web_test_alias="web_test",
)
load(
    "//web/internal:web_test_config.bzl",
    web_test_config_alias="web_test_config")
load(
    "//web/internal:web_test_named_executable.bzl",
    web_test_named_executable_alias="web_test_named_executable")
load(
    "//web/internal:web_test_named_file.bzl",
    web_test_named_file_alias="web_test_named_file")
load(
    "//web/internal:web_test_archive.bzl",
    web_test_archive_alias="web_test_archive")
load("//web/internal:web_test_files.bzl", web_test_files_alias="web_test_files")


def web_test_suite(name,
                   browsers,
                   test,
                   browser_overrides=None,
                   test_suite_tags=DEFAULT_TEST_SUITE_TAGS,
                   visibility=None,
                   **kwargs):
  """Defines a test_suite of web_test targets to be run.

  Args:
    name: Name; required. A unique name for this rule.
    browsers: List of labels; required. The browsers with which to run the test.
    test: Label; required. A single *_test or *_binary target. The test that
      web_test should run with the specified browser.
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
        unqualified_browser) or {}
    overridden_kwargs = _apply_browser_overrides(kwargs, overrides)

    web_test(
        name=test_name,
        browser=browser,
        test=test,
        visibility=visibility,
        **overridden_kwargs)
    tests += [test_name]

  native.test_suite(
      name=name, tests=tests, tags=test_suite_tags, visibility=visibility)


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


def browser(deps=None, data=None, **kwargs):
  """Wrapper around browser to correctly set defaults."""
  data = lists.clone(data)
  if deps:
    lists.ensure_contains_all(data, deps)
  browser_alias(data=data, deps=deps, testonly=True, **kwargs)


def web_test(
    browser,  # pylint: disable=redefined-outer-name
    test,
    config=None,
    data=None,
    launcher=None,
    size=None,
    **kwargs):
  """Wrapper around web_test to correctly set defaults."""
  config = config or str(Label("//web:default_config"))
  launcher = launcher or str(Label("//go/launcher/main"))
  data = lists.clone(data)
  lists.ensure_contains_all(data, [browser, config, launcher, test])
  size = size or "large"
  web_test_alias(
      browser=browser,
      config=config,
      launcher=launcher,
      test=test,
      data=data,
      size=size,
      **kwargs)


def web_test_config(**kwargs):
  """Wrapper around web_test_config to correctly set defaults."""
  web_test_config_alias(testonly=True, **kwargs)


def web_test_named_executable(executable, data=None, **kwargs):
  """Wrapper around web_test_named_executable to correctly set defaults."""
  data = lists.clone(data)
  lists.ensure_contains(data, executable)
  web_test_named_executable_alias(
      data=data, executable=executable, testonly=True, **kwargs)


def web_test_named_file(**kwargs):
  """Wrapper around web_test_named_file to correctly set defaults."""
  web_test_named_file_alias(testonly=True, **kwargs)


def web_test_archive(**kwargs):
  """Wrapper around web_test_archive to correctly set defaults."""
  web_test_archive_alias(testonly=True, **kwargs)


def web_test_files(**kwargs):
  """Wrapper around web_test_files to correctly set defaults."""
  web_test_files_alias(testonly=True, **kwargs)
