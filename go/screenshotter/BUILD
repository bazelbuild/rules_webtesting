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

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//web:go.bzl", "go_web_test_suite")

go_library(
    name = "screenshotter",
    srcs = ["screenshotter.go"],
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_tebeka_selenium//:selenium",
    ],
)

go_web_test_suite(
    name = "screenshotter_test",
    srcs = ["screenshotter_test.go"],
    browsers = ["//browsers:chromium-native"],
    data = ["//go/launcher/proxy/testdata:testpage.html"],
    go_test_tags = [
        "manual",
        "noci",
    ],
    library = ":screenshotter",
    local = True,
    tags = ["noci"],
    deps = [
        "//go:webtest",
        "//go/util:bazel",
        "@com_github_tebeka_selenium//:selenium",
    ],
)
