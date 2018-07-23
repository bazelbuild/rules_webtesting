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

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_skylib",
    sha256 = "b5f6abe419da897b7901f90cbab08af958b97a8f3575b0d3dd062ac7ce78541f",
    strip_prefix = "bazel-skylib-0.5.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/archive/0.5.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/archive/0.5.0.tar.gz",
    ],
)

# NOTE: URLs are mirrored by an asynchronous review process. They must
#       be greppable for that to happen. It's OK to submit broken mirror
#       URLs, so long as they're correctly formatted. Bazel's downloader
#       has fast failover.

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "ba79c532ac400cefd1859cbc8a9829346aa69e3b99482cd5a54432092cbc3933",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/0.13.0/rules_go-0.13.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.13.0/rules_go-0.13.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "bc653d3e058964a5a26dcad02b6c72d7d63e6bb88d94704990b908a1445b8758",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/0.13.0/bazel-gazelle-0.13.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.13.0/bazel-gazelle-0.13.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//web:repositories.bzl", "browser_repositories", "web_test_repositories")

web_test_repositories()

browser_repositories(
    chromium = True,
    firefox = True,
    sauce = True,
)

http_archive(
    name = "io_bazel_rules_scala",
    sha256 = "0ac3ef277205d10ca353a5760da9710c8dbf2fca5bd1d2b6143a073ed2a1a50f",
    strip_prefix = "rules_scala-55c5bd2c4af311008bcfa0f989af39026ed567fe",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_scala/archive/55c5bd2c4af311008bcfa0f989af39026ed567fe.tar.gz",
        "https://github.com/bazelbuild/rules_scala/archive/55c5bd2c4af311008bcfa0f989af39026ed567fe.tar.gz",
    ],
)

load("@io_bazel_rules_scala//scala:scala.bzl", "scala_repositories")

scala_repositories()

load("@io_bazel_rules_scala//scala:toolchains.bzl", "scala_register_toolchains")

scala_register_toolchains()
