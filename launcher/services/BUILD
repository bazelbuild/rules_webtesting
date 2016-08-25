# Copyright 2016 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "service",
    srcs = [
        "cmd.go",
        "server.go",
        "service.go",
    ],
    deps = [
        "//launcher:cmdhelper",
        "//launcher:errors",
        "//launcher:healthreporter",
        "//util:bazel",
        "//util:httphelper",
        "//util:portpicker",
    ],
)

go_library(
    name = "selenium",
    srcs = ["selenium.go"],
    deps = [":service"],
)
