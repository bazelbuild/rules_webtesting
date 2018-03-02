// Copyright 2018 Google Inc.
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

import com.google.testing.bazel.Bazel
import com.google.testing.web.WebTest
import java.awt.Image
import java.nio.file.Files
import java.nio.file.Path
import org.openqa.selenium.By
import org.openqa.selenium.WebDriver
import org.scalatest._

class ScreenshotterTest
  extends WordSpec
  with BeforeAndAfter
  with BeforeAndAfterAll
  with Matchers {

  var tmpDir: Path = _

  private var driver: WebDriver = _
  private var screenshotter: Screenshotter = _

  override def beforeAll() = {
    tmpDir = Bazel.getInstance.newTmpDir(this.getClass.getSimpleName)
  }

  before {
    driver = new WebTest().newWebDriverSession()
    screenshotter = new Screenshotter(driver)
    driver.get(Bazel.getInstance.runfile("testdata/testpage.html").toUri.toString)
  }


  after {
    try {
      driver.quit()
    } finally {
      driver = null
      screenshotter = null
    }
  }

  "A WebDriver" should {
    "take a screenshot" in {
      val ss = screenshotter.takeScreenshot()
      val out = tmpDir.resolve("basicScreenshot.png")
      Files.write(out, ss.asBytes())
      System.out.println("image written to: " + out)
      // No assertions. Check output manually.
    }

    "take a screenshot of an element" in {
      val ss = screenshotter.of(driver.findElement(By.tagName("input"))).takeScreenshot()
      val out = tmpDir.resolve("ofInputElement.png")
      Files.write(out, ss.asBytes())
      System.out.println("image written to: " + out)
      val img = ss.asImage()
      200 shouldEqual img.getHeight(null)
      200 shouldEqual img.getWidth(null)
    }

    "exclude elements from a screenshot" in {
      val ss = screenshotter
        .of(driver.findElement(By.id("outer-div")))
        .excluding(driver.findElement(By.id("inner-div1")), driver.findElement(By.id("b")))
        .takeScreenshot()
      val out = tmpDir.resolve("excludingElements.png")
      Files.write(out, ss.asBytes())
      System.out.println("image written to: " + out)
      // No assertions. Check output manually.
    }
  }
}
