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
    "//rules:browser.bzl",
    browser_alias="browser",)
load(
    "//rules:web_test.bzl",
    web_test_alias="web_test",)
load(
    "//rules:web_test_config.bzl",
    web_test_config_alias="web_test_config",)


def web_test_suite(name,
                   test,
                   browsers,
                   data=None,
                   deprecation=None,
                   shard_count=None,
                   size=None,
                   args=None,
                   tags=None,
                   timeout=None,
                   browser_overrides=None,
                   flaky=None,
                   config=None,
                   visibility=None,
                   local=None):
  """Defines a test_suite of web_test targets to be run.

  Args:
    name: Name; required. A unique name for this rule.
    test: Label; required. A single *_test or *_binary target. The test that
      web_test should run with the specified browser.
    browsers: List of labels; required. The browsers with which to run the test.
    data: Label List; optional.
    deprecation: String; optional.
    shard_count: Integer; optional; default is 1.
    size: String; optional.
    args: String list; optional; list of arguments to pass to test.
    tags: String list; optional.
    timeout: String; optional.
    browser_overrides: Dictionary; optional; default is an empty dictionary. A
      dictionary mapping from browser names to browser-specific web_test
      attributes, such as shard_count, flakiness, timeout, etc. For example:
      {'chrome-linux': {'shard_count': 3, 'flaky': 1}
       'firefox-linux': {'shard_count': 1, 'size': 'medium', 'timeout': 100}}.
    flaky: Boolean; optional.
    config: Label; optional; default is //testing/web/configs:default.
      Configuration of web test features.
    visibility: Label List; optional.
    local: Boolean; optional.
  """
  # pylint: disable=unidiomatic-typecheck
  if type(browsers) != type([]) and type(browsers) != type(()):
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
        unqualified_browser, {})
    overridable_attributes = _apply_browser_overrides(
        data=data or [],
        config=config,
        shard_count=shard_count,
        size=size,
        flaky=flaky,
        timeout=timeout,
        tags=tags,
        local=local,
        overrides=overrides)

    web_test(
        name=test_name,
        test=test,
        browser=browser,
        deprecation=deprecation,
        args=args,
        visibility=visibility,
        **overridable_attributes)
    tests += [test_name]

  native.test_suite(
      name=name, tests=tests, tags=["manual"], visibility=visibility)

def web_test(name,
             browser,  # pylint: disable=redefined-outer-name
             test,
             config=None,
             data=None,
             args=None,
             flaky=None,
             local=None,
             shard_count=None,
             size=None,
             timeout=None,
             deprecation=None,
             features=None,
             tags=None,
             visibility=None):
  """Runs a provided test against a provided browser configuration.

  This target is intended to be used for browser-based tests such as WebDriver
  or JsUnit tests.

  Args:
    name: Name; required. A unique name for this rule.
    browser: Label; required. The browser with which to run the test. To run
      your test against multiple browsers, you should be using the build
      extensions that generate web_test targets. See the Web Testing
      documentation for your test type, and for details of the supported
      browser configurations.
    test: Label; required. A single *_test or *_binary target. The test that
      web_test should run with the specified browser.
    config: Label; optional; default is //testing/web/configs:default.
      Configuration of web test features.
    data: Label List; optional.
    args: String List; optional. list of arguments to pass to test
    flaky: Boolean; optional; default is False.
    local: Boolean; optional; default is False.
    shard_count: Integer; optional; default is 1.
    size: String; optional; default is "large".
    timeout: String; optional.
    deprecation: String; optional.
    features: optional.
    tags: String List; optional.
    visibility: Label List; optional.
  """
  # "large" is the default size for web_test.
  size = size or "large"
  data = _clone_or_init_list(data)
  tags = _clone_or_init_list(tags)
  args = _clone_or_init_list(args)

  # Ensure necessary data is present.
  data = _ensure_in_list(data, test)
  data = _ensure_in_list(data, browser)

  # Ensure necessary tags are present.
  # The browser-* tags is used for metrics collection.
  browser_name = browser.split(":", 2)[1]
  tags = _ensure_in_list(tags, "browser-" + browser_name)

  web_test_alias(
      name=name,
      browser=browser,
      config=config,
      test=test,
      data=data,
      args=args,
      flaky=flaky,
      local=local,
      shard_count=shard_count,
      size=size,
      timeout=timeout,
      deprecation=deprecation,
      features=features,
      tags=tags,
      testonly=True,
      visibility=visibility)


def web_test_config(name, configs=None, record=None):
  """Configuration of web test features.

  Args:
    name: Name; required. A unique name for this rule.
    configs: Label List; optional. A list of configs to inherit from.
    record: String; optional; default is "". Whether to record, never record,
      or record video only if a test fails. Valid values are "", "never",
      "failed", "always".
  """
  web_test_config_alias(name=name, configs=configs, testonly=True)


def browser(name,
            metadata,
            disabled=None,
            data=None,
            deprecation=None,
            features=None,
            tags=None,
            visibility=None):
  """Defines a browser configuration for use with web_test.

  Args:
    name: Name; required. A unique name for this rule.
    metadata: Label; required.
    disabled: String; optional; default is "". Disables all tests for the
      browser. Tests will pass, but emit warnings (including this attribute
      value) that the test is disabled.
    data: Label List; optional.
    deprecation: String; optional.
    features: optional.
    tags: String List; optional.
    visibility: Label List; optional.
  """
  data = _clone_or_init_list(data)

  browser_alias(
      name=name,
      metadata=metadata,
      data=data,
      disabled=disabled,
      deprecation=deprecation,
      features=features,
      tags=tags,
      testonly=True,
      visibility=visibility)

  _run_browser(name)


def _run_browser(name):
  web_test(
      name=name + "_run",
      browser=":" + name,
      test="//test:debug_test",
      local=True,
      tags=["local", "manual"],
      timeout="eternal",)


def _apply_browser_overrides(data, config, shard_count, size, flaky, timeout,
                             tags, local, overrides):
  """Handles browser-specific options that override the top-level definitions.

  Args:
    data: Additional data dependencies for the web_test() target.
    config: Label; optional; default is //testing/web/configs:default.
      Configuration of web test features.
    shard_count: The number of test shards to use per browser.
    size: A string specifying the test size.
    flaky: A boolean specifying the test is flaky.
    timeout: A string specifying the test timeout.
    tags: Tags to use for a specific browser.
    local: A boolean specifying the test should be run locally only.
    overrides: A dictionary of attributes with the new attributes that should
      replace the top-level definitions.

  Returns:
    A dictionary of updated attributes.  For example:
    {'shard_count': 4, 'size': 'medium', 'timeout': 100, 'flaky': 1}
  """

  output = {"data": data,
            "config": config,
            "shard_count": shard_count,
            "size": size,
            "flaky": flaky,
            "timeout": timeout,
            "tags": tags,
            "local": local}
  for attribute in overrides:
    if attribute in output:
      output[attribute] = overrides[attribute]
    else:
      fail("Unrecognized attribute in browser_overrides: %s." % attribute)
  return output


def _ensure_in_list(lst, item):
  """Appends the specified item to the list if its not already a member."""
  if item not in lst:
    lst += [item]
  return lst


def _clone_or_init_list(original):
  new_list = []
  # pylint: disable=unidiomatic-typecheck
  if original and type(original) == type(""):
    fail('got "' + original + '", but expected a list')
  if original:
    new_list += original
  return new_list
