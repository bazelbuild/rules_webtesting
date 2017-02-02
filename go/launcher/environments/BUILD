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

load("@io_bazel_rules_go//go:def.bzl", "go_library")

licenses(["notice"])  # Apache 2.0

go_library(
    name = "environment",
    srcs = ["environment.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        "//go/launcher:diagnostics",
        "//go/launcher:errors",
        "//go/launcher:healthreporter",
        "//go/metadata",
        "//go/metadata:capabilities",
    ],
)

go_library(
    name = "external",
    srcs = ["external.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":environment",
        "//go/launcher:diagnostics",
        "//go/launcher:errors",
        "//go/metadata",
    ],
)

go_library(
    name = "chrome",
    srcs = ["chrome.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":environment",
        "//go/launcher:cmdhelper",
        "//go/launcher:diagnostics",
        "//go/launcher/services:chromedriver",
        "//go/metadata",
    ],
)

go_library(
    name = "firefox",
    srcs = ["firefox.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":environment",
        "//go/launcher:cmdhelper",
        "//go/launcher:diagnostics",
        "//go/launcher/services:geckodriver",
        "//go/metadata",
    ],
)

go_library(
    name = "sauce",
    srcs = ["sauce.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":environment",
        "//go/launcher:diagnostics",
        "//go/metadata",
    ],
)