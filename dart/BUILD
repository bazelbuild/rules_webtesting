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

load("@io_bazel_rules_dart//dart/build_rules:core.bzl", "dart_library")
load("//web:dart.bzl", "dart_web_test_suite")

licenses(["notice"])  # Apache 2.0

dart_library(
    name = "dart",
    srcs = glob(["lib/**/*.dart"]),
    enable_ddc = False,
    pub_pkg_name = "webtest",
    visibility = ["//visibility:public"],
    deps = ["@com_github_google_webdriver_dart//:webdriver"],
)

dart_web_test_suite(
    name = "webtest_test",
    srcs = glob(["test/**/*.dart"]),
    browser_overrides = {
        "//browsers:chrome55-win10-sauce": {
            "tags": ["noci", "requires-network"],
        },
    },
    browsers = [
        "//browsers:chrome55-win10-sauce",
        "//browsers:chromium-native",
    ],
    dart_test_tags = [
        "manual",
        "noci",
    ],
    local = True,
    tags = ["noci"],
    deps = [
        ":dart",
        "@org_dartlang_pub_test//:test",
    ],
)
