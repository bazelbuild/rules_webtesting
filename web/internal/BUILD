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

load("@io_bazel_skydoc//skylark:skylark.bzl", "skylark_library")

licenses(["notice"])  # Apache 2.0

exports_files(
    glob(["*.sh.template"]),
    visibility = ["//visibility:public"],
)

exports_files(
    glob(["*.bzl"]),
    visibility = ["//doc:__pkg__"],
)

skylark_library(
    name = "browser",
    srcs = ["browser.bzl"],
    deps = [":metadata"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "collections",
    srcs = ["collections.bzl"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "constants",
    srcs = ["constants.bzl"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "files",
    srcs = ["files.bzl"],
    deps = [":collections"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "metadata",
    srcs = ["metadata.bzl"],
    deps = [":files"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "platform_http_file",
    srcs = ["platform_http_file.bzl"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "web_test",
    srcs = ["web_test.bzl"],
    deps = [
        ":collections",
        ":files",
        ":metadata",
    ],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "web_test_archive",
    srcs = ["web_test_archive.bzl"],
    deps = [":metadata"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "web_test_config",
    srcs = ["web_test_config.bzl"],
    deps = [":metadata"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "web_test_named_executable",
    srcs = ["web_test_named_executable.bzl"],
    deps = [":metadata"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "web_test_named_file",
    srcs = ["web_test_named_file.bzl"],
    deps = [":metadata"],
    visibility = ["//visibility:public"],
)

skylark_library(
    name = "wrap_web_test_suite",
    srcs = ["wrap_web_test_suite.bzl"],
    deps = [
        ":constants",
        "//web",
    ],
    visibility = ["//visibility:public"],
)
