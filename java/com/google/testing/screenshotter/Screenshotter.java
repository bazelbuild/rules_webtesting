// Copyright 2017 Google Inc.
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
package com.google.testing.screenshotter;

import com.google.common.collect.ImmutableSet;
import java.io.IOException;
import java.util.Collection;
import java.util.Optional;
import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.remote.Dialect;
import org.openqa.selenium.remote.RemoteWebDriver;
import org.openqa.selenium.remote.RemoteWebElement;

/**
 * API for interacting with advanced screenshot handler installed into Web Test Launcher.
 *
 * <p>Screenshotter objects are immutable; methods that return Screenshotter return an new
 * Screenshotter object that is a copy of the original with appropriate modifications.
 */
public final class Screenshotter {

  private static final MediaType JSON = MediaType.parse("application/json; charset=utf-8");

  private static final OkHttpClient client = new OkHttpClient();

  private final String url;
  private final Optional<WebElement> element;
  private final ImmutableSet<WebElement> excluding;

  /**
   * Create a new Screenshotter that interacts with the WebDriver session to which driver is
   * connected.
   */
  public Screenshotter(WebDriver driver) {
    this(
        String.format(
            "%s/session/%s/google/screenshot",
            System.getenv("WEB_TEST_WEBDRIVER_SERVER"),
            String.valueOf(((RemoteWebDriver) driver).getSessionId())),
        Optional.empty(),
        ImmutableSet.of());
  }

  private Screenshotter(
      String url, Optional<WebElement> element, Collection<WebElement> excluding) {
    this.url = url;
    this.element = element;
    this.excluding = ImmutableSet.copyOf(excluding);
  }

  /** Returns a copy of this configured to take a screenshot of element. */
  public Screenshotter of(WebElement element) {
    return new Screenshotter(this.url, Optional.of(element), this.excluding);
  }

  /** Returns a copy of this configured to blackout elements. */
  public Screenshotter excluding(WebElement... elements) {
    ImmutableSet<WebElement> ex =
        new ImmutableSet.Builder<WebElement>().addAll(this.excluding).add(elements).build();
    return new Screenshotter(this.url, this.element, ex);
  }

  public Screenshot takeScreenshot() throws IOException, JSONException {
    Response response =
        client
            .newCall(
                new Request.Builder()
                    .url(url)
                    .post(RequestBody.create(JSON, getRequestBodyJson().toString()))
                    .build())
            .execute();
    JSONObject responseBody = new JSONObject(response.body().string());

    if (!response.isSuccessful()) {
      throw new IOException(responseBody.getString("message"));
    }
    return new Screenshot(responseBody);
  }

  private JSONObject getRequestBodyJson() throws JSONException {
    JSONObject params = new JSONObject();

    if (element.isPresent()) {
      params.put("Element", webElementToJSON(element.get()));
    }

    if (!excluding.isEmpty()) {
      JSONArray ex = new JSONArray();
      for (WebElement el : excluding) {
        ex.put(webElementToJSON(el));
      }
      params.put("Exclude", ex);
    }

    return params;
  }

  private static JSONObject webElementToJSON(WebElement element) {
    String id = ((RemoteWebElement) element).getId();
    JSONObject object = new JSONObject();
    object.put(Dialect.OSS.getEncodedElementKey(), id);
    object.put(Dialect.W3C.getEncodedElementKey(), id);
    return object;
  }
}
