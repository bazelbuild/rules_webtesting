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
"""Web Test rules for Go."""

load("//web/internal:constants.bzl", "DEFAULT_WRAPPED_TEST_TAGS")
load("//web/internal:wrap_web_test_suite.bzl", "wrap_web_test_suite")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

def go_web_test_suite(name, go_test_tags = DEFAULT_WRAPPED_TEST_TAGS, **kwargs):
    """Defines a test_suite of web_test targets that wrap a go_test target.

    Args:
        name: The base name of the test.
        go_test_tags: A list of test tag strings to use for the go_test target.
        **kwargs: Arguments for wrapped_web_test_suite
    """
    wrap_web_test_suite(
        name = name,
        rule = go_test,
        wrapped_test_tags = go_test_tags,
        **kwargs
    )
