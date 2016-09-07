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
workspace(name = "io_bazel_rules_web")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.1.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

load("//web:repositories.bzl", "web_test_repositories")

web_test_repositories(
    default_config = "//web:default_config",
    go = True,
    java = True,
    launcher = "//go/launcher:main",
    merger = "//go/metadata:merger",
    prefix = "",
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

git_repository(
    name = "io_bazel_rules_sass",
    remote = "https://github.com/bazelbuild/rules_sass.git",
    tag = "0.0.1",
)

load("@io_bazel_rules_sass//sass:sass.bzl", "sass_repositories")

sass_repositories()

git_repository(
    name = "io_bazel_skydoc",
    commit = "c57ff682364dbb1ae808b769f9e3add77cdbfad1",
    remote = "https://github.com/bazelbuild/skydoc.git",
)

load("@io_bazel_skydoc//skylark:skylark.bzl", "skydoc_repositories")

skydoc_repositories()

new_http_archive(
    name = "org_chromium_chromedriver",
    build_file = "//build_files:BUILD.chromedriver",
    sha256 = "d9bae7125a772742adc249dcd476f52e766928b83258a510d64a779744b7abfc",
    url = "http://chromedriver.storage.googleapis.com/2.23/chromedriver_linux64.zip",
)

new_http_archive(
    name = "com_google_chrome",
    build_file = "//build_files:BUILD.chrome",
    sha256 = "e3902a68d176f1057579e87fb20389637491d64cef00975608e9d784e2e631e4",
    url = "https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/52.0.2743.116/precise64/chrome-precise64.zip",
)

new_http_archive(
    name = "org_mozilla_firefox",
    build_file = "//build_files:BUILD.firefox",
    sha256 = "af171a6f605703a2322cca95c856c3c8111fd4ba59219efda01f867c2cc34608",
    url = "https://ftp.mozilla.org/pub/firefox/releases/48.0.2/firefox-48.0.2.linux-x86_64.sdk.tar.bz2",
)

new_http_archive(
    name = "org_mozilla_geckodriver",
    build_file = "//build_files:BUILD.geckodriver",
    sha256 = "dee64571aefb5ef0279df7358d5f74fdf19a316adbab13c67e3c2d2c14da9e97",
    url = "https://github.com/mozilla/geckodriver/releases/download/v0.10.0/geckodriver-v0.10.0-linux64.tar.gz",
)
