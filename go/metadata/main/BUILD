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
    srcs = ["merger.go"],
    visibility = ["//visibility:private"],
    deps = ["//go/metadata:go_default_library"],
)

go_binary(
    name = "main",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

sh_test(
    name = "merger_test",
    srcs = ["merger_test.sh"],
    data = [
        ":main",
        "//go/metadata/testdata:android-browser-gingerbread-nexus-s.json",
        "//go/metadata/testdata:bad-named-files.json",
        "//go/metadata/testdata:chrome-linux.json",
        "//go/metadata/testdata:merger-result.json",
        "//go/metadata/testdata:named-files1.json",
        "//go/metadata/testdata:named-files2.json",
    ],
)
