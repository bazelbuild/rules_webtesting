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

load(":web.bzl", "web_test_config")
load("@bazel_skylib//:skylark_library.bzl", "skylark_library")

licenses(["notice"])  # Apache 2.0

exports_files(
    glob(["*.bzl"]),
    visibility = ["//doc:__pkg__"],
)

web_test_config(
    name = "default_config",
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "repositories",
    srcs = ["repositories.bzl"],
    visibility = ["//visibility:public"],
    deps = ["//web/internal:platform_http_file"],
)

skylark_library(
    name = "web",
    srcs = ["web.bzl"],
    visibility = ["//visibility:public"],
    deps = [
        "//web/internal:browser",
        "//web/internal:collections",
        "//web/internal:constants",
        "//web/internal:web_test",
        "//web/internal:web_test_archive",
        "//web/internal:web_test_config",
        "//web/internal:web_test_named_executable",
        "//web/internal:web_test_named_file",
    ],
)

skylark_library(
    name = "go",
    srcs = ["go.bzl"],
    deps = [
        "//web/internal:wrap_web_test_suite",
        # should depend on go build rules.
    ],
)

skylark_library(
    name = "java",
    srcs = ["java.bzl"],
    deps = ["//web/internal:wrap_web_test_suite"],
)

skylark_library(
    name = "py",
    srcs = ["py.bzl"],
    deps = ["//web/internal:wrap_web_test_suite"],
)
