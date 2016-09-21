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

http_file(
    name = "org_chromium_chromedriver_linux",
    sha256 = "d9bae7125a772742adc249dcd476f52e766928b83258a510d64a779744b7abfc",
    url = "http://chromedriver.storage.googleapis.com/2.23/chromedriver_linux64.zip",
)

http_file(
    name = "com_google_chrome_linux",
    sha256 = "e3902a68d176f1057579e87fb20389637491d64cef00975608e9d784e2e631e4",
    url = "https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/52.0.2743.116/precise64/chrome-precise64.zip",
)

http_file(
    name = "org_chromium_chromedriver_mac",
    sha256 = "47a8caec6ce251f2dbaa9005e4dc783cb1fa6c09ecd76afafa41eab540a32e86",
    url = "http://chromedriver.storage.googleapis.com/2.23/chromedriver_mac64.zip",
)

http_file(
    name = "com_google_chrome_mac",
    sha256 = "c4a1b0e4d0cf14486e1bca7c3cc210efdbbd5f229dd50613f76e4e36649d8491",
    url = "https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/52.0.2743.116/mac64/chrome-mac.zip",
)

http_file(
    name = "org_mozilla_firefox_linux",
    sha256 = "95884070af8870a550ef70600793b6e6d5207f34af24f8b437b6c67b095e5517",
    url = "https://ftp.mozilla.org/pub/firefox/releases/49.0/firefox-49.0.linux-x86_64.sdk.tar.bz2",
)

http_file(
    name = "org_mozilla_firefox_mac",
    sha256 = "50f27d0fe0eb4c3ba55f4447076021db491463fe32070164394bb836766d1968",
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
