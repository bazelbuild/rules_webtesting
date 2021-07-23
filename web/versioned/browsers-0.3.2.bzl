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
        amd64_sha256 = "6eb18a5a3f77b190fa0bb48bcda4694d26731703ac3ee56499f72f820fe10ef1",
        amd64_urls = [
            "https://saucelabs.com/downloads/sc-4.5.4-linux.tar.gz",
        ],
        macos_sha256 = "7dd691a46a57c7c39f527688abd4825531d25a8a1c5b074f684783e397529ba6",
        macos_urls = [
            "https://saucelabs.com/downloads/sc-4.5.4-osx.zip",
        ],
        windows_sha256 =
            "4b2baaeb32624aa4e60ea4a2ca51f7c5656d476ba29f650a5dabb0faaf6cb793",
        windows_urls = [
            "https://saucelabs.com/downloads/sc-4.5.4-win32.zip",
        ],
    )

# To update Chromium, do the following:
# Step 1: Go to https://omahaproxy.appspot.com/
# Step 2: Look for branch_base_position of current stable releases
# Step 3: Go to https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/ etc to verify presence of that branch release for that platform.
#         If no results, delete the last digit to broaden your search til you find a result.
# Step 4: Verify both Chromium and ChromeDriver are released at that version.
# Step 5: Update the URL to the new release.
def org_chromium_chromedriver():
    platform_http_file(
        name = "org_chromium_chromedriver",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        amd64_sha256 =
            "c8b8be2fc6835bd3003c16d73b9574242e215e81e9b3e01d6fed457988d052f4",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/870763/chromedriver_linux64.zip",
        ],
        macos_sha256 =
            "aa0124085146556d5d32ad172670e5dcef79b7429380112ad02898047ba7a8b7",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/870776/chromedriver_mac64.zip",
        ],
        windows_sha256 =
            "038624e31c327c40df979d699e7c1bba0f322025277f9c875266258169a56faa",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/870788/chromedriver_win32.zip",
        ],
    )
def org_chromium_chromium():
    platform_http_file(
        name = "org_chromium_chromium",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        amd64_sha256 =
            "3a2ae26b7cc56018ea3435bbe22470a82c26340aac72330d6a87555bc3946ab1",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/870763/chrome-linux.zip",
        ],
        macos_sha256 =
            "39118c96db1b3fdb0129f434912a329c5ca07d3a1c6c6cda673d3383d83e2f9a",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/870776/chrome-mac.zip",
        ],
        windows_sha256 =
            "c0ef527ab7e4776b43da164b96969350cc87f1d18de2f6dfc6b74781092fcce5",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/870788/chrome-win.zip",
        ],
    )

def org_mozilla_firefox():
    platform_http_file(
        name = "org_mozilla_firefox",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "284f58b5ee75daec5eaf8c994fe2c8b14aff6c65331e5deeaed6ba650673357c",
        amd64_urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/68.0.2/linux-x86_64/en-US/firefox-68.0.2.tar.bz2",
        ],
        macos_sha256 =
            "173440ca6147c6e1eebbe36f332da2c4347e37269152ad55c431f6b0d7078862",
        macos_urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/68.0.2/mac/en-US/Firefox%2068.0.2.dmg",
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
            "4739ef8f8af5d89bd4a8015788b4dc45c2f5f16b2fdc001254c9a92fe7261947",
        macos_urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.26.0/geckodriver-v0.26.0-macos.tar.gz",
        ],
    )
