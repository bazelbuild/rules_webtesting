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
    tag = "0.2.0",
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
    noop_web_test_template = "//web/internal:noop_web_test.sh.template",
    prefix = "",
    web_test_template = "//web/internal:web_test.sh.template",
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

http_file(
    name = "org_chromium_chromedriver_linux",
    sha256 = "0c01b05276da98f203dc7eb4236c2ee7fe799b432734e088549bd0aadc71958e",
    url = "http://chromedriver.storage.googleapis.com/2.24/chromedriver_linux64.zip",
)

http_file(
    name = "com_google_chrome_linux",
    sha256 = "6e26d74fd814c38cd419d1ffe87b3e81ad6cfe453e27c7a672ab3c452968e71d",
    url = "https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/53.0.2785.116/precise64/chrome-precise64.zip",
)

http_file(
    name = "org_chromium_chromedriver_mac",
    sha256 = "d4f6e13d88ecf20735138f16ab1545e855a42bce41bebe73667a028871777790",
    url = "http://chromedriver.storage.googleapis.com/2.24/chromedriver_mac64.zip",
)

http_file(
    name = "com_google_chrome_mac",
    sha256 = "84b3cf4f7a9f85fa90dda837b1e38820c83c383fcb6346bbec6448d5128dd7f9",
    url = "https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/53.0.2785.116/mac64/chrome-mac.zip",
)

http_file(
    name = "org_mozilla_firefox_linux",
    sha256 = "95884070af8870a550ef70600793b6e6d5207f34af24f8b437b6c67b095e5517",
    url = "https://ftp.mozilla.org/pub/firefox/releases/49.0/firefox-49.0.linux-x86_64.sdk.tar.bz2",
)

http_file(
    name = "org_mozilla_firefox_mac",
    sha256 = "c068696c69af2da2b916e33e93755f7dda478fa6e9d17a60643cf2009bbaf8e2",
    url = "https://ftp.mozilla.org/pub/firefox/releases/49.0/firefox-49.0.mac-x86_64.sdk.tar.bz2",
)

http_file(
    name = "org_mozilla_geckodriver_linux",
    sha256 = "dee64571aefb5ef0279df7358d5f74fdf19a316adbab13c67e3c2d2c14da9e97",
    url = "https://github.com/mozilla/geckodriver/releases/download/v0.10.0/geckodriver-v0.10.0-linux64.tar.gz",
)

http_file(
    name = "org_mozilla_geckodriver_mac",
    sha256 = "acb05a7671948167e6c1b6930f32ea71dcaa2c12b2c2963e829c7b232f9125d0",
    url = "https://github.com/mozilla/geckodriver/releases/download/v0.10.0/geckodriver-v0.10.0-macos.tar.gz",
)

http_file(
    name = "org_phantomjs_linux",
    sha256 = "86dd9a4bf4aee45f1a84c9f61cf1947c1d6dce9b9e8d2a907105da7852460d2f",
    url = "https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2",
)

http_file(
    name = "org_phantomjs_mac",
    sha256 = "538cf488219ab27e309eafc629e2bcee9976990fe90b1ec334f541779150f8c1",
    url = "https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-macosx.zip",
)
