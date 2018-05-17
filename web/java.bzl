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

def java_web_test_suite(name, java_test_tags = None, test_class = None, **kwargs):
    """Defines a test_suite of web_test targets that wrap a java_test target.

    Args:
        name: The base name of the test.
        java_test_tags: A list of test tag strings to use for the java_test target.
        test_class: Optional; default computed from name and blaze package.
        **kwargs: Arguments for wrapped_web_test_suite
    """
    if test_class == None:
        test_package = native.package_name().replace("javatests/", "")
        test_package = test_package.replace("/", ".")
        test_class = test_package + "." + name

    wrap_web_test_suite(
        name = name,
        rule = native.java_test,
        test_class = test_class,
        wrapped_test_tags = java_test_tags,
        **kwargs
    )
