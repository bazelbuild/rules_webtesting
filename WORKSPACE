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

load("//web:repositories.bzl", "web_test_repositories")

web_test_repositories(
    go = True,
    java = True,
    python = True,
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

load("//web/internal:platform_http_file.bzl", "platform_http_file")

platform_http_file(
    name = "org_chromium_chromedriver",
    amd64_sha256 = "0c01b05276da98f203dc7eb4236c2ee7fe799b432734e088549bd0aadc71958e",
    amd64_url = "http://chromedriver.storage.googleapis.com/2.24/chromedriver_linux64.zip",
    macos_sha256 = "d4f6e13d88ecf20735138f16ab1545e855a42bce41bebe73667a028871777790",
    macos_url = "http://chromedriver.storage.googleapis.com/2.24/chromedriver_mac64.zip",
)

platform_http_file(
    name = "com_google_chrome",
    amd64_sha256 = "6e26d74fd814c38cd419d1ffe87b3e81ad6cfe453e27c7a672ab3c452968e71d",
    amd64_url = "https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/53.0.2785.116/precise64/chrome-precise64.zip",
    macos_sha256 = "84b3cf4f7a9f85fa90dda837b1e38820c83c383fcb6346bbec6448d5128dd7f9",
    macos_url = "https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/53.0.2785.116/mac64/chrome-mac.zip",
)

platform_http_file(
    name = "org_mozilla_firefox",
    amd64_sha256 = "95884070af8870a550ef70600793b6e6d5207f34af24f8b437b6c67b095e5517",
    amd64_url = "https://ftp.mozilla.org/pub/firefox/releases/49.0/firefox-49.0.linux-x86_64.sdk.tar.bz2",
    macos_sha256 = "c068696c69af2da2b916e33e93755f7dda478fa6e9d17a60643cf2009bbaf8e2",
    macos_url = "https://ftp.mozilla.org/pub/firefox/releases/49.0/firefox-49.0.mac-x86_64.sdk.tar.bz2",
)

platform_http_file(
    name = "org_mozilla_geckodriver",
    amd64_sha256 = "dee64571aefb5ef0279df7358d5f74fdf19a316adbab13c67e3c2d2c14da9e97",
    amd64_url = "https://github.com/mozilla/geckodriver/releases/download/v0.10.0/geckodriver-v0.10.0-linux64.tar.gz",
    macos_sha256 = "acb05a7671948167e6c1b6930f32ea71dcaa2c12b2c2963e829c7b232f9125d0",
    macos_url = "https://github.com/mozilla/geckodriver/releases/download/v0.10.0/geckodriver-v0.10.0-macos.tar.gz",
)

platform_http_file(
    name = "org_phantomjs",
    amd64_sha256 = "86dd9a4bf4aee45f1a84c9f61cf1947c1d6dce9b9e8d2a907105da7852460d2f",
    amd64_url = "http://bazel-mirror.storage.googleapis.com/bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2",
    macos_sha256 = "538cf488219ab27e309eafc629e2bcee9976990fe90b1ec334f541779150f8c1",
    macos_url = "http://bazel-mirror.storage.googleapis.com/bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-macosx.zip",
)
