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

load("//web/internal:collections.bzl", "lists")
load(
    "//web/internal:browser.bzl",
    browser_alias="browser",)
load(
    "//web/internal:web_test.bzl",
    web_test_alias="web_test",)
load(
    "//web/internal:web_test_config.bzl",
    web_test_config_alias="web_test_config",)
load(
    "//web/internal:web_test_named_executable.bzl",
    web_test_named_executable_alias="web_test_named_executable")
load(
    "//web/internal:web_test_named_file.bzl",
    web_test_named_file_alias="web_test_named_file")
load(
    "//web/internal:web_test_archive.bzl",
    web_test_archive_alias="web_test_archive")


def web_test_suite(name,
                   browsers,
                   test,
                   args=None,
                   browser_overrides=None,
                   config=None,
                   data=None,
                   deprecation=None,
                   flaky=None,
                   local=None,
                   shard_count=None,
                   size=None,
                   tags=None,
                   timeout=None,
                   visibility=None):
  """Defines a test_suite of web_test targets to be run.

  Args:
    name: Name; required. A unique name for this rule.
    browsers: List of labels; required. The browsers with which to run the test.
    test: Label; required. A single *_test or *_binary target. The test that
      web_test should run with the specified browser.
    args: String list; optional; list of arguments to pass to test.
    browser_overrides: Dictionary; optional; default is an empty dictionary. A
      dictionary mapping from browser names to browser-specific web_test
      attributes, such as shard_count, flakiness, timeout, etc. For example:
      {'\\browsers:chrome-native': {'shard_count': 3, 'flaky': 1}
       '\\browsers:firefox-native': {'shard_count': 1, 'timeout': 100}}.
    config: Label; optional; default is //external:web_test_default_config.
      Configuration of web test features.
    data: Label List; optional.
    deprecation: String; optional.
    flaky: Boolean; optional.
    local: Boolean; optional.
    shard_count: Integer; optional; default is 1.
    size: String; optional; default is 'large'.
    tags: String list; optional.
    timeout: String; optional.
    visibility: Label List; optional.
  """
  if not lists.is_list_like(browsers):
    fail("expected value of type 'list' or 'tuple' for attribute 'browsers' " +
         "but got '%s'" % type(browsers))
  if not browsers:
    fail("expected non-empty list for attribute 'browsers'")

  tests = []
  browser_overrides = browser_overrides or {}

  for browser in browsers:  # pylint: disable=redefined-outer-name
    unqualified_browser = browser.split(":", 2)[1]
    test_name = name + "_" + unqualified_browser

    # Replace current browser attributes with those specified in the browser
    # overrides.
    overrides = browser_overrides.get(browser) or browser_overrides.get(
        unqualified_browser) or {}
    overridable_attributes = _apply_browser_overrides(
        config=config,
        data=data,
        flaky=flaky,
        local=local,
        shard_count=shard_count,
        size=size,
        tags=tags,
        timeout=timeout,
        overrides=overrides)

    web_test(
        name=test_name,
        args=args,
        browser=browser,
        deprecation=deprecation,
        test=test,
        visibility=visibility,
        **overridable_attributes)
    tests += [test_name]

  native.test_suite(
      name=name, tests=tests, tags=["manual"], visibility=visibility)


def _apply_browser_overrides(config, data, flaky, local, shard_count, size,
                             tags, timeout, overrides):
  """Handles browser-specific options that override the top-level definitions.

  Args:
    config: Label; optional; default is //testing/web/configs:default.
      Configuration of web test features.
    data: Additional data dependencies for the web_test() target.
    flaky: A boolean specifying the test is flaky.
    local: A boolean specifying the test should be run locally only.
    shard_count: The number of test shards to use per browser.
    size: A string specifying the test size.
    tags: Tags to use for a specific browser.
    timeout: A string specifying the test timeout.
    overrides: A dictionary of attributes with the new attributes that should
      replace the top-level definitions.

  Returns:
    A dictionary of updated attributes.  For example:
    {'shard_count': 4, 'size': 'medium', 'timeout': 100, 'flaky': 1}
  """

  output = {
      "data": data,
      "config": config,
      "flaky": flaky,
      "local": local,
      "shard_count": shard_count,
      "size": size,
      "timeout": timeout,
      "tags": tags,
  }
  for attribute in overrides:
    if attribute in output:
      output[attribute] = overrides[attribute]
    else:
      fail("Unrecognized attribute in browser_overrides: %s." % attribute)
  return output


def browser(testonly=True, **kwargs):
  """Wrapper around browser to correctly set defaults."""
  if testonly == None:
    testonly = True
  browser_alias(testonly=testonly, **kwargs)


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
  launcher = launcher or str(Label("//go/launcher"))
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


def web_test_config(testonly=True, **kwargs):
  """Wrapper around web_test_config to correctly set defaults."""
  if testonly == None:
    testonly = True
  web_test_config_alias(testonly=testonly, **kwargs)


def web_test_named_executable(executable, data=None, testonly=True, **kwargs):
  """Wrapper around web_test_named_executable to correctly set defaults."""
  data = lists.clone(data)
  lists.ensure_contains(data, executable)
  if testonly == None:
    testonly = True
  web_test_named_executable_alias(
      data=data, executable=executable, testonly=testonly, **kwargs)


def web_test_named_file(testonly=True, **kwargs):
  """Wrapper around web_test_named_file to correctly set defaults."""
  if testonly == None:
    testonly = True
  web_test_named_file_alias(testonly=testonly, **kwargs)


def web_test_archive(testonly=True, **kwargs):
  """Wrapper around web_test_archive to correctly set defaults."""
  if testonly == None:
    testonly = True
  web_test_archive_alias(testonly=testonly, **kwargs)
