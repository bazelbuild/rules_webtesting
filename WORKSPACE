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
    sha256 = "890e055a9956daa85287ed309e8efaa9d5336b2bc5a71ad3676c220c99015a9d",
    strip_prefix = "rules_go-0.3.2",
    url = "https://github.com/bazelbuild/rules_go/archive/0.3.2.tar.gz",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

load("//web:repositories.bzl", "browser_repositories", "web_test_repositories")

web_test_repositories(
    go = True,
    java = True,
    python = True,
)

browser_repositories(
    chromium = True,
    firefox = True,
    phantomjs = True,
)

maven_jar(
    name = "junit_junit",
    artifact = "junit:junit:4.12",
    sha1 = "2973d150c0dc1fefe998f834810d68f278ea58ec",
)

http_archive(
    name = "io_bazel_rules_sass",
    sha256 = "4bd44d81747d06e8334570d413b714218b7759db0883df807e28127e9d59fe80",
    strip_prefix = "rules_sass-931508528093364b86abd44a5b9401e5150f1ba7",
    url = "https://github.com/bazelbuild/rules_sass/archive/931508528093364b86abd44a5b9401e5150f1ba7.tar.gz",
)

load("@io_bazel_rules_sass//sass:sass.bzl", "sass_repositories")

sass_repositories()

http_archive(
    name = "io_bazel_skydoc",
    sha256 = "06d855d8412cae1461d4131481a26d71ba9457914473803df65b110ea4dd6a88",
    strip_prefix = "skydoc-0.1.1",
    url = "https://github.com/bazelbuild/skydoc/archive/0.1.1.tar.gz",
)

load("@io_bazel_skydoc//skylark:skylark.bzl", "skydoc_repositories")

skydoc_repositories()

http_archive(
    name = "io_bazel_rules_dart",
    sha256 = "c46db0431001b1a4aa7fdd2ec43642c07e443eba4161a3b2b0d60bd19ff7370e",
    strip_prefix = "rules_dart-b7ef091c339a55cd321b92399462f24aecf42fd6",
    url = "https://github.com/dart-lang/rules_dart/archive/b7ef091c339a55cd321b92399462f24aecf42fd6.tar.gz",
)

load("@io_bazel_rules_dart//dart/build_rules:repositories.bzl", "dart_repositories")

dart_repositories()

http_archive(
    name = "com_github_google_webdriver_dart",
    sha256 = "675aa61528e86136386939faa8b0a626c8a58db01917f438e4d8b7b24d5b3f5c",
    strip_prefix = "webdriver.dart-1.2.2",
    url = "https://github.com/google/webdriver.dart/archive/v1.2.2.tar.gz",
)

load("@com_github_google_webdriver_dart//:repositories.bzl", "webdriver_dart_repositories")

webdriver_dart_repositories()
