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

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("//web:web.bzl", "web_test", "web_test_suite")

licenses(["notice"])  # Apache 2.0

go_library(
    name = "webtest",
    srcs = ["webtest.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//go/metadata",
        "//go/util:bazel",
        "@com_github_tebeka_selenium//:selenium",
    ],
)

go_test(
    name = "browserinfo_test",
    srcs = ["browserinfo_test.go"],
    data = [
        "//go/metadata/testdata:all-fields.json",
    ],
    library = ":webtest",
)

go_test(
    name = "webtest_test_wrapped",
    srcs = ["webtest_test.go"],
    library = ":webtest",
    tags = [
        "manual",
        "noci",
    ],
    deps = [
        "@com_github_tebeka_selenium//:selenium",
    ],
)

web_test_suite(
    name = "webtest_test",
    browsers = [
        "//browsers:chromium-native",
        "//browsers:firefox-native",
        "//browsers:phantomjs",
    ],
    local = True,
    test = ":webtest_test_wrapped",
)
