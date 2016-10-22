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
workspace(name = "io_bazel_rules_webtesting")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "0c0ec7b9c7935883cbfb2df48fbf524e857859a5c05ae1b24d5442956e6bb5e8",
    strip_prefix = "rules_go-0.2.0",
    url = "https://github.com/bazelbuild/rules_go/archive/0.2.0.tar.gz",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

load("//web:repositories.bzl", "browser_repositories", 
    "web_test_repositories")

web_test_repositories(
    go = True,
    java = True,
    python = True,
)

browser_repositories(
    chrome = True,
    firefox = True, 
    phantomjs = True,
)

maven_jar(
    name = "junit_junit",
    artifact = "junit:junit:4.12",
    sha1 = "2973d150c0dc1fefe998f834810d68f278ea58ec",
)

maven_jar(
    name = "com_google_truth_truth",
    artifact = "com.google.truth:truth:0.29",
    sha1 = "b6ad12d98295a7d17b3fe4b8969d0f7905626b30",
)

http_archive(
    name = "io_bazel_rules_sass",
    sha256 = "d39d40c39a0fa2c7d05230ccf95aac3628936e4e76c0379ad324ff0b8488160f",
    strip_prefix = "rules_sass-0.0.1",
    url = "https://github.com/bazelbuild/rules_sass/archive/0.0.1.tar.gz",
)

load("@io_bazel_rules_sass//sass:sass.bzl", "sass_repositories")

sass_repositories()

http_archive(
    name = "io_bazel_skydoc",
    sha256 = "256bf8b64269d21fd46b8696007b5b9ef10070d79c106e74fb37979c04b6d519",
    strip_prefix = "skydoc-c57ff682364dbb1ae808b769f9e3add77cdbfad1",
    url = "https://github.com/bazelbuild/skydoc/archive/c57ff682364dbb1ae808b769f9e3add77cdbfad1.tar.gz",
)

load("@io_bazel_skydoc//skylark:skylark.bzl", "skydoc_repositories")

skydoc_repositories()
