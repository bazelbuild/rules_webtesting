# Copyright 2017 Google Inc.
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

licenses(["notice"])  # Apache 2.0

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("//web:go.bzl", "go_web_test_suite")

go_library(
    name = "go_default_library",
    srcs = ["webtest.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//go/bazel:go_default_library",
        "//go/metadata:go_default_library",
        "@com_github_tebeka_selenium//:go_default_library",
    ],
)

go_web_test_suite(
    name = "go_default_test",
    srcs = [
        "browserinfo_test.go",
        "webtest_test.go",
    ],
    browser_overrides = {
        "//browsers:chrome55-win10-sauce": {
            "tags": ["requires-network"],
        },
    },
    browsers = [
        "//browsers:chrome55-win10-sauce",
        "//browsers:chromium-native",
        "//browsers:firefox-native",
    ],
    data = ["//go/metadata/testdata:all-fields.json"],
    go_test_tags = [
        "manual",
        "noci",
    ],
    library = ":go_default_library",
    local = True,
    tags = ["noci"],
    deps = ["@com_github_tebeka_selenium//:go_default_library"],
)
