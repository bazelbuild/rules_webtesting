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
        sha256 = "8b9823d07706db02d0f83189c7d658fff14796ccb07f3eac3b327f3f0230f6c7",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Linux_x64/1121551/chromedriver_linux64.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/linux_x64/driver-bin.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_linux64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_macos_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "1c66bd01e53ee406f9f30d5f9ccbf4ea0f9c0f1b959c6ace9758cf0d35a6e4b3",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac/1121551/chromedriver_mac64.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/mac_x64/driver-bin.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_mac64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_macos_arm64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "7eba8dd97537ca787628ae11346e5c897473c0c0871df0fc4a313bd4a48a83dc",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/1121551/chromedriver_mac64.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/mac_arm64/driver-bin.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_mac64/chromedriver",
        },
    )

    platform_archive(
        name = "org_chromium_chromedriver_windows_x64",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        sha256 = "db94b7f35041e3a76fa9a50808f196e61c27f43762df99526c1876244a196526",
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Win/1121551/chromedriver_win32.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/windows_x64/driver-bin.zip",
        ],
        named_files = {
            "CHROMEDRIVER": "chromedriver_win32/chromedriver.exe",
        },
    )

def org_chromium_chromium():
    platform_archive(
        name = "org_chromium_chromium_linux_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "4e89a56b61db2fe494d4072d551b24e81833608318c5ba347b5d16a19687674e",
        # 114.0.5673.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Linux_x64/f/chrome-linux.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/linux_x64/browser-bin.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-linux/chrome",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_macos_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "c03e32f338dffee3404881b4950563d26812d0246c1372ad2f4800547382bb91",
        # 114.0.5673.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac/1121551/chrome-mac.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/mac_x64/browser-bin.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_macos_arm64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "4eb94b113fc995d20fafeca366b4b0cddf172ac1b2cdedc053464b764b74d1c0",
        # 114.0.5673.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/1121551/chrome-mac.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/mac_arm64/browser-bin.zip",
        ],
        named_files = {
            "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium",
        },
    )

    platform_archive(
        name = "org_chromium_chromium_windows_x64",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        sha256 = "fdc221bb1e898ab851c4a5bc50ca1f88a5b388acb5510df4c4606c87d8be0230",
        # 114.0.5673.0
        urls = [
            "https://storage.googleapis.com/chromium-browser-snapshots/Win/1121551/chrome-win.zip",
            "https://storage.googleapis.com/dev-infra-mirror/chromium/1121551/windows_x64/browser-bin.zip",
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
        name = "org_mozilla_firefox_macos",
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
        sha256 = "12c37f41d11ed982b7be43d02411ff2c75fb7a484e46966d000b47d1665baa88",
        # Geckodriver v0.30.0
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.30.0/geckodriver-v0.30.0-linux64.tar.gz",
            "https://storage.googleapis.com/dev-infra-mirror/firefox/97.0/linux_x64/driver-bin.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )

    platform_archive(
        name = "org_mozilla_geckodriver_macos_x64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "560ba192666c1fe8796404153cfdf2d12551515601c4b3937aabcba6ee300f8c",
        # Geckodriver v0.30.0
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.30.0/geckodriver-v0.30.0-macos.tar.gz",
            "https://storage.googleapis.com/dev-infra-mirror/firefox/97.0/mac_x64/driver-bin.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )

    platform_archive(
        name = "org_mozilla_geckodriver_macos_arm64",
        licenses = ["reciprocal"],  # MPL 2.0
        sha256 = "895bc2146edaea434d57a3b5d9a141be5cb3c5f8e8804916bd4869978ddfd4db",
        # Geckodriver v0.30.0
        urls = [
            "https://github.com/mozilla/geckodriver/releases/download/v0.30.0/geckodriver-v0.30.0-macos-aarch64.tar.gz",
            "https://storage.googleapis.com/dev-infra-mirror/firefox/97.0/mac_arm64/driver-bin.tar.gz",
        ],
        named_files = {
            "GECKODRIVER": "geckodriver",
        },
    )
