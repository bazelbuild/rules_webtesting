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

load("//web/internal:shared.bzl",
     "is_list_like",)
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
    size: String; optional; default is 'large'
    args: String list; optional; list of arguments to pass to test.
    tags: String list; optional.
    timeout: String; optional.
    browser_overrides: Dictionary; optional; default is an empty dictionary. A
      dictionary mapping from browser names to browser-specific web_test
      attributes, such as shard_count, flakiness, timeout, etc. For example:
      {'\\browsers:chrome-native': {'shard_count': 3, 'flaky': 1}
       '\\browsers:firefox-native': {'shard_count': 1, 'size': 'medium',
         'timeout': 100}}.
    flaky: Boolean; optional.
    config: Label; optional; default is //external:web_test_default_config.
      Configuration of web test features.
    visibility: Label List; optional.
    local: Boolean; optional.
  """
  # pylint: disable=unidiomatic-typecheck
  if not is_list_like(browsers):
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
    overrides = browser_overrides.get(browser) or {}
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


def browser(testonly=True, **kwargs):
  """Wrapper around browser to correctly set defaults.

  Args:
    testonly: default = True
    **kwargs: see browser in //web/internal:browser.bzl
  """
  if testonly == None:
    testonly = True
  browser_alias(testonly=testonly, **kwargs)


def web_test(size="large", **kwargs):
  """Wrapper around web_test to correctly set defaults.
  
  Args:
    size: default = "large"
    **kwargs: see web_test in //web/internal:web_test.bzl
  """
  size = size or "large"
  web_test_alias(size=size, **kwargs)


def web_test_config(testonly=True, **kwargs):
  """Wrapper around web_test_config to correctly set defaults.

  Args:
    testonly: default = True
    **kwargs: see web_test_config in //web/internal:web_test_config.bzl
  """
  if testonly == None:
    testonly = True
  web_test_config_alias(testonly=testonly, **kwargs)


def web_test_named_executable(testonly=True, **kwargs):
  """Wrapper around web_test_named_executable to correctly set defaults.

  Args:
    testonly: default = True
    **kwargs: see web_test_named_executable in 
      //web/internal:web_test_named_executable.bzl
  """
  if testonly == None:
    testonly = True
  web_test_named_executable_alias(testonly=testonly, **kwargs)
