package com.google.testing.web;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;
import org.openqa.selenium.WebDriver;

@RunWith(JUnit4.class)
public class BrowserTest {

  @Test
  public void newSession() {
    WebDriver driver = new Browser().newSession();
    driver.quit();
  }
}
