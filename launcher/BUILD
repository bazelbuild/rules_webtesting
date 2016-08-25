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
package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        ":cmdhelper",
        "//launcher/environments:environment",
        "//launcher/environments:external",
        "//launcher/environments:native",
        "//launcher/proxy",
        "//launcher/services:selenium",
        "//metadata",
        "//util:bazel",
    ],
)

go_library(
    name = "errors",
    srcs = ["errors.go"],
)

go_test(
    name = "errors_test",
    srcs = ["errors_test.go"],
    library = ":errors",
)

go_library(
    name = "healthreporter",
    srcs = ["health_reporter.go"],
    deps = [":errors"],
)

go_test(
    name = "healthreporter_test",
    srcs = ["health_reporter_test.go"],
    library = ":healthreporter",
)

go_library(
    name = "cmdhelper",
    srcs = ["cmd_helper.go"],
    deps = [":errors"],
)
