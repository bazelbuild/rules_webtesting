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
  _check_bazel_version("Web Testing Rules", "0.4.2")
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
          "http://repo1.maven.org/maven2/cglib/cglib-nodep/3.2.5/cglib-nodep-3.2.5.jar",
          "http://maven.ibiblio.org/maven2/cglib/cglib-nodep/3.2.5/cglib-nodep-3.2.5.jar",
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
          "https://mirror.bazel.build/github.com/blang/semver/archive/v3.5.1.tar.gz",
          "https://github.com/blang/semver/archive/v3.5.1.tar.gz",
      ])


def com_github_gorilla_context():
  go_repository(
      name="com_github_gorilla_context",
      importpath="github.com/gorilla/context",
      sha256="12a849b4e9a08619233d4490a281aa2d34a69f9eaf85c2295f5357927e4d1763",
      strip_prefix="context-1.1",
      urls=[
          "https://mirror.bazel.build/github.com/gorilla/context/archive/v1.1.tar.gz",
          "https://github.com/gorilla/context/archive/v1.1.tar.gz",
      ])


def com_github_gorilla_mux():
  go_repository(
      name="com_github_gorilla_mux",
      importpath="github.com/gorilla/mux",
      sha256="e639e6935f3327bed70b583e3311eba262266e79f2c50d436bb99ca70700879f",
      strip_prefix="mux-1.5.0",
      urls=[
          "https://mirror.bazel.build/github.com/gorilla/mux/archive/v1.5.0.tar.gz",
          "https://github.com/gorilla/mux/archive/v1.5.0.tar.gz",
      ])


def com_github_tebeka_selenium():
  go_repository(
      name="com_github_tebeka_selenium",
      importpath="github.com/tebeka/selenium",
      sha256="22db05f2b3e7b1bfe0f84c45333be798627e1b091918a8c65e241b5b1e9df5cc",
      strip_prefix="selenium-a7e71178aa40d9b298d464b912c4e831b2be4455",
      urls=[
          "https://mirror.bazel.build/github.com/tebeka/selenium/archive/a7e71178aa40d9b298d464b912c4e831b2be4455.tar.gz",
          "https://github.com/tebeka/selenium/archive/a7e71178aa40d9b298d464b912c4e831b2be4455.tar.gz",
      ])


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
      "b7134929f7cc7c04021ec1cc27ef63ab907e410cf0588e397b8851181eb91092",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/google/code/gson/gson/2.8.2/gson-2.8.2.jar",
          "http://repo1.maven.org/maven2/com/google/code/gson/gson/2.8.2/gson-2.8.2.jar",
          "http://maven.ibiblio.org/maven2/com/google/code/gson/gson/2.8.2/gson-2.8.2.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      deps=["@com_google_code_findbugs_jsr305"])


def com_google_errorprone_error_prone_annotations():
  java_import_external(
      name="com_google_errorprone_error_prone_annotations",
      jar_sha256=
      "98590bc81d8065ec4ea8dc18ea8f482d996d050faa191d4595bafb324414166a",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.1.1/error_prone_annotations-2.1.1.jar",
          "http://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.1.1/error_prone_annotations-2.1.1.jar",
          "http://maven.ibiblio.org/maven2/com/google/errorprone/error_prone_annotations/2.1.1/error_prone_annotations-2.1.1.jar",
      ],
      licenses=["notice"]  # Apache 2.0
  )


def com_google_guava():
  java_import_external(
      name="com_google_guava",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/google/guava/guava/23.2-jre/guava-23.2-jre.jar",
          "http://repo1.maven.org/maven2/com/google/guava/guava/23.2-jre/guava-23.2-jre.jar",
          "http://maven.ibiblio.org/maven2/com/google/guava/guava/23.2-jre/guava-23.2-jre.jar",
      ],
      jar_sha256=
      "5be9a7d05ba0ccd74708bc8018ae412255f85843c0b92302e9b9befa6ed52564",
      licenses=["notice"],  # Apache 2.0
      exports=[
          "@com_google_code_findbugs_jsr305",
          "@com_google_errorprone_error_prone_annotations",
      ])


def com_squareup_okhttp3_okhttp():
  java_import_external(
      name="com_squareup_okhttp3_okhttp",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/3.9.0/okhttp-3.9.0.jar",
          "http://repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/3.9.0/okhttp-3.9.0.jar",
          "http://maven.ibiblio.org/maven2/com/squareup/okhttp3/okhttp/3.9.0/okhttp-3.9.0.jar",
      ],
      jar_sha256=
      "7265adbd6f028aade307f58569d814835cd02bc9beffb70c25f72c9de50d61c4",
      licenses=["notice"],  # Apache 2.0
      deps=[
          "@com_squareup_okio",
          "@com_google_code_findbugs_jsr305",
      ])


def com_squareup_okio():
  java_import_external(
      name="com_squareup_okio",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/com/squareup/okio/okio/1.13.0/okio-1.13.0.jar",
          "http://repo1.maven.org/maven2/com/squareup/okio/okio/1.13.0/okio-1.13.0.jar",
          "http://maven.ibiblio.org/maven2/com/squareup/okio/okio/1.13.0/okio-1.13.0.jar",
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
          "http://repo1.maven.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
          "http://maven.ibiblio.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
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


def net_bytebuddy():
  java_import_external(
      name="net_bytebuddy",
      jar_sha256=
      "c7f9861c94b07192ce970820cab9b046dc910e59567855d77c66d4532b9515bf",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.7.6/byte-buddy-1.7.6.jar",
          "http://repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.7.6/byte-buddy-1.7.6.jar",
          "http://maven.ibiblio.org/maven2/net/bytebuddy/byte-buddy/1.7.6/byte-buddy-1.7.6.jar",
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
      "617a8d75f66a57296255a13654a99f10f72f0964336e352211247ed046da3e94",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/net/java/dev/jna/jna/4.5.0/jna-4.5.0.jar",
          "http://maven.ibiblio.org/maven2/net/java/dev/jna/jna/4.5.0/jna-4.5.0.jar",
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna/4.5.0/jna-4.5.0.jar",
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
      "68ee6431c6c07dda48deaa2627c56beeea0dec5927fe7848983e06f7a6a76a08",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/net/java/dev/jna/jna-platform/4.5.0/jna-platform-4.5.0.jar",
          "http://repo1.maven.org/maven2/net/java/dev/jna/jna-platform/4.5.0/jna-platform-4.5.0.jar",
          "http://maven.ibiblio.org/maven2/net/java/dev/jna/jna-platform/4.5.0/jna-platform-4.5.0.jar",
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
      "f5408b3c74b43d87c0c121cab55f350c7b39d62f5f9db018fd261387c087130b",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.8/httpcore-4.4.8.jar",
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.8/httpcore-4.4.8.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpcore/4.4.8/httpcore-4.4.8.jar",
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
          "http://repo1.maven.org/maven2/org/apache/httpcomponents/httpmime/4.5.3/httpmime-4.5.3.jar",
          "http://maven.ibiblio.org/maven2/org/apache/httpcomponents/httpmime/4.5.3/httpmime-4.5.3.jar",
      ],
      licenses=["notice"],  # Apache License, Version 2.0
      deps=["@org_apache_httpcomponents_httpclient"])


def org_chromium_chromedriver():
  platform_http_file(
      name="org_chromium_chromedriver",
      amd64_sha256=
      "87d0059ab1579ec9c10ef34ab9817feea59e19a96c029d78349a57c36db5bb74",
      amd64_urls=[
          "http://chromedriver.storage.googleapis.com/2.33/chromedriver_linux64.zip"
      ],
      macos_sha256=
      "064b243c4236380cc705f183e100d7a229815db7b143f6ad3eaae072a48cc827",
      macos_urls=[
          "http://chromedriver.storage.googleapis.com/2.33/chromedriver_mac64.zip"
      ],
      windows_sha256=
      "e76941bf314e0c7967a8f0ccea10c331d69e4c1de0172fec33ed20df8c50e253",
      windows_urls=[
          "http://chromedriver.storage.googleapis.com/2.33/chromedriver_win32.zip"
      ])


def org_chromium_chromium():
  # Roughly corresponds to Chrome 58
  platform_http_file(
      name="org_chromium_chromium",
      amd64_sha256=
      "f585d71becaaa71f971e9c45ca9ef919b5a3fc7e51ab48a8e3f1d3a9ee705f42",
      amd64_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/488534/chrome-linux.zip"
      ],
      macos_sha256=
      "56195feefbe57139073429f2cfae33e010f6a73211ba66ce919c16c2c231b0fe",
      macos_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/488533/chrome-mac.zip"
      ],
      windows_sha256=
      "c0bb7a5ba47112d5f3789d2b0b9d3d2e38a5be592bf3332078fd90179c71dd27",
      windows_urls=[
          "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Win_x64/488538/chrome-win32.zip"
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
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/json/json/20170516/json-20170516.jar",
          "http://repo1.maven.org/maven2/org/json/json/20170516/json-20170516.jar",
          "http://maven.ibiblio.org/maven2/org/json/json/20170516/json-20170516.jar",
      ],
      licenses=["notice"]  # MIT-style license
  )


def org_mozilla_firefox():
  platform_http_file(
      name="org_mozilla_firefox",
      amd64_sha256=
      "9adf41f9c8ed525906d270c66ed2cfbd9e2154a3aeccaadbb0bdfb6ce3a3ca73",
      amd64_urls=[
          "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/56.0.1/linux-x86_64/en-US/firefox-56.0.1.tar.bz2",
          "https://ftp.mozilla.org/pub/firefox/releases/56.0.1/linux-x86_64/en-US/firefox-56.0.1.tar.bz2",
      ],
      macos_sha256=
      "b9e9f383fd12d6deb6b36b3c2844647ce065142e1eb934499559a8a9842d01ad",
      macos_urls=[
          "http://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/52.1.2esr/firefox-52.1.2esr.mac-x86_64.sdk.tar.bz2",
          "https://ftp.mozilla.org/pub/firefox/releases/52.1.2esr/firefox-52.1.2esr.mac-x86_64.sdk.tar.bz2",
      ])


def org_mozilla_geckodriver():
  platform_http_file(
      name="org_mozilla_geckodriver",
      amd64_sha256=
      "1c93b9cd82a28e4545829ae4686081ac7c76ba4e1f3faa0afda1fd3e5f6eda79",
      amd64_urls=[
          "http://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.19.0/geckodriver-v0.19.0-linux64.tar.gz",
          "https://github.com/mozilla/geckodriver/releases/download/v0.19.0/geckodriver-v0.19.0-linux64.tar.gz",
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
      sha256="9563021c5cba084041e13c218eb531d9c6920dcfa0ba1024bb5bf2b6c0df1797",
      strip_prefix="selenium-3.6.0",
      urls=[
          "http://mirror.bazel.build/pypi.python.org/packages/e1/25/ad1ee3c019e45933c201ae3c8b3c84ab335a64a8172051ace583b7371b35/selenium-3.6.0.tar.gz",
          "https://pypi.python.org/packages/e1/25/ad1ee3c019e45933c201ae3c8b3c84ab335a64a8172051ace583b7371b35/selenium-3.6.0.tar.gz"
      ])


def org_seleniumhq_selenium_api():
  java_import_external(
      name="org_seleniumhq_selenium_api",
      jar_sha256=
      "2f59e4ab0b63845eb5be785c18b1837b375b61ea3a6d29e7a97765456623e0ac",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.6.0/selenium-api-3.6.0.jar",
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.6.0/selenium-api-3.6.0.jar",
          "http://maven.ibiblio.org/maven2/org/seleniumhq/selenium/selenium-api/3.6.0/selenium-api-3.6.0.jar",
      ],
      licenses=["notice"],  # The Apache Software License, Version 2.0
      testonly_=1)


def org_seleniumhq_selenium_remote_driver():
  java_import_external(
      name="org_seleniumhq_selenium_remote_driver",
      jar_sha256=
      "71a3ee7d68018b5ecf9ffa80aeead6bc049fbf7286b74e1e0f1e4ecbe759da54",
      jar_urls=[
          "http://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.6.0/selenium-remote-driver-3.6.0.jar",
          "http://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.6.0/selenium-remote-driver-3.6.0.jar",
          "http://maven.ibiblio.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.6.0/selenium-remote-driver-3.6.0.jar",
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
    version_tuple += (int(number),)
  return version_tuple
