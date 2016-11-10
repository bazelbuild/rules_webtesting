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
"""Defines external repositories needs by rules_webtesting."""

load("//web/internal:platform_http_file.bzl", "platform_http_file")


def web_test_repositories(java=False,
                          go=False,
                          python=False,
                          omit_com_github_gorilla_mux=False,
                          omit_org_seleniumhq_java=False,
                          omit_com_google_code_findbugs_jsr305=False,
                          omit_com_google_guava_guava=False,
                          omit_com_github_tebeka_selenium=False,
                          omit_org_seleniumhq_py=False):
  """Configure repositories for Web Test Launcher and for client languages.    
   
  Args:    
    java: Configure Java client-side libraries.    
    go: Configure Go client-side libraries.    
    python: Configure Python client libraries.   
    omit_com_github_gorilla_mux*: Do not install Gorilla MUX. Gorilla    
      MUX is required to compile the test launcher.    
    omit_org_seleniumhq_java: Do not install Java Selenium client bindings.    
      These bindings are only installed if java=True.    
    omit_com_google_code_findbugs_jsr305: Do not install JSR305 annotations    
      library. This library is only installed if java=True.    
    omit_com_google_guava_guava: Do not install Guava libraries. This    
      library is only installed if java=True.    
    omit_com_github_tebeka_selenium: Do not install Go WebDriver client    
      bindings. These binding are only installed if go=True.   
    omit_org_seleniumhq_py: Do not install Python Selenium client bindings.    
      These bindings are only installed if python=True.    
  """
  if not omit_com_github_gorilla_mux:
    native.new_http_archive(
        name="com_github_gorilla_mux",
        build_file=str(Label("//build_files:gorilla_mux.BUILD")),
        url="https://github.com/gorilla/mux/archive/cf79e51a62d8219d52060dfc1b4e810414ba2d15.tar.gz",
        sha256="80077e14b2f0f8f2796b6bfcf5c8e41e148e3c8c45b4c20d1e6856b348d5efb7",
        strip_prefix="mux-cf79e51a62d8219d52060dfc1b4e810414ba2d15")

  if java:
    if not omit_org_seleniumhq_java:
      native.new_http_archive(
          name="org_seleniumhq_java",
          build_file=str(Label("//build_files:selenium_java.BUILD")),
          sha256="a26a449388abd46d1e152771e3641859ac4acee9c0ea24a101ca369048a81ecb",
          url="http://selenium-release.storage.googleapis.com/3.0-beta3/selenium-java-3.0.0-beta3.zip"
      )

    if not omit_com_google_code_findbugs_jsr305:
      native.maven_jar(
          name="com_google_code_findbugs_jsr305",
          artifact="com.google.code.findbugs:jsr305:3.0.1",
          sha1="f7be08ec23c21485b9b5a1cf1654c2ec8c58168d")

    if not omit_com_google_guava_guava:
      native.maven_jar(
          name="com_google_guava_guava",
          artifact="com.google.guava:guava:19.0",
          sha1="6ce200f6b23222af3d8abb6b6459e6c44f4bb0e9")

  if go:
    if not omit_com_github_tebeka_selenium:
      native.new_http_archive(
          name="com_github_tebeka_selenium",
          build_file=str(Label("//build_files:selenium_go.BUILD")),
          url="https://github.com/tebeka/selenium/archive/v0.9.2.tar.gz",
          sha256="c5f21652eda6230ee8bb5f9f02b740fa8d8b22c0cddc832ec666a7654bb0d1a4",
          strip_prefix="selenium-0.9.2")

  if python:
    if not omit_org_seleniumhq_py:
      native.new_http_archive(
          name="org_seleniumhq_py",
          build_file=str(Label("//build_files:selenium_py.BUILD")),
          sha256="0705803349964c7a2a144f1796a5d29905fe2a09931b2bb945ee0cb4deab75d7",
          strip_prefix="selenium-3.0.1/py",
          url="https://pypi.python.org/packages/3a/a3/e4ab60a0229a85f468a36367bc0672a4bca2720f24391eda33704a5f0ad5/selenium-3.0.1.tar.gz"
      )


def browser_repositories(firefox=False, chrome=False, phantomjs=False):
  """Sets up repositories for browsers defined in //browsers/....

  This should only be used on an experimental basis; trojects should define
  their own.browsers.

  Args:
    firefox: Configure repositories for //browsers:firefox-native.
    chrome: Configure repositories for //browsers:chrome-native.
    phantomjs: Configure repositories for //browsers:phantomjs-native.
  """
  if chrome:
    platform_http_file(
        name="org_chromium_chromedriver",
        amd64_sha256="0c01b05276da98f203dc7eb4236c2ee7fe799b432734e088549bd0aadc71958e",
        amd64_url="http://chromedriver.storage.googleapis.com/2.24/chromedriver_linux64.zip",
        macos_sha256="d4f6e13d88ecf20735138f16ab1545e855a42bce41bebe73667a028871777790",
        macos_url="http://chromedriver.storage.googleapis.com/2.24/chromedriver_mac64.zip"
    )

    platform_http_file(
        name="com_google_chrome",
        amd64_sha256="6e26d74fd814c38cd419d1ffe87b3e81ad6cfe453e27c7a672ab3c452968e71d",
        amd64_url="https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/53.0.2785.116/precise64/chrome-precise64.zip",
        macos_sha256="84b3cf4f7a9f85fa90dda837b1e38820c83c383fcb6346bbec6448d5128dd7f9",
        macos_url="https://commondatastorage.googleapis.com/chrome-unsigned/desktop-5c0tCh/53.0.2785.116/mac64/chrome-mac.zip"
    )

  if firefox:
    platform_http_file(
        name="org_mozilla_firefox",
        amd64_sha256="95884070af8870a550ef70600793b6e6d5207f34af24f8b437b6c67b095e5517",
        amd64_url="https://ftp.mozilla.org/pub/firefox/releases/49.0/firefox-49.0.linux-x86_64.sdk.tar.bz2",
        macos_sha256="c068696c69af2da2b916e33e93755f7dda478fa6e9d17a60643cf2009bbaf8e2",
        macos_url="https://ftp.mozilla.org/pub/firefox/releases/49.0/firefox-49.0.mac-x86_64.sdk.tar.bz2"
    )

    platform_http_file(
        name="org_mozilla_geckodriver",
        amd64_sha256="dee64571aefb5ef0279df7358d5f74fdf19a316adbab13c67e3c2d2c14da9e97",
        amd64_url="https://github.com/mozilla/geckodriver/releases/download/v0.10.0/geckodriver-v0.10.0-linux64.tar.gz",
        macos_sha256="acb05a7671948167e6c1b6930f32ea71dcaa2c12b2c2963e829c7b232f9125d0",
        macos_url="https://github.com/mozilla/geckodriver/releases/download/v0.10.0/geckodriver-v0.10.0-macos.tar.gz"
    )

  if phantomjs:
    platform_http_file(
        name="org_phantomjs",
        amd64_sha256="86dd9a4bf4aee45f1a84c9f61cf1947c1d6dce9b9e8d2a907105da7852460d2f",
        amd64_url="http://bazel-mirror.storage.googleapis.com/bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2",
        macos_sha256="538cf488219ab27e309eafc629e2bcee9976990fe90b1ec334f541779150f8c1",
        macos_url="http://bazel-mirror.storage.googleapis.com/bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-macosx.zip"
    )

    native.http_jar(
        name="org_seleniumhq_server",
        sha256="f5ada04a651ba7ec70fcbc68bd4a59342a928ef7dce858ec594a8d5c49576ace",
        url="http://selenium-release.storage.googleapis.com/3.0-beta3/selenium-server-standalone-3.0.0-beta3.jar"
    )
