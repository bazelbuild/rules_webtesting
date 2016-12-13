@TestOn("vm")
library webtest.webtest_test;

import "package:test/test.dart";

import "package:webtest/webtest.dart" as webtest;

void main() {
  test("browser provisioning, no capabilities", () async {
    final driver = await webtest.newWebDriverSession();

    try {
      await driver.get("about:");
      expect(await driver.currentUrl, isNotEmpty);
    } finally {
      await driver.quit();
    }
  });

  test("browser provisioning, with capabilities", () async {
    final driver = await webtest.newWebDriverSession(const <String, dynamic>{
      "unexpectedAlertBehaviour": "dismiss",
      "elementScrollBehavior": 1,
    });

    try {
      await driver.get("about:");
      expect(await driver.currentUrl, isNotEmpty);
    } finally {
      await driver.quit();
    }
  });
}
