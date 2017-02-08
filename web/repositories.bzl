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
"""Defines external repositories needed by rules_webtesting."""

load("//web/internal:java_import_external.bzl", "java_import_external")
load("//web/internal:platform_http_file.bzl", "platform_http_file")
load("//web/internal:go_repository.bzl", "go_repository")

def web_test_repositories(
    omit_cglib_nodep=False,
    omit_com_github_gorilla_mux=False,
    omit_com_github_tebeka_selenium=False,
    omit_com_google_code_findbugs_jsr305=False,
    omit_com_google_code_gson=False,
    omit_com_google_errorprone_error_prone_annotations=False,
    omit_com_google_guava=False,
    omit_commons_codec=False,
    omit_commons_logging=False,
    omit_junit=False,
    omit_net_java_dev_jna=False,
    omit_net_java_dev_jna_platform=False,
    omit_org_apache_commons_exec=False,
    omit_org_apache_httpcomponents_httpclient=False,
    omit_org_apache_httpcomponents_httpcore=False,
    omit_org_apache_httpcomponents_httpmime=False,
    omit_org_eclipse_jetty_io=False,
    omit_org_eclipse_jetty_util=False,
    omit_org_eclipse_jetty_websocket_api=False,
    omit_org_eclipse_jetty_websocket_client=False,
    omit_org_eclipse_jetty_websocket_common=False,
    omit_org_hamcrest_core=False,
    omit_org_seleniumhq_py=False,
    omit_org_seleniumhq_selenium_api=False,
    omit_org_seleniumhq_selenium_remote_driver=False,
    **kwargs):
  """Defines external repositories required by Webtesting Rules.

  This function exists for other Bazel projects to call from their WORKSPACE
  file when depending on rules_webtesting using http_archive. This function
  makes it easy to import these transitive dependencies into the parent
  workspace using a blacklist model. Individual dependencies may be excluded
  with the omit parameters. This is useful for users who want to be rigorous
  about declaring their own direct dependencies, or when another Bazel project
  is depended upon (e.g. rules_closure) that defines the same dependencies as
  this one (e.g. com_google_guava.) Alternatively, a whitelist model may be
  used by calling the individual functions this method references.

  Please note that while these dependencies are defined, they are not actually
  downloaded, unless a target is built that depends on them.
  """
  _check_bazel_version("Web Testing Rules", "0.4.2")
  if kwargs.keys():
    print("The following parameters are deprecated: " + str(kwargs.keys()))
  if not omit_cglib_nodep:
    cglib_nodep()
  if not omit_com_github_gorilla_mux:
    com_github_gorilla_mux()
  if not omit_com_github_tebeka_selenium:
    com_github_tebeka_selenium()
  if not omit_com_google_code_findbugs_jsr305:
    com_google_code_findbugs_jsr305()
  if not omit_com_google_code_gson:
    com_google_code_gson()
  if not omit_com_google_errorprone_error_prone_annotations:
    com_google_errorprone_error_prone_annotations()
  if not omit_com_google_guava:
    com_google_guava()
  if not omit_commons_codec:
    commons_codec()
  if not omit_commons_logging:
    commons_logging()
  if not omit_junit:
    junit()
  if not omit_net_java_dev_jna:
    net_java_dev_jna()
  if not omit_net_java_dev_jna_platform:
    net_java_dev_jna_platform()
  if not omit_org_apache_commons_exec:
    org_apache_commons_exec()
  if not omit_org_apache_httpcomponents_httpclient:
    org_apache_httpcomponents_httpclient()
  if not omit_org_apache_httpcomponents_httpcore:
    org_apache_httpcomponents_httpcore()
  if not omit_org_apache_httpcomponents_httpmime:
    org_apache_httpcomponents_httpmime()
  if not omit_org_hamcrest_core:
    org_hamcrest_core()
  if not omit_org_seleniumhq_py:
    org_seleniumhq_py()
  if not omit_org_seleniumhq_selenium_api:
    org_seleniumhq_selenium_api()
  if not omit_org_seleniumhq_selenium_remote_driver:
    org_seleniumhq_selenium_remote_driver()


def browser_repositories(firefox=False, chromium=False):
  """Sets up repositories for browsers defined in //browsers/....

  This should only be used on an experimental basis; projects should define
  their own browsers.

  Args:
    firefox: Configure repositories for //browsers:firefox-native.
    chromium: Configure repositories for //browsers:chromium-native.
  """
  _check_bazel_version("Web Testing Rules", "0.4.2")
  if chromium:
    org_chromium_chromedriver()
    org_chromium_chromium()
  if firefox:
    org_mozilla_firefox()
    org_mozilla_geckodriver()


def cglib_nodep():
  java_import_external(
      name="cglib_nodep",
      jar_sha256="b40d7ac4400ea21dcf818f436a346ddd66b67c550811ccac5bbcbec095ab1287",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/cglib/cglib-nodep/3.2.4/cglib-nodep-3.2.4.jar",
          "http://maven.ibiblio.org/maven2/cglib/cglib-nodep/3.2.4/cglib-nodep-3.2.4.jar",
          "http://repo1.maven.org/maven2/cglib/cglib-nodep/3.2.4/cglib-nodep-3.2.4.jar",
      ],
      licenses=["notice"],  # ASF 2.0
  )


def com_github_gorilla_mux():
  go_repository(
      name="com_github_gorilla_mux",
      import_name="github.com/gorilla/mux",
      sha256="a32c13a36c58cb321136231ae8b67b0c6ad3c5f462e65eb6771f59c44b44ccba",
      strip_prefix="mux-757bef944d0f21880861c2dd9c871ca543023cba",
      excluded_srcs = ["context_gorilla.go"],
      license = "licenses([\"notice\"])",
      urls=[
          "http://bazel-mirror.storage.googleapis.com/github.com/gorilla/mux/archive/757bef944d0f21880861c2dd9c871ca543023cba.tar.gz",
          "https://github.com/gorilla/mux/archive/757bef944d0f21880861c2dd9c871ca543023cba.tar.gz",
      ])


def com_github_tebeka_selenium():
  go_repository(
      name="com_github_tebeka_selenium",
      import_name="github.com/tebeka/selenium",
      sha256="c33decb47a9b81d5221cda29c8f040ca5cf874956bbb002ef82b06e07ed78c3d",
      strip_prefix="selenium-f6f9a3638fa049f85b0aaf42e693e1c4ab257d4f",
      license="licenses([\"notice\"])  # MIT.",
      urls=[
          "http://bazel-mirror.storage.googleapis.com/github.com/tebeka/selenium/archive/f6f9a3638fa049f85b0aaf42e693e1c4ab257d4f.tar.gz",
          "https://github.com/tebeka/selenium/archive/f6f9a3638fa049f85b0aaf42e693e1c4ab257d4f.tar.gz",
      ])


def com_google_code_findbugs_jsr305():
  java_import_external(
      name="com_google_code_findbugs_jsr305",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/com/google/code/findbugs/jsr305/1.3.9/jsr305-1.3.9.jar",
          "http://repo1.maven.org/maven2/com/google/code/findbugs/jsr305/1.3.9/jsr305-1.3.9.jar",
          "http://maven.ibiblio.org/maven2/com/google/code/findbugs/jsr305/1.3.9/jsr305-1.3.9.jar",
      ],
      jar_sha256="905721a0eea90a81534abb7ee6ef4ea2e5e645fa1def0a5cd88402df1b46c9ed",
      licenses=["notice"],  # BSD 3-clause
  )


def com_google_code_gson():
  java_import_external(
      name="com_google_code_gson",
      jar_sha256="13f44a2f6ead058da80a91ee650c073871942468e684a9bf6a0d0319138924ce",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/com/google/code/gson/gson/2.3.1/gson-2.3.1.jar",
          "http://repo1.maven.org/maven2/com/google/code/gson/gson/2.3.1/gson-2.3.1.jar",
          "http://maven.ibiblio.org/maven2/com/google/code/gson/gson/2.3.1/gson-2.3.1.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      deps=["@com_google_code_findbugs_jsr305"])


def com_google_errorprone_error_prone_annotations():
  java_import_external(
      name="com_google_errorprone_error_prone_annotations",
      jar_sha256="e7749ffdf03fb8ebe08a727ea205acb301c8791da837fee211b99b04f9d79c46",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.0.15/error_prone_annotations-2.0.15.jar",
          "http://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.0.15/error_prone_annotations-2.0.15.jar",
      ],
      licenses=["notice"],  # Apache 2.0
  )


def com_google_guava():
  java_import_external(
      name="com_google_guava",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/com/google/guava/guava/20.0/guava-20.0.jar",
          "http://repo1.maven.org/maven2/com/google/guava/guava/20.0/guava-20.0.jar",
          "http://maven.ibiblio.org/maven2/com/google/guava/guava/20.0/guava-20.0.jar",
      ],
      jar_sha256="36a666e3b71ae7f0f0dca23654b67e086e6c93d192f60ba5dfd5519db6c288c8",
      licenses=["notice"],  # Apache 2.0
      deps=[
          "@com_google_code_findbugs_jsr305",
          "@com_google_errorprone_error_prone_annotations",
      ])


def commons_codec():
  java_import_external(
      name="commons_codec",
      jar_sha256="4241dfa94e711d435f29a4604a3e2de5c4aa3c165e23bd066be6fc1fc4309569",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/commons-codec/commons-codec/1.10/commons-codec-1.10.jar",
          "http://repo1.maven.org/maven2/commons-codec/commons-codec/1.10/commons-codec-1.10.jar",
          "http://maven.ibiblio.org/maven2/commons-codec/commons-codec/1.10/commons-codec-1.10.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
  )


def commons_logging():
  java_import_external(
      name="commons_logging",
      jar_sha256="daddea1ea0be0f56978ab3006b8ac92834afeefbd9b7e4e6316fca57df0fa636",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
          "http://maven.ibiblio.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
          "http://repo1.maven.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
  )


def junit():
  java_import_external(
      name="junit",
      jar_sha256="59721f0805e223d84b90677887d9ff567dc534d7c502ca903c0c2b17f05c116a",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
          "http://repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
          "http://maven.ibiblio.org/maven2/junit/junit/4.12/junit-4.12.jar",
      ],
      licenses=["reciprocal"],  # Eclipse Public License 1.0
      testonly_=1,
      deps=["@org_hamcrest_core"])


def net_java_dev_jna():
  java_import_external(
      name="net_java_dev_jna",
      jar_sha256="1aa37e9ea6baa0ee152d89509f758f0847eac66ec179b955cafe0919e540a92e",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/net/java/dev/jna/jna/4.1.0/jna-4.1.0.jar",
          "http://maven.ibiblio.org/maven2/net/java/dev/jna/jna/4.1.0/jna-4.1.0.jar",
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna/4.1.0/jna-4.1.0.jar",
      ],
      # LGPL, version 2.1
      # http://www.gnu.org/licenses/licenses.html
      # ASL, version 2
      # http://www.apache.org/licenses/
      licenses=["restricted"])


def net_java_dev_jna_platform():
  java_import_external(
      name="net_java_dev_jna_platform",
      jar_sha256="f91ba7c0f26c34f04bf57d2ae30d4b19f906e7bb1de90eb3e1f4fdbf45d0c541",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/net/java/dev/jna/jna-platform/4.1.0/jna-platform-4.1.0.jar",
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna-platform/4.1.0/jna-platform-4.1.0.jar",
          "http://maven.ibiblio.org/maven2/net/java/dev/jna/jna-platform/4.1.0/jna-platform-4.1.0.jar",
      ],
      # LGPL, version 2.1
      # http://www.gnu.org/licenses/licenses.html
      # ASL, version 2
      # http://www.apache.org/licenses/
      licenses=["restricted"],
      deps=["@net_java_dev_jna"])


def org_apache_commons_exec():
  java_import_external(
      name="org_apache_commons_exec",
      jar_sha256="cb49812dc1bfb0ea4f20f398bcae1a88c6406e213e67f7524fb10d4f8ad9347b",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
          "http://repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
          "http://maven.ibiblio.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
  )


def org_apache_httpcomponents_httpclient():
  java_import_external(
      name="org_apache_httpcomponents_httpclient",
      jar_sha256="0dffc621400d6c632f55787d996b8aeca36b30746a716e079a985f24d8074057",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.2/httpclient-4.5.2.jar",
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.2/httpclient-4.5.2.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpclient/4.5.2/httpclient-4.5.2.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
      deps=[
          "@org_apache_httpcomponents_httpcore",
          "@commons_logging",
          "@commons_codec",
      ])


def org_apache_httpcomponents_httpcore():
  java_import_external(
      name="org_apache_httpcomponents_httpcore",
      jar_sha256="f7bc09dc8a7003822d109634ffd3845d579d12e725ae54673e323a7ce7f5e325",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.4/httpcore-4.4.4.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpcore/4.4.4/httpcore-4.4.4.jar",
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.4/httpcore-4.4.4.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
  )


def org_apache_httpcomponents_httpmime():
  java_import_external(
      name="org_apache_httpcomponents_httpmime",
      jar_sha256="231a3f7e4962053db2be8461d5422e68fc458a3a7dd7d8ada803a348e21f8f07",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/org/apache/httpcomponents/httpmime/4.5.2/httpmime-4.5.2.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpmime/4.5.2/httpmime-4.5.2.jar",
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpmime/4.5.2/httpmime-4.5.2.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
      deps=["@org_apache_httpcomponents_httpclient"])


def org_chromium_chromedriver():
  platform_http_file(
      name="org_chromium_chromedriver",
      amd64_sha256="59e6b1b1656a20334d5731b3c5a7400f92a9c6f5043bb4ab67f1ccf1979ee486",
      amd64_urls=[
          "http://chromedriver.storage.googleapis.com/2.26/chromedriver_linux64.zip"
      ],
      macos_sha256="70aae3812941ed94ad8065bb4a9432861d7d4ebacdd93ee47bb2c7c57c7e841e",
      macos_urls=[
          "http://chromedriver.storage.googleapis.com/2.26/chromedriver_mac64.zip"
      ])


def org_chromium_chromium():
  # Roughly corresponds to Chrome 55
  platform_http_file(
      name="org_chromium_chromium",
      amd64_sha256="e3c99954d6acce013174053534b72f47f67f18a0d75f79c794daaa8dd2ae8aaf",
      amd64_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/423768/chrome-linux.zip"
      ],
      macos_sha256="62aeb7a5c6b8a1b7b31400105bf01295bbd45b0627920b8f99f0cc4ca76927ca",
      macos_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/423758/chrome-mac.zip"
      ])


def org_hamcrest_core():
  java_import_external(
      name="org_hamcrest_core",
      jar_sha256="66fdef91e9739348df7a096aa384a5685f4e875584cce89386a7a47251c4d8e9",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
          "http://repo1.maven.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
          "http://maven.ibiblio.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
      ],
      licenses=["notice"],  # New BSD License
      testonly_=1)


def org_mozilla_firefox():
  platform_http_file(
      name="org_mozilla_firefox",
      amd64_sha256="10533f3db9c819a56f6cd72f9340e05c7e3b116454eb81b0d39ed161955bb48f",
      amd64_urls=[
          "http://ftp.mozilla.org/pub/firefox/releases/50.1.0/firefox-50.1.0.linux-x86_64.sdk.tar.bz2",
          "http://bazel-mirror.storage.googleapis.com/ftp.mozilla.org/pub/firefox/releases/50.1.0/firefox-50.1.0.linux-x86_64.sdk.tar.bz2",
      ],
      macos_sha256="5cd449ebedb44b2f882b37e6e5cee1a814bc5ff3c3f86d1a1019b937aa287441",
      macos_urls=[
          "http://ftp.mozilla.org/pub/firefox/releases/50.1.0/firefox-50.1.0.mac-x86_64.sdk.tar.bz2",
          "http://bazel-mirror.storage.googleapis.com/ftp.mozilla.org/pub/firefox/releases/50.1.0/firefox-50.1.0.mac-x86_64.sdk.tar.bz2",
      ])


def org_mozilla_geckodriver():
  platform_http_file(
      name="org_mozilla_geckodriver",
      amd64_sha256="ce4aa8b5cf918a6607b50e73996fb909db42fd803855f0ecc9d7183999c3bedc",
      amd64_urls=[
          "http://bazel-mirror.storage.googleapis.com/github.com/mozilla/geckodriver/releases/download/v0.11.1/geckodriver-v0.11.1-linux64.tar.gz",
          "https://github.com/mozilla/geckodriver/releases/download/v0.11.1/geckodriver-v0.11.1-linux64.tar.gz",
      ],
      macos_sha256="802cc1a33b8ce6f7c3aeb5116730cb6efc20414959d6f750e74437869d37a150",
      macos_urls=[
          "http://bazel-mirror.storage.googleapis.com/github.com/mozilla/geckodriver/releases/download/v0.11.1/geckodriver-v0.11.1-macos.tar.gz",
          "https://github.com/mozilla/geckodriver/releases/download/v0.11.1/geckodriver-v0.11.1-macos.tar.gz",
      ])


def org_seleniumhq_py():
  native.new_http_archive(
      name="org_seleniumhq_py",
      build_file=str(Label("//build_files:org_seleniumhq_py.BUILD")),
      sha256="85daad4d09be86bddd4f45579986ac316c1909c3b4653ed471ea4519eb413c8f",
      strip_prefix="selenium-3.0.2/py",
      urls=[
          "http://bazel-mirror.storage.googleapis.com/pypi.python.org/packages/0c/42/20c235e604bf736bc970c1275a78c4ea28c6453a0934002f95df9c49dad0/selenium-3.0.2.tar.gz",
          "https://pypi.python.org/packages/0c/42/20c235e604bf736bc970c1275a78c4ea28c6453a0934002f95df9c49dad0/selenium-3.0.2.tar.gz",
      ])


def org_seleniumhq_selenium_api():
  java_import_external(
      name="org_seleniumhq_selenium_api",
      jar_sha256="0226cc02880aff06f7fd85e77314182087a524e21ceda02f8197317bbb0390b8",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.0.1/selenium-api-3.0.1.jar",
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.0.1/selenium-api-3.0.1.jar",
          "http://maven.ibiblio.org/maven2/org/seleniumhq/selenium/selenium-api/3.0.1/selenium-api-3.0.1.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      testonly_=1)


def org_seleniumhq_selenium_remote_driver():
  java_import_external(
      name="org_seleniumhq_selenium_remote_driver",
      jar_sha256="97eed1fe99c4b5ced127336270fe56fa53754627f24536bc07141c6451270275",
      jar_urls=[
          "http://bazel-mirror.storage.googleapis.com/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.0.1/selenium-remote-driver-3.0.1.jar",
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.0.1/selenium-remote-driver-3.0.1.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      testonly_=1,
      deps=[
          "@cglib_nodep",
          "@com_google_code_gson",
          "@com_google_guava",
          "@net_java_dev_jna_platform",
          "@org_apache_commons_exec",
          "@org_apache_httpcomponents_httpmime",
          "@org_seleniumhq_selenium_api",
      ])


def _check_bazel_version(project, bazel_version):
  if "bazel_version" not in dir(native):
    fail("%s requires Bazel >=%s but was <0.2.1" % (project, bazel_version))
  elif not native.bazel_version:
    pass  # user probably compiled Bazel from scratch
  else:
    current_bazel_version = _parse_bazel_version(native.bazel_version)
    minimum_bazel_version = _parse_bazel_version(bazel_version)
    if minimum_bazel_version > current_bazel_version:
      fail("%s requires Bazel >=%s but was %s" % (project, bazel_version,
                                                  native.bazel_version))


def _parse_bazel_version(bazel_version):
  # Remove commit from version.
  version = bazel_version.split(" ", 1)[0]
  # Split into (release, date) parts and only return the release
  # as a tuple of integers.
  parts = version.split("-", 1)
  # Turn "release" into a tuple of strings
  version_tuple = ()
  for number in parts[0].split("."):
    version_tuple += (str(number),)
  return version_tuple
