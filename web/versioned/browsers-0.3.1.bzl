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
"""Sets up repositories for use by rules_webtesting at version 0.3.1."""

load("//web:web.bzl", "platform_archive")

def browser_repositories(firefox = False, chromium = False, sauce = False):
    """Sets up repositories for browsers defined in //browsers/....

    Args:
        firefox: Configure repositories for //browsers:firefox-native.
        chromium: Configure repositories for //browsers:chromium-native.
        sauce: Configure repositories for //browser/sauce:chrome-win10-connect.
    """
    if chromium:
        org_chromium_chromedriver()
        org_chromium_chromium()
    if firefox:
        org_mozilla_firefox()
        org_mozilla_geckodriver()
    if sauce:
        com_saucelabs_sauce_connect()

def com_saucelabs_sauce_connect():
    platform_archive(
        name = "com_saucelabs_sauce_connect_linux_x64",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        sha256 = "0de7fcbcb03ad400e886f2c4b34661eda55808e69c7bc4db6aa6aff85e4edb15",
        urls = [
            "https://saucelabs.com/downloads/sc-4.5.3-linux.tar.gz",
        ],
        named_files = {
            "SAUCE_CONNECT": "bin/sc",
        },
    )

    platform_archive(
        name = "com_saucelabs_sauce_connect_macos_x64",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        sha256 = "838d869fbf96ba6595fda2fa40008326337d419e1891a43fee826b995515d4bf",
        urls = [
            "https://saucelabs.com/downloads/sc-4.5.3-osx.zip",
        ],
        named_files = {
            "SAUCE_CONNECT": "bin/sc",
        },
    )

    platform_archive(
        name = "com_saucelabs_sauce_connect_windows_x64",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        sha256 =
            "5fbab3c5ade586049c204698c5e23d550ef8ac4c8fb854210da8d4bcd5f224d5",
        urls = [
            "https://saucelabs.com/downloads/sc-4.5.3-win32.zip",
        ],
        named_files = {
            "SAUCE_CONNECT": "bin/sc.exe",
        },
    )

def org_chromium_chromedriver():
    platform_archive(
        name = "org_chromium_chromedriver_linux_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 =
            "d55d3141a6d9dbff3db6289ff6c6301e1d4ea1248b1cc901376ba0520dfe37b9",
        urls = [
            "https://chromedriver.storage.googleapis.com/72.0.3626.69/chromedriver_linux64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_linux64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_macos_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 =
            "eab0cc3deb77966ed1b1c6569a33f26ee316de7e2063d2200422f7be3667009b",
        urls = [
            "https://chromedriver.storage.googleapis.com/72.0.3626.69/chromedriver_mac64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_mac64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_windows_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 =
            "9138bfd672f9856f1776eb3ac8ba2e54f1efba475a7019854825b3887781545a",
        urls = [
            "https://chromedriver.storage.googleapis.com/72.0.3626.69/chromedriver_win32.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_win32/chromedriver.exe",
        },
    )

def org_chromium_chromium():
    platform_archive(
        name = "org_chromium_chromium_linux_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 =
            "9bb83c07e5b67a6d032e6b9e22500f11e34f86551971f4892dfa68e3544ac39a",
        urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/612434/chrome-linux.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-linux/chrome",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_macos_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 =
            "ab1a75f0d918a0e266f85e43517db5bd701f34544377e3d6aa89f035b508667d",
        urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/612398/chrome-mac.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_windows_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 =
            "8d4e9caa29057ca62635da47ece2774de52ffbc36e6a8894e09e543424ff04c4",
        urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/612432/chrome-win.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-win/chrome.exe",
        },
    )

def org_mozilla_firefox():
    platform_archive(
        name = "org_mozilla_firefox_linux_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 =
            "a89aae224b872d1b5e17ab213ca04a711ae72061828685f999ea2f52784acb56",
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/65.0.1/linux-x86_64/en-US/firefox-65.0.1.tar.bz2",
        ],
        named_files = {
            "FIREFOX": "firefox/firefox",
        },
    )

    platform_archive(
        name = "org_mozilla_firefox_macos_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 =
            "cbbf59f2fad5968db1dfde3d0692b3b08a10f1f283b7dc307f6e4f864a64fae2",
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/65.0.1/mac/en-US/Firefox%2065.0.1.dmg",
        ],
        named_files = {
            "FIREFOX": "Firefox.app/Contents/MacOS/firefox",
        },
    )

def org_mozilla_geckodriver():
    platform_archive(
        name = "org_mozilla_geckodriver_linux_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 =
            "03be3d3b16b57e0f3e7e8ba7c1e4bf090620c147e6804f6c6f3203864f5e3784",
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.24.0/geckodriver-v0.24.0-linux64.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )

    platform_archive(
        name = "org_mozilla_geckodriver_macos_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 =
            "6553195cd6f449e2b90b0bdfe174c6c3337ed571ac6d57a0db028ac5f306cca9",
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.24.0/geckodriver-v0.24.0-macos.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )
