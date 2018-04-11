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

rules_go_version = "0.10.3"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "feba3278c13cde8d67e341a837f69a029f698d7a27ddbb2a202be7a10b22142a",
    urls = [
        "http://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/{v}/rules_go-{v}.tar.gz".format(v = rules_go_version),
        "https://github.com/bazelbuild/rules_go/releases/download/{v}/rules_go-{v}.tar.gz".format(v = rules_go_version),
    ],
)

gazelle_version = "0.10.1"

http_archive(
    name = "bazel_gazelle",
    sha256 = "d03625db67e9fb0905bbd206fa97e32ae9da894fe234a493e7517fd25faec914",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/{v}/bazel-gazelle-{v}.tar.gz".format(v = gazelle_version),
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//web:repositories.bzl", "browser_repositories", "web_test_repositories")

web_test_repositories()

browser_repositories(
    chromium = True,
    firefox = True,
)

rules_scala_version = "55c5bd2c4af311008bcfa0f989af39026ed567fe"

http_archive(
    name = "io_bazel_rules_scala",
    sha256 = "d45b5d621f216eee004e1aaed0f9f9df43c75f55c81fa742b27fc6969d2cbc2b",
    strip_prefix = "rules_scala-{v}".format(v = rules_scala_version),
    type = "zip",
    url = "https://github.com/bazelbuild/rules_scala/archive/{v}.zip".format(v = rules_scala_version),
)

load("@io_bazel_rules_scala//scala:scala.bzl", "scala_repositories")

scala_repositories()

load("@io_bazel_rules_scala//scala:toolchains.bzl", "scala_register_toolchains")

scala_register_toolchains()
