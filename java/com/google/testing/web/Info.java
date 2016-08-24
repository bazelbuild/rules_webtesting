package com.google.testing.web;

/** Represents basic information about the browser defined by the web test environment. */
public class Info {

  private final String browserName;
  private final String formFactor;
  private final String targetBrowserName;
  private final String environment;

  Info(Metadata metadata) {
    this(
        metadata.getBrowserName(),
        metadata.getFormFactor(),
        metadata.getBrowserLabel().split(":")[1],
        metadata.getEnvironment());
  }

  private Info(
      String browserName, String formFactor, String targetBrowserName, String environment) {
    this.browserName = browserName;
    this.formFactor = formFactor;
    this.targetBrowserName = targetBrowserName;
    this.environment = environment;
  }

  /** Returns the name of the browser. */
  public String getBrowserName() {
    return this.browserName;
  }

  /**
   * Returns the browser environment that Web Test Launcher is using for the specified
   * configuration.
   *
   * <p>Note that for LTEP, this returns the backing technology, which is equivalent.
   */
  public String getEnvironment() {
    return this.environment;
  }

  /**
   * Returns a {@link String} representing the form factor of the device on which the browser is
   * running.
   *
   * <p>Examples:
   *
   * <ul>
   * <li>"DESKTOP"
   * <li>"PHONE"
   * <li>"TABLET"
   * </ul>
   */
  public String getFormFactor() {
    return this.formFactor;
  }

  /**
   * Returns the browser name component of the generated test target name; for example, the target
   * browser name of a target ":WebTestSuite_chrome-linux" would be "chrome-linux".
   */
  public String getTargetBrowserName() {
    return this.targetBrowserName;
  }
}
