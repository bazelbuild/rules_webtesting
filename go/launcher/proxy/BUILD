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
    name = "proxy",
    srcs = ["proxy.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":driverhub",
        "//go/launcher:errors",
        "//go/launcher:healthreporter",
        "//go/launcher/environments:environment",
        "//go/metadata",
        "//go/util:httphelper",
        "//go/util:portpicker",
    ],
)

go_library(
    name = "driverhub",
    srcs = [
        "driver_hub.go",
        "driver_responses.go",
        "driver_session.go",
        "screenshot.go",
    ],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":webdriver",
        "//go/launcher:errors",
        "//go/launcher:healthreporter",
        "//go/launcher/environments:environment",
        "//go/metadata",
        "//go/util:httphelper",
        "@com_github_gorilla_mux//:mux",
    ],
)

go_library(
    name = "webdriver",
    srcs = ["webdriver.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        "//go/launcher:errors",
        "//go/launcher:healthreporter",
    ],
)
