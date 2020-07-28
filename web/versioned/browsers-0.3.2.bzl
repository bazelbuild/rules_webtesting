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

def org_chromium_chromedriver():
    platform_http_file(
        name = "org_chromium_chromedriver",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        amd64_sha256 =
            "0ead02145854b60a3317b59031205b362fb4cfdb680fef20e95c89582e6e38be",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/664981/chromedriver_linux64.zip",
        ],
        macos_sha256 =
            "8dd159e27b13b16262afa6993b15321e736c3b484da363c0e03bb050d72522c9",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/665002/chromedriver_mac64.zip",
        ],
        windows_sha256 =
            "1cc881364974102182257a5c5c2b9cfed513689dee28924ca44df082bdf9fd60",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/664999/chromedriver_win32.zip",
        ],
    )

def org_chromium_chromium():
    platform_http_file(
        name = "org_chromium_chromium",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        amd64_sha256 =
            "b1e30c4dec8a451f8fe10d1f2d3c71e491d0333425f32247fe5c80a0a354303d",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/664981/chrome-linux.zip",
        ],
        macos_sha256 =
            "7c0ba93616f44a421330b1c1262e8899fbdf7916bed8b04c775e0426f6f35ec6",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/665002/chrome-mac.zip",
        ],
        windows_sha256 =
            "f2facd0066270078d0e8999e684595274c359cac3946299a1ceedba2a5de1c63",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/664999/chrome-win.zip",
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
