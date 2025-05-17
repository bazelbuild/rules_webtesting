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
"""Tests for testing.web.webtest."""

import unittest

from testing.web import webtest


class BrowserTest(unittest.TestCase):

  def testBrowserProvisioningNoCaps(self):
    driver = webtest.new_webdriver_session()

    try:
      driver.get(webtest.http_address() + "/healthz")
      self.assertTrue(driver.current_url)
    finally:
      driver.quit()

  def testBrowserProvisioningWithCaps(self):
    capabilities = {
        "acceptInsecureCerts": False,
        "pageLoadStrategy": "normal",
    }
    driver = webtest.new_webdriver_session(capabilities)

    try:
      driver.get(webtest.http_address() + "/healthz")
      self.assertTrue(driver.current_url)
    finally:
      driver.quit()


if __name__ == "__main__":
  unittest.main()
