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

load("@io_bazel_skydoc//skylark:skylark.bzl", "skylark_doc")

licenses(["notice"])  # Apache 2.0

skylark_doc(
    name = "web-docs",
    srcs = [
        "//web:repositories.bzl",
        "//web:web.bzl",
        "//web/internal:browser.bzl",
        "//web/internal:platform_http_file.bzl",
        "//web/internal:web_test.bzl",
        "//web/internal:web_test_archive.bzl",
        "//web/internal:web_test_config.bzl",
        "//web/internal:web_test_named_executable.bzl",
        "//web/internal:web_test_named_file.bzl",
    ],
    tags = [
        "manual",
        "noci",
    ],
)
