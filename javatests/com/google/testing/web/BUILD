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
package(default_testonly = True)

load("//web:web.bzl", "web_test", "web_test_suite")

licenses(["notice"])  # Apache 2.0

java_test(
    name = "WebTestTest_wrapped",
    srcs = ["WebTestTest.java"],
    tags = [
        "manual",
        "noci",
    ],
    test_class = "com.google.testing.web.WebTestTest",
    deps = [
        "//java/com/google/testing/web",
        "@junit_junit//jar",
        "@org_seleniumhq_java//:selenium",
    ],
)

web_test_suite(
    name = "WebTestTest",
    browser_overrides = {
        "//browsers:chrome-external": {
            "tags": [
                "manual",
                "noci",
                "requires-network",
            ],
        },
        "//browsers:firefox-external": {
            "tags": [
                "manual",
                "noci",
                "requires-network",
            ],
        },
        "//browsers:chromium-native": {
            "local": True,
        },
        "//browsers:firefox-native": {
            "local": True,
        },
    },
    browsers = [
        "//browsers:chrome-external",
        "//browsers:chromium-native",
        "//browsers:firefox-external",
        "//browsers:firefox-native",
        "//browsers:phantomjs",
    ],
    tags = ["noci"],
    test = ":WebTestTest_wrapped",
)
