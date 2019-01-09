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
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def web_test_repositories(**kwargs):
    """Defines external repositories required by WebTesting Rules.

    This function exists for other Bazel projects to call from their WORKSPACE
    file when depending on rules_webtesting using http_archive. This function
    makes it easy to import these transitive dependencies into the parent
    workspace. This will check to see if a repository has been previously defined
    before defining a new repository.

    Alternatively, individual dependencies may be excluded with an
    "omit_" + name parameter. This is useful for users who want to be rigorous
    about declaring their own direct dependencies, or when another Bazel project
    is depended upon (e.g. rules_closure) that defines the same dependencies as
    this one (e.g. com_google_guava.) Alternatively, a whitelist model may be
    used by calling the individual functions this method references.

    Please note that while these dependencies are defined, they are not actually
    downloaded, unless a target is built that depends on them.

    Args:
        **kwargs: omit_... parameters used to prevent importing specific
          dependencies.
    """
    if should_create_repository("bazel_skylib", kwargs):
        bazel_skylib()
    if kwargs.keys():
        print("The following parameters are unknown: " + str(kwargs.keys()))

def should_create_repository(name, args):
    """Returns whether the name repository should be created.
    This allows creation of a repository to be disabled by either an
    "omit_" _+ name parameter or by previously defining a rule for the repository.
    The args dict will be mutated to remove "omit_" + name.
    Args:
        name: The name of the repository that should be checked.
        args: A dictionary that contains "omit_...": bool pairs.
    Returns:
        boolean indicating whether the repository should be created.
    """
    key = "omit_" + name
    if key in args:
        val = args.pop(key)
        if val:
            return False
    if native.existing_rule(name):
        return False
    return True

def browser_repositories(firefox = False, chromium = False, sauce = False):
    """Sets up repositories for browsers defined in //browsers/....

    This should only be used on an experimental basis; projects should define
    their own browsers.

    Args:
        firefox: Configure repositories for //browsers:firefox-native.
        chromium: Configure repositories for //browsers:chromium-native.
        sauce: Configure repositories for //browser/sauce:chrome-win10.
    """
    if chromium:
        org_chromium_chromedriver()
        org_chromium_chromium()
    if firefox:
        org_mozilla_firefox()
        org_mozilla_geckodriver()
    if sauce:
        com_saucelabs_sauce_connect()

def bazel_skylib():
    http_archive(
        name = "bazel_skylib",
        sha256 = "68ef2998919a92c2c9553f7a6b00a1d0615b57720a13239c0e51d0ded5aa452a",
        strip_prefix = "bazel-skylib-8cecf885c8bf4c51e82fd6b50b9dd68d2c98f757",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/archive/8cecf885c8bf4c51e82fd6b50b9dd68d2c98f757.tar.gz",
            "https://github.com/bazelbuild/bazel-skylib/archive/8cecf885c8bf4c51e82fd6b50b9dd68d2c98f757.tar.gz",
        ],
    )

def com_saucelabs_sauce_connect():
    platform_http_file(
        name = "com_saucelabs_sauce_connect",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        amd64_sha256 = "dd53f2cdcec489fbc2443942b853b51bf44af39f230600573119cdd315ddee52",
        amd64_urls = [
            "https://saucelabs.com/downloads/sc-4.5.1-linux.tar.gz",
        ],
        macos_sha256 = "920ae7bd5657bccdcd27bb596593588654a2820486043e9a12c9062700697e66",
        macos_urls = [
            "https://saucelabs.com/downloads/sc-4.5.1-osx.zip",
        ],
        windows_sha256 =
            "ec11b4ee029c9f0cba316820995df6ab5a4f394053102e1871b9f9589d0a9eb5",
        windows_urls = [
            "https://saucelabs.com/downloads/sc-4.4.12-win32.zip",
        ],
    )

def org_chromium_chromedriver():
    platform_http_file(
        name = "org_chromium_chromedriver",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        amd64_sha256 =
            "d4a5eec0a3b7fec9bcb71353233dde38630e51b29fa7b218cdd196e2e4487da7",
        amd64_urls = [
            "https://chromedriver.storage.googleapis.com/2.45/chromedriver_linux64.zip",
        ],
        macos_sha256 =
            "aa0f416a48e20185da62525869c1f98f994bf99f241d6ce2eb1af6ceb517c425",
        macos_urls = [
            "https://chromedriver.storage.googleapis.com/2.45/chromedriver_mac64.zip",
        ],
        windows_sha256 =
            "8f3373a260a524410e25ea643ecb0175ed49a078088c7ab4d88db323ee19a230",
        windows_urls = [
            "https://chromedriver.storage.googleapis.com/2.45/chromedriver_win32.zip",
        ],
    )

def org_chromium_chromium():
    platform_http_file(
        name = "org_chromium_chromium",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        amd64_sha256 =
            "a83baae72d70d4c337c7af6fb190ec3fbb082b94af1e16c5f28390dfd7771612",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/587811/chrome-linux.zip",
        ],
        macos_sha256 =
            "2878816a49e2eeaac4f9e689d05668eb51d0e5e0d1c071756d3a4dfcf35cfcd0",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/587811/chrome-mac.zip",
        ],
        windows_sha256 =
            "fe5c6fc33db8d69c3270a22d6569fe5e9f2afb24cb0b7a0a451793c521ab6eef",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win_x64/587811/chrome-win32.zip",
        ],
    )

def org_mozilla_firefox():
    platform_http_file(
        name = "org_mozilla_firefox",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "d4c696d77823bb3a3ea24d67d0f1075899fbaa4a0893d069f606015708ad1fca",
        amd64_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/63.0.1/linux-x86_64/en-US/firefox-63.0.1.tar.bz2",
            "https://ftp.mozilla.org/pub/firefox/releases/63.0.1/linux-x86_64/en-US/firefox-63.0.1.tar.bz2",
        ],
        macos_sha256 =
            "73e810389dcbf3b0ac74d0997121d498789d5cd205da16175c4ca5fca2ca8f79",
        macos_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/63.0.1/mac/en-US/Firefox%2063.0.1.dmg",
            "https://ftp.mozilla.org/pub/firefox/releases/63.0.1/mac/en-US/Firefox%2063.0.1.dmg",
        ],
    )

def org_mozilla_geckodriver():
    platform_http_file(
        name = "org_mozilla_geckodriver",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "2abf02cb69b48f2ba65ea344b752ff547e5431659aad80b03bf68cdb4f8df14b",
        amd64_urls = [
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-linux64.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-linux64.tar.gz",
        ],
        macos_sha256 =
            "006e206cc4c93ad9ef857aa5b7efc5a9084fa239381a7afaaa4acbba6f00bac9",
        macos_urls = [
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-macos.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-macos.tar.gz",
        ],
    )
