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
load("//web:go.bzl", "go_web_test_suite")

licenses(["notice"])  # Apache 2.0

go_library(
    name = "webtest",
    srcs = ["webtest.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//go/metadata",
        "//go/util:bazel",
        "@com_github_tebeka_selenium",
    ],
)

go_test(
    name = "browserinfo_test",
    srcs = ["browserinfo_test.go"],
    data = ["//go/metadata/testdata:all-fields.json"],
    library = ":webtest",
)

go_web_test_suite(
    name = "webtest_test",
    srcs = ["webtest_test.go"],
    browser_overrides = {
        "//browsers:chrome55-win10-sauce": {
            "tags": [
                "requires-network",
                "exclusive",
            ],
        },
    },
    browsers = [
        "//browsers:chrome55-win10-sauce",
        "//browsers:chromium-native",
        "//browsers:firefox-native",
    ],
    go_test_tags = [
        "manual",
        "noci",
    ],
    library = ":webtest",
    local = True,
    tags = ["noci"],
    deps = ["@com_github_tebeka_selenium"],
)
