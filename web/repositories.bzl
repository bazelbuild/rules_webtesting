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

load("//web/internal:platform_http_file.bzl", "platform_http_file")
load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:java.bzl", "java_import_external")

# NOTE: URLs are mirrored by an asynchronous review process. They must
#       be greppable for that to happen. It's OK to submit broken mirror
#       URLs, so long as they're correctly formatted. Bazel's downloader
#       has fast failover.

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

    Args:
        **kwargs: omit_... parameters used to prevent importing specific
          dependencies.
    """
    if should_create_repository("bazel_skylib", kwargs):
        bazel_skylib()
    if should_create_repository("com_github_blang_semver", kwargs):
        com_github_blang_semver()
    if should_create_repository("com_github_gorilla_context", kwargs):
        com_github_gorilla_context()
    if should_create_repository("com_github_gorilla_mux", kwargs):
        com_github_gorilla_mux()
    if should_create_repository("com_github_tebeka_selenium", kwargs):
        com_github_tebeka_selenium()
    if should_create_repository("com_github_urllib3", kwargs):
        com_github_urllib3()
    if should_create_repository("com_google_code_findbugs_jsr305", kwargs):
        com_google_code_findbugs_jsr305()
    if should_create_repository("com_google_code_gson", kwargs):
        com_google_code_gson()
    if should_create_repository(
        "com_google_errorprone_error_prone_annotations",
        kwargs,
    ):
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
    if should_create_repository("org_apache_commons_exec", kwargs):
        org_apache_commons_exec()
    if should_create_repository("org_apache_httpcomponents_httpclient", kwargs):
        org_apache_httpcomponents_httpclient()
    if should_create_repository("org_apache_httpcomponents_httpcore", kwargs):
        org_apache_httpcomponents_httpcore()
    if should_create_repository("org_hamcrest_core", kwargs):
        org_hamcrest_core()
    if should_create_repository("org_jetbrains_kotlin_stdlib", kwargs):
        org_jetbrains_kotlin_stdlib()
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

    Args:
        name: The name of the repository that should be checked.
        args: A dictionary that contains "omit_...": bool pairs.

    Returns:
        boolean indicating whether the repository should be created.
    """
    key = "omit_" + name
    if key in args:
        val = args.pop(key)
        if val:
            return False
    if native.existing_rule(name):
        return False
    return True

def browser_repositories(firefox = False, chromium = False, sauce = False):
    """Sets up repositories for browsers defined in //browsers/....

    This should only be used on an experimental basis; projects should define
    their own browsers.

    Args:
        firefox: Configure repositories for //browsers:firefox-native.
        chromium: Configure repositories for //browsers:chromium-native.
        sauce: Configure repositories for //browser/sauce:chrome-win10.
    """
    if chromium:
        org_chromium_chromedriver()
        org_chromium_chromium()
    if firefox:
        org_mozilla_firefox()
        org_mozilla_geckodriver()
    if sauce:
        com_saucelabs_sauce_connect()

def bazel_skylib():
    http_archive(
        name = "bazel_skylib",
        sha256 = "",
        strip_prefix = "bazel-skylib-e9fc4750d427196754bebb0e2e1e38d68893490a",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/archive/e9fc4750d427196754bebb0e2e1e38d68893490a.tar.gz",
            "https://github.com/bazelbuild/bazel-skylib/archive/e9fc4750d427196754bebb0e2e1e38d68893490a.tar.gz",
        ],
    )

def com_github_blang_semver():
    go_repository(
        name = "com_github_blang_semver",
        importpath = "github.com/blang/semver",
        sha256 = "3d9da53f4c2d3169bfa9b25f2f36f301a37556a47259c870881524c643c69c57",
        strip_prefix = "semver-3.5.1",
        urls = [
            "https://mirror.bazel.build/github.com/blang/semver/archive/v3.5.1.tar.gz",
            "https://github.com/blang/semver/archive/v3.5.1.tar.gz",
        ],
    )

def com_github_gorilla_context():
    go_repository(
        name = "com_github_gorilla_context",
        importpath = "github.com/gorilla/context",
        sha256 = "2dfdd051c238695bf9ebfed0bf6a8c533507ac0893bce23be5930e973736bb03",
        strip_prefix = "context-1.1.1",
        urls = [
            "https://mirror.bazel.build/github.com/gorilla/context/archive/v1.1.1.tar.gz",
            "https://github.com/gorilla/context/archive/v1.1.1.tar.gz",
        ],
    )

def com_github_gorilla_mux():
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sha256 = "0dc18fb09413efea7393e9c2bd8b5b442ce08e729058f5f7e328d912c6c3d3e3",
        strip_prefix = "mux-1.6.2",
        urls = [
            "https://mirror.bazel.build/github.com/gorilla/mux/archive/v1.6.2.tar.gz",
            "https://github.com/gorilla/mux/archive/v1.6.2.tar.gz",
        ],
    )

def com_github_tebeka_selenium():
    go_repository(
        name = "com_github_tebeka_selenium",
        importpath = "github.com/tebeka/selenium",
        sha256 = "c506637fd690f4125136233a3ea405908b8255e2d7aa2aa9d3b746d96df50dcd",
        strip_prefix = "selenium-a49cf4b98a36c2b21b1ccb012852bd142d5fc04a",
        urls = [
            "https://mirror.bazel.build/github.com/tebeka/selenium/archive/a49cf4b98a36c2b21b1ccb012852bd142d5fc04a.tar.gz",
            "https://github.com/tebeka/selenium/archive/a49cf4b98a36c2b21b1ccb012852bd142d5fc04a.tar.gz",
        ],
    )

def com_github_urllib3():
    http_archive(
        name = "com_github_urllib3",
        build_file = str(Label("//build_files:com_github_urllib3.BUILD")),
        sha256 = "a68ac5e15e76e7e5dd2b8f94007233e01effe3e50e8daddf69acfd81cb686baf",
        strip_prefix = "urllib3-1.23",
        urls = [
            "https://files.pythonhosted.org/packages/3c/d2/dc5471622bd200db1cd9319e02e71bc655e9ea27b8e0ce65fc69de0dac15/urllib3-1.23.tar.gz",
        ],
    )

def com_google_code_findbugs_jsr305():
    java_import_external(
        name = "com_google_code_findbugs_jsr305",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar",
            "https://repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar",
        ],
        jar_sha256 =
            "766ad2a0783f2687962c8ad74ceecc38a28b9f72a2d085ee438b7813e928d0c7",
        licenses = ["notice"],  # BSD 3-clause
    )

def com_google_code_gson():
    java_import_external(
        name = "com_google_code_gson",
        jar_sha256 =
            "233a0149fc365c9f6edbd683cfe266b19bdc773be98eabdaf6b3c924b48e7d81",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/code/gson/gson/2.8.5/gson-2.8.5.jar",
            "https://repo1.maven.org/maven2/com/google/code/gson/gson/2.8.5/gson-2.8.5.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
    )

def com_google_errorprone_error_prone_annotations():
    java_import_external(
        name = "com_google_errorprone_error_prone_annotations",
        jar_sha256 =
            "10a5949aa0f95c8de4fd47edfe20534d2acefd8c224f8afea1f607e112816120",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.3.1/error_prone_annotations-2.3.1.jar",
            "https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.3.1/error_prone_annotations-2.3.1.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
    )

def com_google_guava():
    java_import_external(
        name = "com_google_guava",
        jar_sha256 = "a0e9cabad665bc20bcd2b01f108e5fc03f756e13aea80abaadb9f407033bea2c",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/guava/guava/26.0-jre/guava-26.9-jre.jar",
            "https://repo1.maven.org/maven2/com/google/guava/guava/26.0-jre/guava-26.0-jre.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        exports = [
            "@com_google_code_findbugs_jsr305",
            "@com_google_errorprone_error_prone_annotations",
        ],
    )

def com_saucelabs_sauce_connect():
    platform_http_file(
        name = "com_saucelabs_sauce_connect",
        licenses = ["by_exception_only"],  # SauceLabs EULA
        amd64_sha256 = "dd53f2cdcec489fbc2443942b853b51bf44af39f230600573119cdd315ddee52",
        amd64_urls = [
            "https://saucelabs.com/downloads/sc-4.5.1-linux.tar.gz",
        ],
        macos_sha256 = "920ae7bd5657bccdcd27bb596593588654a2820486043e9a12c9062700697e66",
        macos_urls = [
            "https://saucelabs.com/downloads/sc-4.5.1-osx.zip",
        ],
    )

def com_squareup_okhttp3_okhttp():
    java_import_external(
        name = "com_squareup_okhttp3_okhttp",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/3.9.1/okhttp-3.9.1.jar",
            "https://repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/3.9.1/okhttp-3.9.1.jar",
        ],
        jar_sha256 =
            "a0d01017a42bba26e507fc6d448bb36e536f4b6e612f7c42de30bbdac2b7785e",
        licenses = ["notice"],  # Apache 2.0
        deps = [
            "@com_squareup_okio",
            "@com_google_code_findbugs_jsr305",
        ],
    )

def com_squareup_okio():
    java_import_external(
        name = "com_squareup_okio",
        jar_sha256 = "79b948cf77504750fdf7aeaf362b5060415136ab6635e5113bd22925e0e9e737",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/squareup/okio/okio/2.0.0/okio-2.0.0.jar",
            "https://repo1.maven.org/maven2/com/squareup/okio/okio/2.0.0/okio-2.0.0.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        deps = [
            "@com_google_code_findbugs_jsr305",
            "@org_jetbrains_kotlin_stdlib",
        ],
    )

def commons_codec():
    java_import_external(
        name = "commons_codec",
        jar_sha256 =
            "e599d5318e97aa48f42136a2927e6dfa4e8881dff0e6c8e3109ddbbff51d7b7d",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/commons-codec/commons-codec/1.11/commons-codec-1.11.jar",
            "https://repo1.maven.org/maven2/commons-codec/commons-codec/1.11/commons-codec-1.11.jar",
        ],
        licenses = ["notice"],  # Apache License, Version 2.0
    )

def commons_logging():
    java_import_external(
        name = "commons_logging",
        jar_sha256 =
            "daddea1ea0be0f56978ab3006b8ac92834afeefbd9b7e4e6316fca57df0fa636",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
            "https://repo1.maven.org/maven2/commons-logging/commons-logging/1.2/commons-logging-1.2.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
    )

def junit():
    java_import_external(
        name = "junit",
        jar_sha256 =
            "59721f0805e223d84b90677887d9ff567dc534d7c502ca903c0c2b17f05c116a",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
            "https://repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
        ],
        licenses = ["reciprocal"],  # Eclipse Public License 1.0
        testonly_ = 1,
        deps = ["@org_hamcrest_core"],
    )

def net_bytebuddy():
    java_import_external(
        name = "net_bytebuddy",
        jar_sha256 = "4b87ad52a8f64a1197508e176e84076584160e3d65229ff757efee870cd4a8e2",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.8.19/byte-buddy-1.8.19.jar",
            "https://repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.8.19/byte-buddy-1.8.19.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        deps = ["@com_google_code_findbugs_jsr305"],
    )

def org_apache_commons_exec():
    java_import_external(
        name = "org_apache_commons_exec",
        jar_sha256 =
            "cb49812dc1bfb0ea4f20f398bcae1a88c6406e213e67f7524fb10d4f8ad9347b",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
            "https://repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
        ],
        licenses = ["notice"],  # Apache License, Version 2.0
    )

def org_apache_httpcomponents_httpclient():
    java_import_external(
        name = "org_apache_httpcomponents_httpclient",
        jar_sha256 =
            "c03f813195e7a80e3608d0ddd8da80b21696a4c92a6a2298865bf149071551c7",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.6/httpclient-4.5.6.jar",
            "https://repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.6/httpclient-4.5.6.jar",
        ],
        licenses = ["notice"],  # Apache License, Version 2.0
        deps = [
            "@org_apache_httpcomponents_httpcore",
            "@commons_logging",
            "@commons_codec",
        ],
    )

def org_apache_httpcomponents_httpcore():
    java_import_external(
        name = "org_apache_httpcomponents_httpcore",
        jar_sha256 =
            "1b4a1c0b9b4222eda70108d3c6e2befd4a6be3d9f78ff53dd7a94966fdf51fc5",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.9/httpcore-4.4.9.jar",
            "https://repo1.maven.org/maven2/org/apache/httpcomponents/httpcore/4.4.9/httpcore-4.4.9.jar",
        ],
        licenses = ["notice"],  # Apache License, Version 2.0
    )

def org_chromium_chromedriver():
    platform_http_file(
        name = "org_chromium_chromedriver",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        amd64_sha256 =
            "71eafe087900dbca4bc0b354a1d172df48b31a4a502e21f7c7b156d7e76c95c7",
        amd64_urls = [
            "https://chromedriver.storage.googleapis.com/2.41/chromedriver_linux64.zip",
        ],
        macos_sha256 =
            "fd32a27148f44796a55f5ce3397015c89ebd9f600d9dda2bcaca54575e2497ae",
        macos_urls = [
            "https://chromedriver.storage.googleapis.com/2.41/chromedriver_mac64.zip",
        ],
    )

def org_chromium_chromium():
    platform_http_file(
        name = "org_chromium_chromium",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        amd64_sha256 =
            "6933d0afce6e17304b62029fbbd246cbe9e130eb0d90d7682d3765d3dbc8e1c8",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/561732/chrome-linux.zip",
        ],
        macos_sha256 =
            "084884e91841a923d7b6e81101f0105bbc3b0026f9f6f7a3477f5b313ee89e32",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/561733/chrome-mac.zip",
        ],
    )

def org_hamcrest_core():
    java_import_external(
        name = "org_hamcrest_core",
        jar_sha256 =
            "66fdef91e9739348df7a096aa384a5685f4e875584cce89386a7a47251c4d8e9",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
            "https://repo1.maven.org/maven2/org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar",
        ],
        licenses = ["notice"],  # New BSD License
        testonly_ = 1,
    )

def org_jetbrains_kotlin_stdlib():
    java_import_external(
        name = "org_jetbrains_kotlin_stdlib",
        jar_sha256 = "62eaf9cc6e746cef4593abe7cdb4dd48694ef5f817c852e0d9fbbd11fcfc564e",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.2.61/kotlin-stdlib-1.2.61.jar",
            "https://repo1.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.2.61/kotlin-stdlib-1.2.61.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
    )

def org_json():
    java_import_external(
        name = "org_json",
        jar_sha256 = "518080049ba83181914419d11a25d9bc9833a2d729b6a6e7469fa52851356da8",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/json/json/20180813/json-20180813.jar",
            "https://repo1.maven.org/maven2/org/json/json/20180813/json-20180813.jar",
        ],
        licenses = ["notice"],  # MIT-style license
    )

def org_mozilla_firefox():
    platform_http_file(
        name = "org_mozilla_firefox",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "3a729ddcb1e0f5d63933177a35177ac6172f12edbf9fbbbf45305f49333608de",
        amd64_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/61.0.2/linux-x86_64/en-US/firefox-61.0.2.tar.bz2",
            "https://ftp.mozilla.org/pub/firefox/releases/61.0.2/linux-x86_64/en-US/firefox-61.0.2.tar.bz2",
        ],
        macos_sha256 =
            "bf23f659ae34832605dd0576affcca060d1077b7bf7395bc9874f62b84936dc5",
        macos_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/61.0.2/mac/en-US/Firefox%2061.0.2.dmg",
            "https://ftp.mozilla.org/pub/firefox/releases/61.0.2/mac/en-US/Firefox%2061.0.2.dmg",
        ],
    )

def org_mozilla_geckodriver():
    platform_http_file(
        name = "org_mozilla_geckodriver",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "c9ae92348cf00aa719be6337a608fae8304691a95668e8e338d92623ba9e0ec6",
        amd64_urls = [
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-linux64.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-linux64.tar.gz",
        ],
        macos_sha256 =
            "ce4a3e9d706db94e8760988de1ad562630412fa8cf898819572522be584f01ce",
        macos_urls = [
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-macos.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.21.0/geckodriver-v0.21.0-macos.tar.gz",
        ],
    )

def org_seleniumhq_py():
    http_archive(
        name = "org_seleniumhq_py",
        build_file = str(Label("//build_files:org_seleniumhq_py.BUILD")),
        sha256 = "f9ca21919b564a0a86012cd2177923e3a7f37c4a574207086e710192452a7c40",
        strip_prefix = "selenium-3.14.0",
        urls = [
            "https://files.pythonhosted.org/packages/af/7c/3f76140976b1c8f8a6b437ccd1f04efaed37bdc2600530e76ba981c677b9/selenium-3.14.0.tar.gz",
        ],
    )

def org_seleniumhq_selenium_api():
    java_import_external(
        name = "org_seleniumhq_selenium_api",
        jar_sha256 = "1fc941f86ba4fefeae9a705c1468e65beeaeb63688e19ad3fcbda74cc883ee5b",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.14.0/selenium-api-3.14.0.jar",
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.14.0/selenium-api-3.14.0.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
    )

def org_seleniumhq_selenium_remote_driver():
    java_import_external(
        name = "org_seleniumhq_selenium_remote_driver",
        jar_sha256 =
            "284cb4ea043539353bd5ecd774cbd726b705d423ea4569376c863d0b66e5eaf2",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.14.0/selenium-remote-driver-3.14.0.jar",
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.14.0/selenium-remote-driver-3.14.0.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
        deps = [
            "@com_google_code_gson",
            "@com_google_guava",
            "@net_bytebuddy",
            "@com_squareup_okhttp3_okhttp",
            "@com_squareup_okio",
            "@commons_codec",
            "@commons_logging",
            "@org_apache_commons_exec",
            "@org_apache_httpcomponents_httpclient",
            "@org_apache_httpcomponents_httpcore",
            "@org_seleniumhq_selenium_api",
        ],
    )
