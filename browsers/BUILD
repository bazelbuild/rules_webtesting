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

web_test_archive(
    name = "chromedriver",
    archive = "@org_chromium_chromedriver//file",
    named_files = {
        "CHROMEDRIVER": "chromedriver",
    },
)

web_test_archive(
    name = "chrome",
    archive = "@com_google_chrome//file",
    named_files = {
        "CHROME": "chrome-precise64/chrome",
    },
)

web_test_archive(
    name = "firefox",
    archive = "@org_mozilla_firefox//file",
    named_files = {
        "FIREFOX": "firefox-sdk/bin/firefox",
    },
)

web_test_archive(
    name = "geckodriver",
    archive = "@org_mozilla_geckodriver//file",
    named_files = {
        "GECKODRIVER": "geckodriver",
    },
)

web_test_archive(
    name = "phantomjs_bin",
    archive = "@org_phantomjs//file",
    named_files = {
        "PHANTOMJS": "phantomjs-2.1.1-linux-x86_64/bin/phantomjs",
    },
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
