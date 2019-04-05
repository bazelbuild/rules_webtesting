# Copyright 2019 Google LLC
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

load("//web/internal:platform_http_file.bzl", "platform_http_file")

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
    platform_http_file(
        name = "com_saucelabs_sauce_connect",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        amd64_sha256 = "0de7fcbcb03ad400e886f2c4b34661eda55808e69c7bc4db6aa6aff85e4edb15",
        amd64_urls = [
            "https://saucelabs.com/downloads/sc-4.5.3-linux.tar.gz",
        ],
        macos_sha256 = "838d869fbf96ba6595fda2fa40008326337d419e1891a43fee826b995515d4bf",
        macos_urls = [
            "https://saucelabs.com/downloads/sc-4.5.3-osx.zip",
        ],
        windows_sha256 =
            "5fbab3c5ade586049c204698c5e23d550ef8ac4c8fb854210da8d4bcd5f224d5",
        windows_urls = [
            "https://saucelabs.com/downloads/sc-4.5.3-win32.zip",
        ],
    )

def org_chromium_chromedriver():
    platform_http_file(
        name = "org_chromium_chromedriver",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        amd64_sha256 =
            "258830acdfd1273c47dd06defa92e37944d5379fb74c3bc9cc6dddbfd9b1d26e",
        amd64_urls = [
            "https://chromedriver.storage.googleapis.com/73.0.3683.68/chromedriver_linux64.zip",
        ],
        macos_sha256 =
            "eaaa1b0b7d47b113d228ca99a5d68de52f660ccd9dd78a069df8cd97ff83308a",
        macos_urls = [
            "https://chromedriver.storage.googleapis.com/73.0.3683.68/chromedriver_mac64.zip",
        ],
        windows_sha256 =
            "8a1f37de201d406f6aeb4a21e81311b685356c6d2ac16512008ba5322939bf69",
        windows_urls = [
            "https://chromedriver.storage.googleapis.com/73.0.3683.68/chromedriver_win32.zip",
        ],
    )

def org_chromium_chromium():
    platform_http_file(
        name = "org_chromium_chromium",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        amd64_sha256 =
            "07088ae85de4a0c380ddb0bff7d39a9b30efe537293830378d06e609d8112cfb",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/625894/chrome-linux.zip",
        ],
        macos_sha256 =
            "65433518ba6269113ad8479369b3563c1c37c2765c88239a0af53f5acd25d8bb",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/625854/chrome-mac.zip",
        ],
        windows_sha256 =
            "bc80a8fa401404d1dc2e1017747fefcc3c7c1abf56dc52c042afafa11e689bae",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/625885/chrome-win.zip",
        ],
    )

def org_mozilla_firefox():
    platform_http_file(
        name = "org_mozilla_firefox",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "eda0f0bc2af3a4b19f77d48c7242453477565906337848b51e363291456f41b1",
        amd64_urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/66.0.2/linux-x86_64/en-US/firefox-66.0.2.tar.bz2",
        ],
        macos_sha256 =
            "bc80a8fa401404d1dc2e1017747fefcc3c7c1abf56dc52c042afafa11e689bae",
        macos_urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/66.0.2/mac/en-US/Firefox%2066.0.2.dmg",
        ],
    )

def org_mozilla_geckodriver():
    platform_http_file(
        name = "org_mozilla_geckodriver",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "03be3d3b16b57e0f3e7e8ba7c1e4bf090620c147e6804f6c6f3203864f5e3784",
        amd64_urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.24.0/geckodriver-v0.24.0-linux64.tar.gz",
        ],
        macos_sha256 =
            "6553195cd6f449e2b90b0bdfe174c6c3337ed571ac6d57a0db028ac5f306cca9",
        macos_urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.24.0/geckodriver-v0.24.0-macos.tar.gz",
        ],
    )
