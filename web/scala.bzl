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
"""Web Test rules for Java."""

load("//web/internal:wrap_web_test_suite.bzl", "wrap_web_test_suite")
load("@io_bazel_rules_scala//scala:scala.bzl", "scala_test")

def scala_web_test_suite(name, scala_test_tags = None, **kwargs):
    """Defines a test_suite of web_test targets that wrap a java_test target.

  Args:
    name: The base name of the test.
    scala_test_tags: A list of test tag strings to use for the scala_test
      target.
    **kwargs: Arguments for wrapped_web_test_suite
  """
    wrap_web_test_suite(
        name = name,
        rule = scala_test,
        wrapped_test_tags = scala_test_tags,
        **kwargs
    )
