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
"""Sets up repositories for use by rules_webtesting at version 0.3.3."""

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
        sha256 = "6eb18a5a3f77b190fa0bb48bcda4694d26731703ac3ee56499f72f820fe10ef1",
        urls = [
            "https://saucelabs.com/downloads/sc-4.5.4-linux.tar.gz",
        ],
        named_files = {
            "SAUCE_CONNECT": "bin/sc",
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
            "SAUCE_CONNECT": "bin/sc",
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
            "SAUCE_CONNECT": "bin/sc.exe",
        },
    )

# To update Chromium, do the following:
# Step 1: Go to https://omahaproxy.appspot.com/
# Step 2: Look for branch_base_position of current stable releases
# Step 3: Go to https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/ etc to verify presence of that branch release for that platform.
#         If no results, delete the last digit to broaden your search til you find a result.
# Step 4: Verify both Chromium and ChromeDriver are released at that version.
# Step 5: Update the URL to the new release.
def org_chromium_chromedriver():
    platform_archive(
        name = "org_chromium_chromedriver_linux_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "1d2e73a19632031f5de876916e12b497d5b0e3dc83d1ce2fbe8665061adfd114",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Linux_x64/902390/chromedriver_linux64.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chromedriver_linux64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_linux64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_macos_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "36cc50c5194767b043913534f6ec16a7d7a85636b319729a67ffff486b30a5f6",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac/902390/chromedriver_mac64.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chromedriver_mac_x64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_mac64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_macos_arm64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "1f100aacf4bab4b3ac4218ecf654b17d66f2e07dd455f887bb3d9aa8d21862e1",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/902390/chromedriver_mac64.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chromedriver_mac_arm64.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_mac64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_windows_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "48392698f2ba338a0b9192f7c2154058a0b0b926aef0a5ef22aa6706b2bbc7b6",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Win/902390/chromedriver_win32.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chromedriver_win32.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_win32/chromedriver.exe",
        },
    )

def org_chromium_chromium():
    platform_archive(
        name = "org_chromium_chromium_linux_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        # 94.0.4578.0
        sha256 = "673ee08b4cfaff128ef0b4f7517acb6b6b25c9315fc6494ec328ab38aaf952d1",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Linux_x64/902390/chrome-linux.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chrome-linux.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-linux/chrome",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_macos_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "75f6bd26744368cd0fcbbec035766dea82e34def60e938fb48630be6799d46c7",
        # 94.0.4578.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac/902390/chrome-mac.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chrome-mac_x64.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_macos_arm64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "4845ce895d030aeb8bfd877a599f1f07d8c7a77d1e08513e80e60bb0093fca24",
        # 94.0.4578.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/902390/chrome-mac.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chrome-mac_arm64.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_windows_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "8919cd2f8a4676af4acc50d022b6a946a5b21a5fec4e078b0ebb0c8e18f1ce90",
        # 94.0.4578.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Win/902390/chrome-win.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chrome-win.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-win/chrome.exe",
        },
    )

def org_mozilla_firefox():
    platform_archive(
        name = "org_mozilla_firefox_linux_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "998607f028043b3780f296eee03027279ef059acab5b50f9754df2bd69ca42b3",
        # Firefox v90.0.1
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/90.0.1/linux-x86_64/en-US/firefox-90.0.1.tar.bz2",
            "https://storage.googleapis.com/dev-infra-mirror/mozilla/firefox/firefox-90.0.1.tar.bz2",
        ],
        named_files = {
            "FIREFOX": "firefox/firefox",
        },
    )

    platform_archive(
        # Firefox has a launcher that conditionally starts x64/arm64
        name = "org_mozilla_firefox_macos_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "76c1b9c42b52c7e5be4c112a98b7d3762a18841367f778a179679ac0de751f05",
        # Firefox v90.0.1
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/90.0.1/mac/en-US/Firefox%2090.0.1.dmg",
            "https://storage.googleapis.com/dev-infra-mirror/mozilla/firefox/Firefox%2090.0.1.dmg",
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
        sha256 = "76c1b9c42b52c7e5be4c112a98b7d3762a18841367f778a179679ac0de751f05",
        # Firefox v90.0.1
        urls = [
            "https://ftp.mozilla.org/pub/firefox/releases/90.0.1/mac/en-US/Firefox%2090.0.1.dmg",
            "https://storage.googleapis.com/dev-infra-mirror/mozilla/firefox/Firefox%2090.0.1.dmg",
        ],
        named_files = {
            "FIREFOX": "Firefox.app/Contents/MacOS/firefox",
        },
    )

def org_mozilla_geckodriver():
    platform_archive(
        name = "org_mozilla_geckodriver_linux_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "ec164910a3de7eec71e596bd2a1814ae27ba4c9d112b611680a6470dbe2ce27b",
        # Geckodriver v0.29.1
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.29.1/geckodriver-v0.29.1-linux64.tar.gz",
            "https://storage.googleapis.com/dev-infra-mirror/mozilla/geckodriver/0.29.1/geckodriver-v0.29.1-linux64.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )

    platform_archive(
        name = "org_mozilla_geckodriver_macos_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "9929c804ad0157ca13fdafca808866c88815b658e7059280a9f08f7e70364963",
        # Geckodriver v0.29.1
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.29.1/geckodriver-v0.29.1-macos.tar.gz",
            "https://storage.googleapis.com/dev-infra-mirror/mozilla/geckodriver/0.29.1/geckodriver-v0.29.1-macos.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )

    platform_archive(
        name = "org_mozilla_geckodriver_macos_arm64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "a1ec058b930fbfb684e30071ea47eec61bc18acb489914a9e0d095ede6088eea",
        # Geckodriver v0.29.1
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.29.1/geckodriver-v0.29.1-macos-aarch64.tar.gz",
            "https://storage.googleapis.com/dev-infra-mirror/mozilla/geckodriver/0.29.1/geckodriver-v0.29.1-macos-aarch64.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )
