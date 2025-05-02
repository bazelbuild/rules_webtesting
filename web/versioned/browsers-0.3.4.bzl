# Copyright 2023 Google LLC
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
"""Sets up repositories for use by rules_webtesting at version 0.3.4."""

load("//web:web.bzl", "platform_archive")

def browser_repositories():
    """Sets up repositories for browsers defined in //browsers/.... """
    org_chromium_chromedriver()
    org_chromium_chromium()
    org_mozilla_firefox()
    org_mozilla_geckodriver()
    com_saucelabs_sauce_connect()

def com_saucelabs_sauce_connect():
    platform_archive(
        name = "com_saucelabs_sauce_connect_linux_x64",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        sha256 = "6eb18a5a3f77b190fa0bb48bcda4694d26731703ac3ee56499f72f820fe10ef1",
        urls = [
            "https://saucelabs.com/downloads/sc-4.5.4-linux.tar.gz",
        ],
        named_files = {
            "SAUCE_CONNECT": "sc-4.5.4-linux/bin/sc",
        },
    )

    platform_archive(
        name = "com_saucelabs_sauce_connect_macos_x64",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        sha256 = "7dd691a46a57c7c39f527688abd4825531d25a8a1c5b074f684783e397529ba6",
        urls = [
            "https://saucelabs.com/downloads/sc-4.5.4-osx.zip",
        ],
        named_files = {
            "SAUCE_CONNECT": "sc-4.5.4-osx/bin/sc",
        },
    )

    platform_archive(
        name = "com_saucelabs_sauce_connect_windows_x64",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        sha256 =
            "4b2baaeb32624aa4e60ea4a2ca51f7c5656d476ba29f650a5dabb0faaf6cb793",
        urls = [
            "https://saucelabs.com/downloads/sc-4.5.4-win32.zip",
        ],
        named_files = {
            "SAUCE_CONNECT": "sc-4.5.4-win32/bin/sc.exe",
        },
    )

# To update Chromium, do the following:
# Step 1: Go to https://chromiumdash.appspot.com/
# Step 2: Look for branch_base_position of current stable releases
# Step 3: Go to https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/ etc to verify presence of that branch release for that platform.
#         If no results, delete the last digit to broaden your search til you find a result.
# Step 4: Verify both Chromium and ChromeDriver are released at that version.
# Step 5: Update the URL to the new release.
def org_chromium_chromedriver():
    platform_archive(
        name = "org_chromium_chromedriver_linux_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "8813b9fee041c1033d332a309c02a4d04658999fada4771c5901966b3558bda4",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Linux_x64/1250580/chromedriver_linux64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_linux64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_macos_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "6144656cbc4f73628adae24581a128d109711c1cecdb88e57fd11e110d7769cd",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac/1250665/chromedriver_mac64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_mac64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_macos_arm64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "07ac40a4a5cca634aab0b178b8118760bbcffe46e6e59429a1dd8e4c59f02d3b",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/1250665/chromedriver_mac64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_mac64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_windows_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "1b5f3c68490e9defc97d876ba8fe5a00b723636be6b7d33dbe070a55be5f7126",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Win/1250665/chromedriver_win32.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_win32/chromedriver.exe",
        },
    )

def org_chromium_chromium():
    platform_archive(
        name = "org_chromium_chromium_linux_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "b2dba543e76293d9718462eccea482fde606f16991775a73bdfbfcd32dd6a339",
        # 122.0.6261.90
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Linux_x64/1250580/chrome-linux.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-linux/chrome",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_macos_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "7a951016dbb3716c99a17eea35b4d6ebcd953d40b944a8f02d9a76de1295deda",
        # 123.0.6262.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac/1250665/chrome-mac.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_macos_arm64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "cf0dcbfdddd0bcdc08891492b21c4b11ed3b27c17922e7ead1f702a22a2e6d28",
        # 123.0.6262.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/1250665/chrome-mac.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_windows_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "35df630d8b3acceb860bcf28c5bcb854d63aab1a3224353e21808f314c1475da",
        # 123.0.6262.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Win/1250665/chrome-win.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-win/chrome.exe",
        },
    )

def org_mozilla_firefox():
    platform_archive(
        name = "org_mozilla_firefox_linux_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "3d0f74790fe6ff5e38324222ab0c47e10edb31970ed67c6dd7a1c84e7017d1a5",
        # Firefox v97.0
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/97.0/linux-x86_64/en-US/firefox-97.0.tar.bz2",
            "https://storage.googleapis.com/dev-infra-mirror/firefox/97.0/linux_x64/browser-bin.tar.bz2",
        ],
        named_files = {
            "FIREFOX": "firefox/firefox",
        },
    )

    platform_archive(
        name = "org_mozilla_firefox_macos_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "c06c4e58179acaf55d05c3be41d0d4cdd68f811a75322a39557d91121aa2ef74",
        # Firefox v97.0
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/97.0/mac/en-US/Firefox%2097.0.dmg",
            "https://storage.googleapis.com/dev-infra-mirror/firefox/97.0/mac_x64/browser-bin.dmg",
        ],
        named_files = {
            "FIREFOX": "Firefox.app/Contents/MacOS/firefox",
        },
    )

    platform_archive(
        # Firefox has a launcher that conditionally starts x64/arm64. This means that the
        # x64 and arm64 repositories download the same binaries. We preserve separate
        # repositories to allow for dedicated ARM/x64 binaries if needed in the future.
        name = "org_mozilla_firefox_macos_arm64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "c06c4e58179acaf55d05c3be41d0d4cdd68f811a75322a39557d91121aa2ef74",
        # Firefox v97.0
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/97.0/mac/en-US/Firefox%2097.0.dmg",
            "https://storage.googleapis.com/dev-infra-mirror/firefox/97.0/mac_x64/browser-bin.dmg",
        ],
        named_files = {
            "FIREFOX": "Firefox.app/Contents/MacOS/firefox",
        },
    )

def org_mozilla_geckodriver():
    platform_archive(
        name = "org_mozilla_geckodriver_linux_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "7fdd8007d22a6f44caa6929a3d74bbd6a00984d88be50255153671bd201e5493",
        # Geckodriver v0.31.0
        urls = ["https://github.com/mozilla/geckodriver/releases/download/v0.31.0/geckodriver-v0.31.0-linux64.tar.gz"],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )

    platform_archive(
        name = "org_mozilla_geckodriver_macos_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "4da5c6effe987e0c9049c69c7018e70a9d79f3c6119657def2cc0c3419f885e6",
        # Geckodriver v0.31.0
        urls = ["https://github.com/mozilla/geckodriver/releases/download/v0.31.0/geckodriver-v0.31.0-macos.tar.gz"],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )

    platform_archive(
        name = "org_mozilla_geckodriver_macos_arm64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "bfd3974b313be378087f4e7bc4c90128e67dc042647181b4c4ac302b1b88de7f",
        # Geckodriver v0.31.0
        urls = ["https://github.com/mozilla/geckodriver/releases/download/v0.31.0/geckodriver-v0.31.0-macos-aarch64.tar.gz"],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )
