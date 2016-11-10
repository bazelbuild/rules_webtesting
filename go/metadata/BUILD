load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

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

go_binary(
    name = "merger",
    srcs = ["merger.go"],
    visibility = ["//visibility:public"],
    deps = [
        ":metadata",
    ],
)

go_library(
    name = "metadata",
    srcs = [
        "metadata.go",
        "web_test_files.go",
    ],
    visibility = ["//go:__subpackages__"],
    deps = [
        ":capabilities",
        "//go/util:bazel",
    ],
)

go_library(
    name = "capabilities",
    srcs = ["capabilities.go"],
    visibility = ["//go:__subpackages__"],
)

go_test(
    name = "metadata_test",
    srcs = ["metadata_test.go"],
    data = [
        "//go/metadata/testdata:all-fields.json",
        "//go/metadata/testdata:android-browser-gingerbread-nexus-s.json",
        "//go/metadata/testdata:bad-named-files.json",
        "//go/metadata/testdata:chrome-linux.json",
        "//go/metadata/testdata:merge-from-file-result.json",
    ],
    library = ":metadata",
    deps = [
        "//go/util:bazel",
    ],
)

go_test(
    name = "capabilities_test",
    srcs = ["capabilities_test.go"],
    library = ":capabilities",
)

sh_test(
    name = "merger_test",
    srcs = ["merger_test.sh"],
    data = [
        ":merger",
        "//go/metadata/testdata:android-browser-gingerbread-nexus-s.json",
        "//go/metadata/testdata:bad-named-files.json",
        "//go/metadata/testdata:chrome-linux.json",
        "//go/metadata/testdata:merger-result.json",
        "//go/metadata/testdata:named-files1.json",
        "//go/metadata/testdata:named-files2.json",
    ],
)
