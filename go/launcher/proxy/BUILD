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
load("//web:web.bzl", "web_test_config", "web_test_suite")

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

go_library(
    name = "screenshot",
    srcs = ["screenshot.go"],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":driverhub",
        ":webdriver",
        "//go/util:cropper",
    ],
)

web_test_config(
    name = "screenshot_test_config",
    metadata = "screenshot_test.json",
)

go_test(
    name = "screenshot_test_wrapped",
    srcs = ["screenshot_test.go"],
    tags = [
        "manual",
        "noci",
    ],
    deps = [
        "//go:webtest",
        "//go/util:bazel",
        "@com_github_tebeka_selenium//:selenium",
    ],
)

web_test_suite(
    name = "screenshot_test",
    browsers = ["//browsers:chromium-native"],
    config = ":screenshot_test_config",
    local = True,
    tags = ["noci"],
    test = ":screenshot_test_wrapped",
)
