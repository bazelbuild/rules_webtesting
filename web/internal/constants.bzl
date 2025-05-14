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
"""Constants used in other build rules."""

DEFAULT_WRAPPED_TEST_TAGS = ("manual", "noci")

DEFAULT_TEST_SUITE_TAGS = ("manual",)

DEFAULT_WEB_TEST_SUITE_TAGS = {
    "chrome-external": [
        "external",
    ],
    "chromium-local": [
        "native",
    ],
    "chromium-local-1024x768": [
        "native",
    ],
    "chrome-win10": [
        "exclusive",
        "sauce",
    ],
    "chrome-win10-connect": [
        "exclusive",
        "noci",
        "sauce",
    ],
    "firefox-external": [
        "external",
    ],
    "firefox-local": [
        "native",
    ],
}
