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

licenses(["notice"])  # Apache 2.0

go_binary(
    name = "launcher",
    srcs = ["launcher.go"],
    visibility = ["//visibility:public"],
    deps = [
        ":cmdhelper",
        "//go/launcher/environments:chrome",
        "//go/launcher/environments:environment",
        "//go/launcher/environments:external",
        "//go/launcher/environments:firefox",
        "//go/launcher/environments:native",
        "//go/launcher/proxy",
        "//go/metadata",
        "//go/util:bazel",
    ],
)

go_library(
    name = "errors",
    srcs = ["errors.go"],
    visibility = ["//go:__subpackages__"],
)

go_test(
    name = "errors_test",
    srcs = ["errors_test.go"],
    library = ":errors",
)

go_library(
    name = "healthreporter",
    srcs = ["health_reporter.go"],
    visibility = ["//go:__subpackages__"],
    deps = [":errors"],
)

go_test(
    name = "healthreporter_test",
    srcs = ["health_reporter_test.go"],
    library = ":healthreporter",
    deps = [":errors"],
)

go_library(
    name = "cmdhelper",
    srcs = ["cmd_helper.go"],
    visibility = ["//go:__subpackages__"],
)
