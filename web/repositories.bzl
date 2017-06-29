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
load("@io_bazel_rules_go//go:def.bzl", "go_repository")


def web_test_repositories(
    omit_cglib_nodep=False,
    omit_com_github_gorilla_mux=False,
    omit_com_github_tebeka_selenium=False,
    omit_com_google_code_findbugs_jsr305=False,
    omit_com_google_code_gson=False,
    omit_com_google_errorprone_error_prone_annotations=False,
    omit_com_google_guava=False,
    omit_com_squareup_okhttp3_okhttp=False,
    omit_com_squareup_okio=False,
    omit_commons_codec=False,
    omit_commons_logging=False,
    omit_junit=False,
    omit_net_java_dev_jna=False,
    omit_net_java_dev_jna_platform=False,
    omit_org_apache_commons_exec=False,
    omit_org_apache_httpcomponents_httpclient=False,
    omit_org_apache_httpcomponents_httpcore=False,
    omit_org_apache_httpcomponents_httpmime=False,
    omit_org_hamcrest_core=False,
    omit_org_json=False,
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
  if not omit_com_squareup_okhttp3_okhttp:
    com_squareup_okhttp3_okhttp()
  if not omit_com_squareup_okio:
    com_squareup_okio()
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
  if not omit_org_json:
    org_json()
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
      jar_sha256=
      "9ab68b361ffdce66747f63bbe0676391c0a96d6de3e743ac03d7e998ff0a1064",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/cglib/cglib-nodep/3.2.5/cglib-nodep-3.2.5.jar",
          "http://maven.ibiblio.org/maven2/cglib/cglib-nodep/3.2.5/cglib-nodep-3.2.5.jar",
          "http://repo1.maven.org/maven2/cglib/cglib-nodep/3.2.5/cglib-nodep-3.2.5.jar",
      ],
      licenses=["notice"]  # ASF 2.0
  )


def com_github_gorilla_mux():
  go_repository(
      name="com_github_gorilla_mux",
      importpath="github.com/gorilla/mux",
      sha256="1a1b35782b0e38534b81d90bb86993fe830a7c1c3974a562554399f850ccdfcd",
      strip_prefix="mux-1.4.0",
      build_tags=["go1.9"],
      urls=["https://github.com/gorilla/mux/archive/v1.4.0.tar.gz"])


def com_github_tebeka_selenium():
  go_repository(
      name="com_github_tebeka_selenium",
      importpath="github.com/tebeka/selenium",
      sha256="345f204a3ece2469dcc82b59860dec095006e179b1c2ba3e9433b14c90dae167",
      strip_prefix="selenium-8f4861d1f09c100da29ceec85424c3c96df15170",
      urls=
      ["https://github.com/tebeka/selenium/archive/8f4861d1f09c100da29ceec85424c3c96df15170.tar.gz"]
  )


def com_google_code_findbugs_jsr305():
  java_import_external(
      name="com_google_code_findbugs_jsr305",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar",
          "http://repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar",
          "http://maven.ibiblio.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar",
      ],
      jar_sha256=
      "766ad2a0783f2687962c8ad74ceecc38a28b9f72a2d085ee438b7813e928d0c7",
      licenses=["notice"]  # BSD 3-clause
  )


def com_google_code_gson():
  java_import_external(
      name="com_google_code_gson",
      jar_sha256=
      "c6221763bd79c4f1c3dc7f750b5f29a0bb38b367b81314c4f71896e340c40825",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/google/code/gson/gson/2.8.0/gson-2.8.0.jar",
          "http://repo1.maven.org/maven2/com/google/code/gson/gson/2.8.0/gson-2.8.0.jar",
          "http://maven.ibiblio.org/maven2/com/google/code/gson/gson/2.8.0/gson-2.8.0.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      deps=["@com_google_code_findbugs_jsr305"])


def com_google_errorprone_error_prone_annotations():
  java_import_external(
      name="com_google_errorprone_error_prone_annotations",
      jar_sha256=
      "cde78ace21e46398299d0d9c6be9f47b7f971c7f045d40c78f95be9a638cbf7e",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.0.19/error_prone_annotations-2.0.19.jar",
          "http://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.0.19/error_prone_annotations-2.0.19.jar",
      ],
      licenses=["notice"]  # Apache 2.0
  )


def com_google_guava():
  java_import_external(
      name="com_google_guava",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/google/guava/guava/22.0/guava-22.0.jar",
          "http://repo1.maven.org/maven2/com/google/guava/guava/22.0/guava-22.0.jar",
          "http://maven.ibiblio.org/maven2/com/google/guava/guava/22.0/guava-22.0.jar",
      ],
      jar_sha256=
      "1158e94c7de4da480873f0b4ab4a1da14c0d23d4b1902cc94a58a6f0f9ab579e",
      licenses=["notice"],  # Apache 2.0
      exports=[
          "@com_google_code_findbugs_jsr305",
          "@com_google_errorprone_error_prone_annotations",
      ])


def com_squareup_okhttp3_okhttp():
  java_import_external(
      name="com_squareup_okhttp3_okhttp",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/3.8.0/okhttp-3.8.0.jar",
      ],
      jar_sha256=
      "19e1db51787716ff0046fa19e408fb34ed32a6274baa0c07475bf724b4eb6800",
      licenses=["notice"],  # Apache 2.0
      deps=[
          "@com_squareup_okio",
          "@com_google_code_findbugs_jsr305",
      ])


def com_squareup_okio():
  java_import_external(
      name="com_squareup_okio",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/squareup/okio/okio/1.13.0/okio-1.13.0.jar",
      ],
      jar_sha256=
      "734269c3ebc5090e3b23566db558f421f0b4027277c79ad5d176b8ec168bb850",
      licenses=["notice"],  # Apache 2.0
      deps=["@com_google_code_findbugs_jsr305"])


def commons_codec():
  java_import_external(
      name="commons_codec",
      jar_sha256=
      "4241dfa94e711d435f29a4604a3e2de5c4aa3c165e23bd066be6fc1fc4309569",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/commons-codec/commons-codec/1.10/commons-codec-1.10.jar",
          "http://repo1.maven.org/maven2/commons-codec/commons-codec/1.10/commons-codec-1.10.jar",
          "http://maven.ibiblio.org/maven2/commons-codec/commons-codec/1.10/commons-codec-1.10.jar",
      ],
      licenses=["notice"]  # Apache License, Version 2.0
  )


def commons_logging():
  java_import_external(
      name="commons_logging",
      jar_sha256=
      "daddea1ea0be0f56978ab3006b8ac92834afeefbd9b7e4e6316fca57df0fa636",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
          "http://maven.ibiblio.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
          "http://repo1.maven.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
      ],
      licenses=["notice"]  # The Apache Software License, Version 2.0
  )


def junit():
  java_import_external(
      name="junit",
      jar_sha256=
      "59721f0805e223d84b90677887d9ff567dc534d7c502ca903c0c2b17f05c116a",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
          "http://repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
          "http://maven.ibiblio.org/maven2/junit/junit/4.12/junit-4.12.jar",
      ],
      licenses=["reciprocal"],  # Eclipse Public License 1.0
      testonly_=1,
      deps=["@org_hamcrest_core"])


def net_java_dev_jna():
  java_import_external(
      name="net_java_dev_jna",
      jar_sha256=
      "c4dadeeecaa90c8847902082aee5eb107fcf59c5d0e63a17fcaf273c0e2d2bd1",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/net/java/dev/jna/jna/4.4.0/jna-4.4.0.jar",
          "http://maven.ibiblio.org/maven2/net/java/dev/jna/jna/4.4.0/jna-4.4.0.jar",
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna/4.4.0/jna-4.4.0.jar",
      ],
      # LGPL, version 2.1
      # http://www.gnu.org/licenses/licenses.html
      # ASL, version 2
      # http://www.apache.org/licenses/
      licenses=["restricted"])


def net_java_dev_jna_platform():
  java_import_external(
      name="net_java_dev_jna_platform",
      jar_sha256=
      "e9dda9e884fc107eb6367710540789a12dfa8ad28be9326b22ca6e352e325499",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/net/java/dev/jna/jna-platform/4.4.0/jna-platform-4.4.0.jar",
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna-platform/4.4.0/jna-platform-4.4.0.jar",
          "http://maven.ibiblio.org/maven2/net/java/dev/jna/jna-platform/4.4.0/jna-platform-4.4.0.jar",
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
      jar_sha256=
      "cb49812dc1bfb0ea4f20f398bcae1a88c6406e213e67f7524fb10d4f8ad9347b",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
          "http://repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
          "http://maven.ibiblio.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
      ],
      licenses=["notice"]  # Apache License, Version 2.0
  )


def org_apache_httpcomponents_httpclient():
  java_import_external(
      name="org_apache_httpcomponents_httpclient",
      jar_sha256=
      "db3d1b6c2d6a5e5ad47577ad61854e2f0e0936199b8e05eb541ed52349263135",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.3/httpclient-4.5.3.jar",
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.3/httpclient-4.5.3.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpclient/4.5.3/httpclient-4.5.3.jar",
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
      jar_sha256=
      "d7f853dee87680b07293d30855b39b9eb56c1297bd16ff1cd6f19ddb8fa745fb",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.6/httpcore-4.4.6.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpcore/4.4.6/httpcore-4.4.6.jar",
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.6/httpcore-4.4.6.jar",
      ],
      licenses=["notice"]  # Apache License, Version 2.0
  )


def org_apache_httpcomponents_httpmime():
  java_import_external(
      name="org_apache_httpcomponents_httpmime",
      jar_sha256=
      "b4865b79a3aaeef794220b532bc7b07f793fa4aad90c29e83cab2b835cd8ee06",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/apache/httpcomponents/httpmime/4.5.3/httpmime-4.5.3.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpmime/4.5.3/httpmime-4.5.3.jar",
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpmime/4.5.3/httpmime-4.5.3.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
      deps=["@org_apache_httpcomponents_httpclient"])


def org_chromium_chromedriver():
  platform_http_file(
      name="org_chromium_chromedriver",
      amd64_sha256=
      "bb2cf08f2c213f061d6fbca9658fc44a367c1ba7e40b3ee1e3ae437be0f901c2",
      amd64_urls=[
          "http://chromedriver.storage.googleapis.com/2.29/chromedriver_linux64.zip"
      ],
      macos_sha256=
      "6c30bba7693ec2d9af7cd9a54729e10aeae85c0953c816d9c4a40a1a72fd8be0",
      macos_urls=[
          "http://chromedriver.storage.googleapis.com/2.29/chromedriver_mac64.zip"
      ])


def org_chromium_chromium():
  # Roughly corresponds to Chrome 58
  platform_http_file(
      name="org_chromium_chromium",
      amd64_sha256=
      "c356dbaee39b6a070388e87566c10254def32641df50b21c28aff1b8b11aeb5f",
      amd64_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/454471/chrome-linux.zip"
      ],
      macos_sha256=
      "740d691b07855e2aace1e524fd67b8732458e52cc8fca0b4c1bddbbb3aa9ee11",
      macos_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/454475/chrome-mac.zip"
      ])


def org_hamcrest_core():
  java_import_external(
      name="org_hamcrest_core",
      jar_sha256=
      "66fdef91e9739348df7a096aa384a5685f4e875584cce89386a7a47251c4d8e9",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
          "http://repo1.maven.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
          "http://maven.ibiblio.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
      ],
      licenses=["notice"],  # New BSD License
      testonly_=1)


def org_json():
  java_import_external(
      name="org_json",
      jar_sha256=
      "813f37e4820f1854e8a4eb4f80df94bf1b1f2ec6c3b72692f23ab9a556256af6",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/json/json/20170516/json-20170516.jar",
      ],
      licenses=["notice"]  # MIT-style license
  )


def org_mozilla_firefox():
  platform_http_file(
      name="org_mozilla_firefox",
      amd64_sha256=
      "6d338e98b1ac1078738dcdab11d6a0c855e5087fa0d7d99dbffa521d16a45327",
      amd64_urls=[
          "https://ftp.mozilla.org/pub/firefox/releases/53.0.3/linux-x86_64/en-US/firefox-53.0.3.tar.bz2",
          "http://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/53.0.3/linux-x86_64/en-US/firefox-53.0.3.tar.bz2",
      ],
      macos_sha256=
      "b9e9f383fd12d6deb6b36b3c2844647ce065142e1eb934499559a8a9842d01ad",
      macos_urls=[
          "https://ftp.mozilla.org/pub/firefox/releases/52.1.2esr/firefox-52.1.2esr.mac-x86_64.sdk.tar.bz2",
          "http://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/52.1.2esr/firefox-52.1.2esr.mac-x86_64.sdk.tar.bz2",
      ])


def org_mozilla_geckodriver():
  platform_http_file(
      name="org_mozilla_geckodriver",
      amd64_sha256=
      "dcadab8586264cf33aae1fff0897520d46e39dad4580c6cae712452fdc59e529",
      amd64_urls=[
          "http://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.16.1/geckodriver-v0.16.1-linux64.tar.gz",
          "https://github.com/mozilla/geckodriver/releases/download/v0.16.1/geckodriver-v0.16.1-linux64.tar.gz",
      ],
      macos_sha256=
      "eb5a2971e5eb4a2fe74a3b8089f0f2cc96eed548c28526b8351f0f459c080836",
      macos_urls=[
          "http://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.16.1/geckodriver-v0.16.1-macos.tar.gz",
          "https://github.com/mozilla/geckodriver/releases/download/v0.16.1/geckodriver-v0.16.1-macos.tar.gz",
      ])


def org_seleniumhq_py():
  native.new_http_archive(
      name="org_seleniumhq_py",
      build_file=str(Label("//build_files:org_seleniumhq_py.BUILD")),
      sha256="2befcd8a18c4a88fe7af1f848277d83aa89682516fb2ac95942cb13ea6eaad02",
      strip_prefix="selenium-3.4.2",
      urls=[
          "http://mirror.bazel.build/pypi.python.org/packages/47/63/611503113c00cb3ba0ab310567b4f7d7c2a3178e7ebc07f66ae4b997d56a/selenium-3.4.2.tar.gz",
          "https://pypi.python.org/packages/47/63/611503113c00cb3ba0ab310567b4f7d7c2a3178e7ebc07f66ae4b997d56a/selenium-3.4.2.tar.gz"
      ])


def org_seleniumhq_selenium_api():
  java_import_external(
      name="org_seleniumhq_selenium_api",
      jar_sha256=
      "a14e8fb856b1840a837dbca8d2dd322cd4a7d8650924aa4dfa19eca06f47401e",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.4.0/selenium-api-3.4.0.jar",
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.4.0/selenium-api-3.4.0.jar",
          "http://maven.ibiblio.org/maven2/org/seleniumhq/selenium/selenium-api/3.4.0/selenium-api-3.4.0.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      testonly_=1)


def org_seleniumhq_selenium_remote_driver():
  java_import_external(
      name="org_seleniumhq_selenium_remote_driver",
      jar_sha256=
      "47b88da5cb9c92f832af51db4fdf6b0a6aa70e7a76ed641137c344a8fad5cc03",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.4.0/selenium-remote-driver-3.4.0.jar",
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.4.0/selenium-remote-driver-3.4.0.jar",
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
