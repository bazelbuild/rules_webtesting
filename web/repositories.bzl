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
load("@bazel_skylib//:lib.bzl", "versions")


def web_test_repositories(**kwargs):
  """Defines external repositories required by Webtesting Rules.

  This function exists for other Bazel projects to call from their WORKSPACE
  file when depending on rules_webtesting using http_archive. This function
  makes it easy to import these transitive dependencies into the parent
  workspace. This will check to see if a repository has been previously defined
  before defining a new repository.

  Alternatively, individual dependencies may be excluded with an
  "omit_" + name parameter. This is useful for users who want to be rigorous
  about declaring their own direct dependencies, or when another Bazel project
  is depended upon (e.g. rules_closure) that defines the same dependencies as
  this one (e.g. com_google_guava.) Alternatively, a whitelist model may be
  used by calling the individual functions this method references.

  Please note that while these dependencies are defined, they are not actually
  downloaded, unless a target is built that depends on them.
  """
  versions.check("0.9.0")
  if should_create_repository("cglib_nodep", kwargs):
    cglib_nodep()
  if should_create_repository("com_github_blang_semver", kwargs):
    com_github_blang_semver()
  if should_create_repository("com_github_gorilla_context", kwargs):
    com_github_gorilla_context()
  if should_create_repository("com_github_gorilla_mux", kwargs):
    com_github_gorilla_mux()
  if should_create_repository("com_github_tebeka_selenium", kwargs):
    com_github_tebeka_selenium()
  if should_create_repository("com_google_code_findbugs_jsr305", kwargs):
    com_google_code_findbugs_jsr305()
  if should_create_repository("com_google_code_gson", kwargs):
    com_google_code_gson()
  if should_create_repository("com_google_errorprone_error_prone_annotations",
                              kwargs):
    com_google_errorprone_error_prone_annotations()
  if should_create_repository("com_google_guava", kwargs):
    com_google_guava()
  if should_create_repository("com_squareup_okhttp3_okhttp", kwargs):
    com_squareup_okhttp3_okhttp()
  if should_create_repository("com_squareup_okio", kwargs):
    com_squareup_okio()
  if should_create_repository("commons_codec", kwargs):
    commons_codec()
  if should_create_repository("commons_logging", kwargs):
    commons_logging()
  if should_create_repository("junit", kwargs):
    junit()
  if should_create_repository("net_bytebuddy", kwargs):
    net_bytebuddy()
  if should_create_repository("net_java_dev_jna", kwargs):
    net_java_dev_jna()
  if should_create_repository("net_java_dev_jna_platform", kwargs):
    net_java_dev_jna_platform()
  if should_create_repository("org_apache_commons_exec", kwargs):
    org_apache_commons_exec()
  if should_create_repository("org_apache_httpcomponents_httpclient", kwargs):
    org_apache_httpcomponents_httpclient()
  if should_create_repository("org_apache_httpcomponents_httpcore", kwargs):
    org_apache_httpcomponents_httpcore()
  if should_create_repository("org_apache_httpcomponents_httpmime", kwargs):
    org_apache_httpcomponents_httpmime()
  if should_create_repository("org_hamcrest_core", kwargs):
    org_hamcrest_core()
  if should_create_repository("org_json", kwargs):
    org_json()
  if should_create_repository("org_seleniumhq_py", kwargs):
    org_seleniumhq_py()
  if should_create_repository("org_seleniumhq_selenium_api", kwargs):
    org_seleniumhq_selenium_api()
  if should_create_repository("org_seleniumhq_selenium_remote_driver", kwargs):
    org_seleniumhq_selenium_remote_driver()
  if kwargs.keys():
    print("The following parameters are unknown: " + str(kwargs.keys()))


def should_create_repository(name, args):
  """Returns whether the name repository should be created.

  This allows creation of a repository to be disabled by either an
  "omit_" _+ name parameter or by previously defining a rule for the repository.

  The args dict will be mutated to remove "omit_" + name.
  """
  key = "omit_" + name
  if key in args:
    val = args.pop(key)
    if val:
      return False
  if native.existing_rule(name):
    return False
  return True


def browser_repositories(firefox=False, chromium=False):
  """Sets up repositories for browsers defined in //browsers/....

  This should only be used on an experimental basis; projects should define
  their own browsers.

  Args:
    firefox: Configure repositories for //browsers:firefox-native.
    chromium: Configure repositories for //browsers:chromium-native.
  """
  versions.check("0.9.0")
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
      "80a0cc14c7c495682b43d4b082ad80a5848ada19fc3700f72d8ec042923633d3",
      jar_urls=[
          "http://repo1.maven.org/maven2/cglib/cglib-nodep/3.2.6/cglib-nodep-3.2.6.jar",
      ],
      licenses=["notice"]  # ASF 2.0
  )


def com_github_blang_semver():
  go_repository(
      name="com_github_blang_semver",
      importpath="github.com/blang/semver",
      sha256="3d9da53f4c2d3169bfa9b25f2f36f301a37556a47259c870881524c643c69c57",
      strip_prefix="semver-3.5.1",
      urls=[
          "https://github.com/blang/semver/archive/v3.5.1.tar.gz",
      ])


def com_github_gorilla_context():
  go_repository(
      name="com_github_gorilla_context",
      importpath="github.com/gorilla/context",
      sha256="12a849b4e9a08619233d4490a281aa2d34a69f9eaf85c2295f5357927e4d1763",
      strip_prefix="context-1.1",
      urls=[
          "https://github.com/gorilla/context/archive/v1.1.tar.gz",
      ])


def com_github_gorilla_mux():
  go_repository(
      name="com_github_gorilla_mux",
      importpath="github.com/gorilla/mux",
      sha256="fe4d6909570b53121eb0d5e6f933ef7c49d5de094705af6ba07fab9c299df0f9",
      strip_prefix="mux-1.6.1",
      urls=[
          "https://github.com/gorilla/mux/archive/v1.6.1.tar.gz",
      ])


def com_github_tebeka_selenium():
  go_repository(
      name="com_github_tebeka_selenium",
      importpath="github.com/tebeka/selenium",
      sha256="c731518d8f4724b1c581bdca90215978bf4d2658a5dc49b720d7c61294888396",
      strip_prefix="selenium-a789e65b0e7f126888873e84f528c1c8537dff3e",
      urls=[
          "https://github.com/tebeka/selenium/archive/a789e65b0e7f126888873e84f528c1c8537dff3e.tar.gz",
      ])


def com_google_code_findbugs_jsr305():
  java_import_external(
      name="com_google_code_findbugs_jsr305",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar",
      ],
      jar_sha256=
      "766ad2a0783f2687962c8ad74ceecc38a28b9f72a2d085ee438b7813e928d0c7",
      licenses=["notice"]  # BSD 3-clause
  )


def com_google_code_gson():
  java_import_external(
      name="com_google_code_gson",
      jar_sha256=
      "b7134929f7cc7c04021ec1cc27ef63ab907e410cf0588e397b8851181eb91092",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/google/code/gson/gson/2.8.2/gson-2.8.2.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      deps=["@com_google_code_findbugs_jsr305"])


def com_google_errorprone_error_prone_annotations():
  java_import_external(
      name="com_google_errorprone_error_prone_annotations",
      jar_sha256=
      "6ebd22ca1b9d8ec06d41de8d64e0596981d9607b42035f9ed374f9de271a481a",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.2.0/error_prone_annotations-2.2.0.jar",
      ],
      licenses=["notice"]  # Apache 2.0
  )


def com_google_guava():
  java_import_external(
      name="com_google_guava",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/google/guava/guava/24.0-jre/guava-24.0-jre.jar",
      ],
      jar_sha256=
      "e0274470b16ba1154e926b5f54ef8ae159197fbc356406bda9b261ba67e3e599",
      licenses=["notice"],  # Apache 2.0
      exports=[
          "@com_google_code_findbugs_jsr305",
          "@com_google_errorprone_error_prone_annotations",
      ])


def com_squareup_okhttp3_okhttp():
  java_import_external(
      name="com_squareup_okhttp3_okhttp",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/3.9.1/okhttp-3.9.1.jar",
      ],
      jar_sha256=
      "a0d01017a42bba26e507fc6d448bb36e536f4b6e612f7c42de30bbdac2b7785e",
      licenses=["notice"],  # Apache 2.0
      deps=[
          "@com_squareup_okio",
          "@com_google_code_findbugs_jsr305",
      ])


def com_squareup_okio():
  java_import_external(
      name="com_squareup_okio",
      jar_urls=[
          "http://repo1.maven.org/maven2/com/squareup/okio/okio/1.14.0/okio-1.14.0.jar",
      ],
      jar_sha256=
      "4633c331f50642ebe795dc089d6a5928aff43071c9d17e7840a009eea2fe95a3",
      licenses=["notice"],  # Apache 2.0
      deps=["@com_google_code_findbugs_jsr305"])


def commons_codec():
  java_import_external(
      name="commons_codec",
      jar_sha256=
      "e599d5318e97aa48f42136a2927e6dfa4e8881dff0e6c8e3109ddbbff51d7b7d",
      jar_urls=[
          "http://repo1.maven.org/maven2/commons-codec/commons-codec/1.11/commons-codec-1.11.jar",
      ],
      licenses=["notice"]  # Apache License, Version 2.0
  )


def commons_logging():
  java_import_external(
      name="commons_logging",
      jar_sha256=
      "daddea1ea0be0f56978ab3006b8ac92834afeefbd9b7e4e6316fca57df0fa636",
      jar_urls=[
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
          "http://repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
      ],
      licenses=["reciprocal"],  # Eclipse Public License 1.0
      testonly_=1,
      deps=["@org_hamcrest_core"])


def net_bytebuddy():
  java_import_external(
      name="net_bytebuddy",
      jar_sha256=
      "2ea2ada12b790d16ac7f6e6c065cb55cbcdb6ba519355f5958851159cad3b16a",
      jar_urls=[
          "http://repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.7.9/byte-buddy-1.7.9.jar",
      ],
      # LGPL, version 2.1
      # http://www.gnu.org/licenses/licenses.html
      # ASL, version 2
      # http://www.apache.org/licenses/
      licenses=["restricted"])


def net_java_dev_jna():
  java_import_external(
      name="net_java_dev_jna",
      jar_sha256=
      "fbc9de96a0cc193a125b4008dbc348e9ed54e5e13fc67b8ed40e645d303cc51b",
      jar_urls=[
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna/4.5.1/jna-4.5.1.jar",
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
      "84c8667555ee8dd91fef44b451419f6f16f71f727d5fc475a10c2663eba83abb",
      jar_urls=[
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna-platform/4.5.1/jna-platform-4.5.1.jar",
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
          "http://repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
      ],
      licenses=["notice"]  # Apache License, Version 2.0
  )


def org_apache_httpcomponents_httpclient():
  java_import_external(
      name="org_apache_httpcomponents_httpclient",
      jar_sha256=
      "7e97724443ad2a25ad8c73183431d47cc7946271bcbbdfa91a8a17522a566573",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.5/httpclient-4.5.5.jar",
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
      "1b4a1c0b9b4222eda70108d3c6e2befd4a6be3d9f78ff53dd7a94966fdf51fc5",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.9/httpcore-4.4.9.jar",
      ],
      licenses=["notice"]  # Apache License, Version 2.0
  )


def org_apache_httpcomponents_httpmime():
  java_import_external(
      name="org_apache_httpcomponents_httpmime",
      jar_sha256=
      "e46206931b7426102e658f086f74ee582761264a8f9977fba02c1e200c51a9c5",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpmime/4.5.5/httpmime-4.5.5.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
      deps=["@org_apache_httpcomponents_httpclient"])


def org_chromium_chromedriver():
  platform_http_file(
      name="org_chromium_chromedriver",
      amd64_sha256=
      "67fad24c4a85e3f33f51c97924a98b619722db15ce92dcd27484fb748af93e8e",
      amd64_urls=[
          "http://chromedriver.storage.googleapis.com/2.35/chromedriver_linux64.zip"
      ],
      macos_sha256=
      "c11521bdc991874a0a29cf36beea6a4b8d73616ce6c8d1e6b90067d85718aa87",
      macos_urls=[
          "http://chromedriver.storage.googleapis.com/2.35/chromedriver_mac64.zip"
      ],
      windows_sha256=
      "b32fdcc1c19bb829032f0447a8aac4f5436565ec1d0f105b63c4451ba4e6ae8a",
      windows_urls=[
          "http://chromedriver.storage.googleapis.com/2.35/chromedriver_win32.zip"
      ])


def org_chromium_chromium():
  platform_http_file(
      name="org_chromium_chromium",
      amd64_sha256=
      "51a189382cb5272d240a729da0ae77d0211c1bbc0d10b701a2723b5b068c1e3a",
      amd64_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/539259/chrome-linux.zip"
      ],
      macos_sha256=
      "866ec9aa4e07cc86ae1d5aeb6e9bdafb5f94989c7c0be661302930ad667f41f3",
      macos_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/539251/chrome-mac.zip"
      ],
      windows_sha256=
      "be4fcc7257d85c12ae2de10aef0150ddbb7b9ecbd5ada6a898d247cf867a058a",
      windows_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Win_x64/539249/chrome-win32.zip"
      ])


def org_hamcrest_core():
  java_import_external(
      name="org_hamcrest_core",
      jar_sha256=
      "66fdef91e9739348df7a096aa384a5685f4e875584cce89386a7a47251c4d8e9",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
      ],
      licenses=["notice"],  # New BSD License
      testonly_=1)


def org_json():
  java_import_external(
      name="org_json",
      jar_sha256=
      "3eddf6d9d50e770650e62abe62885f4393aa911430ecde73ebafb1ffd2cfad16",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/json/json/20180130/json-20180130.jar",
      ],
      licenses=["notice"]  # MIT-style license
  )


def org_mozilla_firefox():
  platform_http_file(
      name="org_mozilla_firefox",
      amd64_sha256=
      "134fec04819eb56fa7b644cdd6d89623b21f4020bbedc3bd122db2a2caa4e434",
      amd64_urls=[
          "https://ftp.mozilla.org/pub/firefox/releases/58.0/linux-x86_64/en-US/firefox-58.0.tar.bz2",
      ],
      macos_sha256=
      "b9e9f383fd12d6deb6b36b3c2844647ce065142e1eb934499559a8a9842d01ad",
      macos_urls=[
          "https://ftp.mozilla.org/pub/firefox/releases/52.1.2esr/firefox-52.1.2esr.mac-x86_64.sdk.tar.bz2",
      ])


def org_mozilla_geckodriver():
  platform_http_file(
      name="org_mozilla_geckodriver",
      amd64_sha256=
      "7f55c4c89695fd1e6f8fc7372345acc1e2dbaa4a8003cee4bd282eed88145937",
      amd64_urls=[
          "https://github.com/mozilla/geckodriver/releases/download/v0.19.1/geckodriver-v0.19.1-linux64.tar.gz",
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
      sha256="a34a833d89bcfb463bfba5e5515a9276bb94221787b409f0ad28d2f91903e31d",
      strip_prefix="selenium-3.9.0",
      urls=[
          "https://pypi.python.org/packages/b4/54/ba7059b254a72fc30f1d8b838eb951003ee6e5ba716bb9b0ce0e4c58e308/selenium-3.9.0.tar.gz"
      ])


def org_seleniumhq_selenium_api():
  java_import_external(
      name="org_seleniumhq_selenium_api",
      jar_sha256=
      "040871bcfeb0ac522b2c2a1507ab0046c10fead3c22468fef78d2a815b55ad00",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.9.1/selenium-api-3.9.1.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      testonly_=1)


def org_seleniumhq_selenium_remote_driver():
  java_import_external(
      name="org_seleniumhq_selenium_remote_driver",
      jar_sha256=
      "1bf029a5c0f034072f11655662710ed72ffa577166daa94de82c1c6073515b11",
      jar_urls=[
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.8.1/selenium-remote-driver-3.8.1.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      testonly_=1,
      deps=[
          "@cglib_nodep",
          "@com_google_code_gson",
          "@com_google_guava",
          "@net_bytebuddy",
          "@net_java_dev_jna_platform",
          "@org_apache_commons_exec",
          "@org_apache_httpcomponents_httpmime",
          "@org_seleniumhq_selenium_api",
      ])
