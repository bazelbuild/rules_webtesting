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
#
################################################################################
#
package(default_testonly = 1)

load("//web:web.bzl", "browser", "web_test_archive", "web_test_named_executable")

config_setting(
    name = "mac",
    values = {"cpu": "darwin"},
)

config_setting(
    name = "linux",
    values = {"cpu": "k8"},
)

web_test_archive(
    name = "chromedriver",
    archive = select({
        ":linux": "@org_chromium_chromedriver_linux//file",
        ":mac": "@org_chromium_chromedriver_mac//file",
    }),
    named_files = {
        "CHROMEDRIVER": "chromedriver",
    },
)

web_test_archive(
    name = "chrome",
    archive = select({
        ":linux": "@com_google_chrome_linux//file",
        ":mac": "@com_google_chrome_mac//file",
    }),
    named_files = select({
        ":linux": {
            "CHROME": "chrome-precise64/chrome",
        },
        ":mac": {
            "CHROME": "chrome-mac/Google Chrome.app/Contents/MacOS/Google Chrome",
        },
    }),
)

web_test_archive(
    name = "firefox",
    archive = select({
        ":linux": "@org_mozilla_firefox_linux//file",
        ":mac": "@org_mozilla_firefox_mac//file",
    }),
    named_files = select({
        ":linux": {
            "FIREFOX": "firefox-sdk/bin/firefox",
        },
        ":mac": {
            "FIREFOX": "firefox-sdk/bin/Firefox.app/Contents/MacOS/firefox",
        },
    }),
)

web_test_archive(
    name = "geckodriver",
    archive = select({
        ":linux": "@org_mozilla_geckodriver_linux//file",
        ":mac": "@org_mozilla_geckodriver_mac//file",
    }),
    named_files = {
        "GECKODRIVER": "geckodriver",
    },
)

web_test_archive(
    name = "phantomjs_bin",
    archive = select({
        ":linux": "@org_phantomjs_linux//file",
        ":mac": "@org_phantomjs_mac//file",
    }),
    named_files = select({
        ":linux": {
            "PHANTOMJS": "phantomjs-2.1.1-linux-x86_64/bin/phantomjs",
        },
        ":mac": {
            "PHANTOMJS": "phantomjs-2.1.1-macosx/bin/phantomjs",
        },
    }),
)

browser(
    name = "firefox-external",
    metadata = "firefox-external.json",
    required_tags = ["requires-network"],
    visibility = ["//visibility:public"],
)

browser(
    name = "chrome-external",
    metadata = "chrome-external.json",
    required_tags = ["requires-network"],
    visibility = ["//visibility:public"],
)

browser(
    name = "firefox-native",
    data = [
        ":firefox",
        ":geckodriver",
        "//java:SELENIUM_SERVER",
    ],
    environment = select({
        ":linux": None,
        ":mac": {"DISABLE_X_DISPLAY": "1"},
    }),
    metadata = "firefox-native.json",
    # if not tagged local, then bazel does not exit cleanly
    required_tags = ["local"],
    visibility = ["//visibility:public"],
)

browser(
    name = "chrome-native",
    data = [
        ":chrome",
        ":chromedriver",
        "//java:SELENIUM_SERVER",
    ],
    environment = select({
        ":linux": None,
        ":mac": {"DISABLE_X_DISPLAY": "1"},
    }),
    metadata = "chrome-native.json",
    # if not tagged local, then bazel does not exit cleanly
    required_tags = ["local"],
    visibility = ["//visibility:public"],
)

browser(
    name = "phantomjs",
    data = [
        ":phantomjs_bin",
        "//java:SELENIUM_SERVER",
    ],
    metadata = "phantomjs.json",
    visibility = ["//visibility:public"],
)
