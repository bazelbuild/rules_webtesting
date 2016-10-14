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
"""Tests for google3.testing.web.py.browser.Browser."""

import unittest
from testing.web import browser


class BrowserTest(unittest.TestCase):

  def testBrowserProvisioningNoCaps(self):
    driver = browser.Browser().new_session()

    try:
      driver.get("about:")
      self.assertTrue(driver.current_url)
    finally:
      driver.quit()

  def testBrowserProvisioningWithCaps(self):
    capabilities = {
        "unexpectedAlertBehaviour": "dismiss",
        "elementScrollBehavior": 1,
    }
    driver = browser.Browser().new_session(capabilities)

    try:
      driver.get("about:")
      self.assertTrue(driver.current_url)
    finally:
      driver.quit()


if __name__ == "__main__":
  unittest.main()
