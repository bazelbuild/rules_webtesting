library webtest;

/// Browser provisioning and information API.
///
/// Provision a browser:
///     import "package:webtest/webtest.dart" as webtest;
///
///     WebDriver driver = webtest.newWebDriverSession();
///
/// Provision a browser with capabilities (as an example, profiling):
///     var capabilities = <String, dynamic>{
///       Capabilities.enableProfiling: true,
///     };
///
///     WebDriver driver =  webtest.newWebDriverSession(capabilities);

import "dart:async" show Future;
import "dart:io" show Platform;

import "package:webdriver/core.dart" show WebDriver;
import "package:webdriver/io.dart" show createDriver;

/// Provisions and returns a new [WebDriver] session.
Future<WebDriver> newWebDriverSession(
    [Map<String, dynamic> capabilities = const <String, dynamic>{}]) async {
  var address = Platform.environment["WEB_TEST_WEBDRIVER_SERVER"];
  if (!address.endsWith("/")) {
    address += "/";
  }

  return createDriver(uri: Uri.parse(address), desired: capabilities);
}
