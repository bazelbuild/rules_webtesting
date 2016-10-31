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
"""Browser provisioning and information API.

Provision a browser:
  driver = webtest.new_webdriver_session()

Provision a browser with capabilities:
  capabilities = {"webdriver.logging.profiler.enabled": true}
  driver = browser.Browser().new_webdriver_session(capabilities)
"""
import os

from selenium.webdriver.remote import remote_connection
from selenium.webdriver.remote import webdriver


def new_webdriver_session(capabilities={}):
  """Provisions a new WebDriver session.

  Args:
    capabilities: a dict with the json capabilities desired for this browser
      session.

  Returns:
    A new WebDriver connected to a browser defined by the web test
    environment.
  """
  address = os.environ['REMOTE_WEBDRIVER_SERVER'].rstrip('/')

  # Set the timeout for WebDriver http requests so that the socket default
  # timeout is not used.
  remote_connection.RemoteConnection.set_timeout(450)

  return webdriver.WebDriver(address, desired_capabilities=capabilities)
