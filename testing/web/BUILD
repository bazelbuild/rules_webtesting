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
#
################################################################################
#
package(default_testonly = True)

load("//web:py.bzl", "py_web_test_suite")

licenses(["notice"])  # Apache 2.0

py_library(
    name = "web",
    srcs = glob(
        ["*.py"],
        exclude = ["*_test.py"],
    ),
    srcs_version = "PY2AND3",
    visibility = ["//visibility:public"],
    deps = ["@org_seleniumhq_py//:selenium"],
)

py_web_test_suite(
    name = "webtest_test",
    srcs = ["webtest_test.py"],
    browsers = [
        "//browsers:chromium-native",
        "//browsers:firefox-native",
    ],
    local = True,
    py_test_tags = [
        "manual",
        "noci",
    ],
    srcs_version = "PY2AND3",
    tags = ["noci"],
    deps = [":web"],
)
