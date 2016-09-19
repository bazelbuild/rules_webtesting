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

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("//web:web.bzl", "web_test", "web_test_suite")

go_library(
    name = "browser",
    srcs = ["browser.go"],
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_tebeka_selenium//:selenium",
    ],
)

go_test(
    name = "browser_test_wrapped",
    srcs = ["browser_test.go"],
    library = ":browser",
    tags = ["manual"],
    deps = [
        "@com_github_tebeka_selenium//:selenium",
    ],
)

web_test_suite(
    name = "browser_test",
    browser_overrides = {
        "//browsers:chrome-external": {
            "tags": [
                "manual",
                "noci",
                "requires-network",
            ],
        },
        "//browsers:firefox-external": {
            "tags": [
                "manual",
                "noci",
                "requires-network",
            ],
        },
        "//browsers:chrome-native": {
            "local": True,
        },
        "//browsers:firefox-native": {
            "local": True,
        },
    },
    browsers = [
        "//browsers:chrome-external",
        "//browsers:chrome-native",
        "//browsers:firefox-external",
        "//browsers:firefox-native",
        "//browsers:phantomjs",
    ],
    tags = ["noci"],
    test = ":browser_test_wrapped",
)

web_test(
    name = "browser_test_ci_chrome-native",
    browser = select({
        "//browsers:mac": "//browsers:chrome-native",
        "//browsers:linux": "//browsers:disabled",
    }),
    test = ":browser_test_wrapped",
)

web_test(
    name = "browser_test_ci_firefox-native",
    browser = select({
        "//browsers:mac": "//browsers:firefox-native",
        "//browsers:linux": "//browsers:disabled",
    }),
    test = ":browser_test_wrapped",
)

web_test(
    name = "browser_test_ci_phantomjs",
    browser = select({
        "//browsers:mac": "//browsers:phantomjs",
        "//browsers:linux": "//browsers:disabled",
    }),
    test = ":browser_test_wrapped",
)
