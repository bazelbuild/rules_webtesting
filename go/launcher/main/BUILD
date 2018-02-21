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

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "launcher.go",
    ],
    importpath = "github.com/bazelbuild/rules_webtesting/go/launcher/main",
    visibility = ["//visibility:private"],
    deps = [
        "//go/bazel:go_default_library",
        "//go/launcher/cmdhelper:go_default_library",
        "//go/launcher/diagnostics:go_default_library",
        "//go/launcher/environment:go_default_library",
        "//go/launcher/environment/chrome:go_default_library",
        "//go/launcher/environment/external:go_default_library",
        "//go/launcher/environment/firefox:go_default_library",
        "//go/launcher/environment/local:go_default_library",
        "//go/launcher/environment/sauce:go_default_library",
        "//go/launcher/errors:go_default_library",
        "//go/launcher/proxy:go_default_library",
        "//go/launcher/proxy/driverhub:go_default_library",
        "//go/launcher/proxy/driverhub/drivermu:go_default_library",
        "//go/launcher/proxy/driverhub/googlescreenshot:go_default_library",
        "//go/launcher/proxy/driverhub/mobileemulation:go_default_library",
        "//go/launcher/proxy/driverhub/quithandler:go_default_library",
        "//go/launcher/proxy/driverhub/scripttimeout:go_default_library",
        "//go/launcher/proxy/healthz:go_default_library",
        "//go/metadata:go_default_library",
    ],
)

go_binary(
    name = "main",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
