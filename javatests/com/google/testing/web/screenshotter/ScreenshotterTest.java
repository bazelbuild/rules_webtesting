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
package com.google.testing.web.screenshotter;

import static org.junit.Assert.assertEquals;

import com.google.testing.util.Bazel;
import com.google.testing.web.WebTest;
import java.awt.Image;
import java.nio.file.Files;
import java.nio.file.Path;
import org.junit.After;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;
import org.openqa.selenium.By;
import org.openqa.selenium.WebDriver;

@RunWith(JUnit4.class)
public class ScreenshotterTest {

  private static Path tmpDir;

  private WebDriver driver;
  private Screenshotter screenshotter;

  @BeforeClass
  public static void createTmpDir() throws Exception {
    tmpDir = Bazel.getInstance().newTmpDir(ScreenshotterTest.class.getSimpleName());
  }

  @Before
  public void createDriver() throws Exception {
    driver = new WebTest().newWebDriverSession();
    screenshotter = new Screenshotter(driver);
    driver.get(
        Bazel.getInstance().runfile("go/launcher/proxy/testdata/testpage.html").toUri().toString());
  }

  @After
  public void quitDriver() throws Exception {
    try {
      driver.quit();
    } finally {
      driver = null;
      screenshotter = null;
    }
  }

  @Test
  public void basicScreenshot() throws Exception {
    Screenshot ss = screenshotter.takeScreenshot();
    Path out = tmpDir.resolve("basicScreenshot.png");
    Files.write(out, ss.asBytes());
    System.out.println("image written to: " + out);
    // No assertions. Check output manually.
  }

  @Test
  public void ofInputElement() throws Exception {
    Screenshot ss = screenshotter.of(driver.findElement(By.tagName("input"))).takeScreenshot();
    Path out = tmpDir.resolve("ofInputElement.png");
    Files.write(out, ss.asBytes());
    System.out.println("image written to: " + out);
    Image img = ss.asImage();
    assertEquals(200, img.getHeight(null));
    assertEquals(200, img.getWidth(null));
  }

  @Test
  public void excludingElements() throws Exception {
    Screenshot ss =
        screenshotter
            .of(driver.findElement(By.id("outer-div")))
            .excluding(driver.findElement(By.id("inner-div1")), driver.findElement(By.id("b")))
            .takeScreenshot();
    Path out = tmpDir.resolve("excludingElements.png");
    Files.write(out, ss.asBytes());
    System.out.println("image written to: " + out);
    // No assertions. Check output manually.
  }
}
