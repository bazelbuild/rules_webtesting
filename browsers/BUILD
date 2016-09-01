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

load("//web:web.bzl", "browser", "web_test_named_executable")

web_test_named_executable(
    name = "CHROMEDRIVER",
    executable = "@org_chromium_chromedriver//:chromedriver",
)

web_test_named_executable(
    name = "CHROME",
    data = ["@com_google_chrome//:deps"],
    executable = "@com_google_chrome//:chrome-precise64/chrome",
)

web_test_named_executable(
    name = "FIREFOX",
    data = ["@org_mozilla_firefox//:deps"],
    executable = "@org_mozilla_firefox//:firefox-sdk/bin/firefox",
)

web_test_named_executable(
    name = "GECKODRIVER",
    executable = "@org_mozilla_geckodriver//:geckodriver",
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
        ":FIREFOX",
        ":GECKODRIVER",
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
        ":CHROME",
        ":CHROMEDRIVER",
        "//java:SELENIUM_SERVER",
    ],
    metadata = "chrome-native.json",
    # if not tagged local, then bazel does not exit cleanly
    required_tags = ["local"],
    visibility = ["//visibility:public"],
)
