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
"""Defines external repositories needed by rules_webtesting."""

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
    omit_com_github_gorilla_mux: Do not install Gorilla MUX. Gorilla
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
        sha256="a32c13a36c58cb321136231ae8b67b0c6ad3c5f462e65eb6771f59c44b44ccba",
        strip_prefix="mux-757bef944d0f21880861c2dd9c871ca543023cba",
        url="https://github.com/gorilla/mux/archive/757bef944d0f21880861c2dd9c871ca543023cba.tar.gz"
    )

  if java:
    if not omit_org_seleniumhq_java:
      native.new_http_archive(
          name="org_seleniumhq_java",
          build_file=str(Label("//build_files:selenium_java.BUILD")),
          sha256="0001f86f575d27a1886be3b383e198bcda621330d70c21e0cd0258170c2bc514",
          url="http://selenium-release.storage.googleapis.com/3.0/selenium-java-3.0.1.zip"
      )

    if not omit_com_google_code_findbugs_jsr305:
      native.maven_jar(
          name="com_google_code_findbugs_jsr305",
          artifact="com.google.code.findbugs:jsr305:3.0.1",
          sha1="f7be08ec23c21485b9b5a1cf1654c2ec8c58168d")

    if not omit_com_google_guava_guava:
      native.maven_jar(
          name="com_google_guava_guava",
          artifact="com.google.guava:guava:20.0",
          sha1="89507701249388e1ed5ddcf8c41f4ce1be7831ef")

  if go:
    if not omit_com_github_tebeka_selenium:
      native.new_http_archive(
          name="com_github_tebeka_selenium",
          build_file=str(Label("//build_files:selenium_go.BUILD")),
          sha256="84949fd190b82935af672732b02709b9291933d7d2f5916dcae8b4668c0004cc",
          strip_prefix="selenium-0.9.3",
          url="https://github.com/tebeka/selenium/archive/v0.9.3.tar.gz"
      )

  if python:
    if not omit_org_seleniumhq_py:
      native.new_http_archive(
          name="org_seleniumhq_py",
          build_file=str(Label("//build_files:selenium_py.BUILD")),
          sha256="85daad4d09be86bddd4f45579986ac316c1909c3b4653ed471ea4519eb413c8f",
          strip_prefix="selenium-3.0.2/py",
          url="https://pypi.python.org/packages/0c/42/20c235e604bf736bc970c1275a78c4ea28c6453a0934002f95df9c49dad0/selenium-3.0.2.tar.gz"
      )


def browser_repositories(firefox=False, chromium=False):
  """Sets up repositories for browsers defined in //browsers/....

  This should only be used on an experimental basis; projects should define
  their own browsers.

  Args:
    firefox: Configure repositories for //browsers:firefox-native.
    chromium: Configure repositories for //browsers:chromium-native.
  """
  if chromium:
    platform_http_file(
        name="org_chromium_chromedriver",
        amd64_sha256="59e6b1b1656a20334d5731b3c5a7400f92a9c6f5043bb4ab67f1ccf1979ee486",
        amd64_url="http://chromedriver.storage.googleapis.com/2.26/chromedriver_linux64.zip",
        macos_sha256="70aae3812941ed94ad8065bb4a9432861d7d4ebacdd93ee47bb2c7c57c7e841e",
        macos_url="http://chromedriver.storage.googleapis.com/2.26/chromedriver_mac64.zip"
    )

    # Roughly corresponds to Chrome 55
    platform_http_file(
        name="org_chromium_chromium",
        amd64_sha256="e3c99954d6acce013174053534b72f47f67f18a0d75f79c794daaa8dd2ae8aaf",
        amd64_url="https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/423768/chrome-linux.zip",
        macos_sha256="62aeb7a5c6b8a1b7b31400105bf01295bbd45b0627920b8f99f0cc4ca76927ca",
        macos_url="https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/423758/chrome-mac.zip"
    )

  if firefox:
    platform_http_file(
        name="org_mozilla_firefox",
        amd64_sha256="10533f3db9c819a56f6cd72f9340e05c7e3b116454eb81b0d39ed161955bb48f",
        amd64_url="https://ftp.mozilla.org/pub/firefox/releases/50.1.0/firefox-50.1.0.linux-x86_64.sdk.tar.bz2",
        macos_sha256="5cd449ebedb44b2f882b37e6e5cee1a814bc5ff3c3f86d1a1019b937aa287441",
        macos_url="https://ftp.mozilla.org/pub/firefox/releases/50.1.0/firefox-50.1.0.mac-x86_64.sdk.tar.bz2"
    )

    platform_http_file(
        name="org_mozilla_geckodriver",
        amd64_sha256="ce4aa8b5cf918a6607b50e73996fb909db42fd803855f0ecc9d7183999c3bedc",
        amd64_url="https://github.com/mozilla/geckodriver/releases/download/v0.11.1/geckodriver-v0.11.1-linux64.tar.gz",
        macos_sha256="802cc1a33b8ce6f7c3aeb5116730cb6efc20414959d6f750e74437869d37a150",
        macos_url="https://github.com/mozilla/geckodriver/releases/download/v0.11.1/geckodriver-v0.11.1-macos.tar.gz"
    )
