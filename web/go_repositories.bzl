# Copyright 2019 Google Inc.
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
"""Defines external repositories needed by rules_webtesting."""

load(":repositories.bzl", "should_create_repository")
load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories(**kwargs):
    """Defines external repositories required to use Go webtest and screenshotter APIs.

    This function exists for other Bazel projects to call from their WORKSPACE
    file when depending on rules_webtesting using http_archive. This function
    makes it easy to import these transitive dependencies into the parent
    workspace. This will check to see if a repository has been previously defined
    before defining a new repository.

    Alternatively, individual dependencies may be excluded with an
    "omit_" + name parameter. This is useful for users who want to be rigorous
    about declaring their own direct dependencies, or when another Bazel project
    is depended upon (e.g. rules_closure) that defines the same dependencies as
    this one (e.g. com_google_guava_guava.) Alternatively, a whitelist model may be
    used by calling the individual functions this method references.

    Please note that while these dependencies are defined, they are not actually
    downloaded, unless a target is built that depends on them.

    Args:
        **kwargs: omit_... parameters used to prevent importing specific
          dependencies.
    """
    if should_create_repository("com_github_blang_semver", kwargs):
        com_github_blang_semver()
    if should_create_repository("com_github_tebeka_selenium", kwargs):
        com_github_tebeka_selenium()
    if kwargs.keys():
        print("The following parameters are unknown: " + str(kwargs.keys()))

def go_internal_repositories(**kwargs):
    """Defines external repositories required to build WebTesting rules release.

    Alternatively, individual dependencies may be excluded with an
    "omit_" + name parameter. This is useful for users who want to be rigorous
    about declaring their own direct dependencies, or when another Bazel project
    is depended upon (e.g. rules_closure) that defines the same dependencies as
    this one (e.g. com_google_guava_guava.) Alternatively, a whitelist model may be
    used by calling the individual functions this method references.

    Please note that while these dependencies are defined, they are not actually
    downloaded, unless a target is built that depends on them.

    Args:
        **kwargs: omit_... parameters used to prevent importing specific
          dependencies.
    """
    if should_create_repository("com_github_gorilla_mux", kwargs):
        com_github_gorilla_mux()
    if kwargs.keys():
        print("The following parameters are unknown: " + str(kwargs.keys()))

def com_github_blang_semver():
    go_repository(
        name = "com_github_blang_semver",
        importpath = "github.com/blang/semver",
        sha256 = "dc85076e7c5a7a44e33fc24df320904b95e2fa12c94a3ac758a574dadd54ee53",
        strip_prefix = "semver-3.6.1",
        urls = [
            "https://github.com/blang/semver/archive/v3.6.1.tar.gz",
        ],
    )

def com_github_gorilla_mux():
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sha256 = "92adb9aea022f8b35686b75be50ba1206c4457c2f8a0e2a9d10d8721f35b3f11",
        strip_prefix = "mux-1.7.3",
        urls = [
            "https://github.com/gorilla/mux/archive/v1.7.3.tar.gz",
        ],
    )

def com_github_tebeka_selenium():
    go_repository(
        name = "com_github_tebeka_selenium",
        importpath = "github.com/tebeka/selenium",
        sha256 = "82846f237b742983a293619e712dcf167e3d7998df3239f3443303d9036ad412",
        strip_prefix = "selenium-0.9.9",
        urls = [
            "https://github.com/tebeka/selenium/archive/v0.9.9.tar.gz",
        ],
    )
