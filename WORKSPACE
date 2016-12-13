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

load(
    "//web:repositories.bzl",
    "browser_repositories",
    "web_test_repositories",
)

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

http_file(
    name = "fonts_noto_hinted_deb",
    sha256 = "25b362c9437a7859ce034f22d94b698e8ed25007b443e5a26228ed5b3d2d32d4",
    url = "http://bazel-mirror.storage.googleapis.com/http.us.debian.org/debian/pool/main/f/fonts-noto/fonts-noto-hinted_20160116-1_all.deb",
)

http_file(
    name = "fonts_noto_mono_deb",
    sha256 = "74b457715f275ed893998a70d6bc955f67da6d36b36b422dbeeb045160edacb6",
    url = "http://bazel-mirror.storage.googleapis.com/http.us.debian.org/debian/pool/main/f/fonts-noto/fonts-noto-mono_20160116-1_all.deb",
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
    sha256 = "04c7a457631c8af9ce47464f671ce1b7ddadb1e8d0d284218e8ef6a3623f27ba",
    strip_prefix = "skydoc-b14ff1002f53c24054960ede321090282644c4dc",
    url = "https://github.com/bazelbuild/skydoc/archive/b14ff1002f53c24054960ede321090282644c4dc.tar.gz",
)

load("@io_bazel_skydoc//skylark:skylark.bzl", "skydoc_repositories")

skydoc_repositories()

http_archive(
    name = "io_bazel_rules_dart",
    sha256 = "83220a6d89b2299887506911a4a12439190baa1633f6da2fd32a67c4442d6f33",
    strip_prefix = "rules_dart-84052073c0c6fdb60013b65358d887bb808c6478",
    url = "https://github.com/dart-lang/rules_dart/archive/84052073c0c6fdb60013b65358d887bb808c6478.tar.gz",
)

load("@io_bazel_rules_dart//dart/build_rules:repositories.bzl", "dart_repositories")

dart_repositories()

http_archive(
    name = "com_github_google_webdriver_dart",
    sha256 = "98ad4f57102823b770b0ed8a51097ede621fe14bf2b106a0b581167252f84ab7",
    strip_prefix = "webdriver.dart-cac534aaa9e80c65ebf08fa87f16f01d8f0d25a1",
    url = "https://github.com/google/webdriver.dart/archive/cac534aaa9e80c65ebf08fa87f16f01d8f0d25a1.tar.gz",
)

load("@com_github_google_webdriver_dart//:repositories.bzl", "webdriver_dart_repositories")

webdriver_dart_repositories()
