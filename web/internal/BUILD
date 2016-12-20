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

licenses(["notice"])  # Apache 2.0

exports_files(
    glob(["*.sh.template"]),
    visibility = ["//visibility:public"],
)

exports_files(
    glob(["*.bzl"]),
    visibility = ["//doc:__pkg__"],
)

filegroup(
    name = "browser",
    srcs = ["browser.bzl"],
    data = [":metadata"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "collections",
    srcs = ["collections.bzl"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "files",
    srcs = ["files.bzl"],
    data = [":collections"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "metadata",
    srcs = ["metadata.bzl"],
    data = [":files"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "platform_http_file",
    srcs = ["platform_http_file.bzl"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "web_test",
    srcs = ["web_test.bzl"],
    data = [
        ":collections",
        ":files",
        ":metadata",
    ],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "web_test_archive",
    srcs = ["web_test_archive.bzl"],
    data = [":metadata"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "web_test_config",
    srcs = ["web_test_config.bzl"],
    data = [":metadata"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "web_test_named_executable",
    srcs = ["web_test_named_executable.bzl"],
    data = [":metadata"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "web_test_named_file",
    srcs = ["web_test_named_file.bzl"],
    data = [":metadata"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "wrap_web_test_suite",
    srcs = ["wrap_web_test_suite.bzl"],
    data = ["//web"],
    visibility = ["//visibility:public"],
)
