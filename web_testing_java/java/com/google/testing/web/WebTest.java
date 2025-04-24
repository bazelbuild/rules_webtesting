// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// //////////////////////////////////////////////////////////////////////////////
//
package com.google.testing.web;

import java.net.MalformedURLException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URL;
import java.util.Optional;
import org.openqa.selenium.Capabilities;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.remote.Augmenter;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.openqa.selenium.remote.RemoteWebDriver;

/**
 * Browser provisioning and information API.
 *
 * <p>Provision a browser:
 *
 * <pre class="code">
 * import com.google.testing.web.WebTest;
 *
 * WebDriver driver = new WebTest().newWebDriverSession();
 * </pre>
 *
 * <p>Provision a browser with capabilities (as an example, profiling):
 *
 * <pre class="code">
 * DesiredCapabilities capabilities = new DesiredCapabilities();
 * capabilities.setCapability(CapabilityType.ENABLE_PROFILING_CAPABILITY, true);
 *
 * WebDriver driver = new WebTest().newWebDriverSession(capabilities);
 * </pre>
 */
public class WebTest {

  private final URL wd;
  private final URI http;
  private final Optional<URI> https;

  public WebTest() {
    this(
        System.getenv("WEB_TEST_WEBDRIVER_SERVER"),
        System.getenv("WEB_TEST_HTTP_SERVER"),
        System.getenv("WEB_TEST_HTTPS_SERVER"));
  }

  private WebTest(String wd, String http, String https) {
    try {
      this.wd = new URL(wd);
      this.http = new URI(http);
      if (https != null && !https.isEmpty()) {
        this.https = Optional.of(new URI(https));
      } else {
        this.https = Optional.empty();
      }
    } catch (MalformedURLException | URISyntaxException e) {
      throw new RuntimeException(e);
    }
  }

  /** Provisions and returns a new default {@link WebDriver} session. */
  public WebDriver newWebDriverSession() {
    return newWebDriverSession(new DesiredCapabilities());
  }

  /**
   * Provisions and returns a new {@link WebDriver} session.
   *
   * @param capabilities Configuration of the browser.
   */
  public WebDriver newWebDriverSession(Capabilities capabilities) {
    WebDriver driver = new Augmenter().augment(new RemoteWebDriver(wd, capabilities));

    return driver;
  }

  /** Returns the HTTP address of WTL. */
  public URI HTTPAddress() {
    return http;
  }

  /** Returns the HTTPS address of WTL. */
  public Optional<URI> HTTPSAddress() {
    return https;
  }
}
