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

licenses(["notice"])  # Apache 2.0

exports_files(
    glob(["*.bzl"]),
    visibility = ["//doc:__pkg__"],
)

web_test_config(
    name = "default_config",
    visibility = ["//visibility:public"],
)

filegroup(
    name = "repositories",
    srcs = ["repositories.bzl"],
    data = ["//web/internal:platform_http_file"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "web",
    srcs = ["web.bzl"],
    data = [
        "//web/internal:browser",
        "//web/internal:collections",
        "//web/internal:web_test",
        "//web/internal:web_test_archive",
        "//web/internal:web_test_config",
        "//web/internal:web_test_named_executable",
        "//web/internal:web_test_named_file",
    ],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "dart",
    srcs = ["dart.bzl"],
    data = [
        ":web",
        # should depend on dart build rules.
    ],
)

filegroup(
    name = "go",
    srcs = ["go.bzl"],
    data = [
        ":web",
        # should depend on go build rules.
    ],
)

filegroup(
    name = "java",
    srcs = ["java.bzl"],
    data = [":web"],
)

filegroup(
    name = "py",
    srcs = ["py.bzl"],
    data = [":web"],
)
