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
load("@bazel_gazelle//:deps.bzl", "go_repository")

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

def browser_repositories(firefox = False, chromium = False):
    """Sets up repositories for browsers defined in //browsers/....

    This should only be used on an experimental basis; projects should define
    their own browsers.

    Args:
        firefox: Configure repositories for //browsers:firefox-native.
        chromium: Configure repositories for //browsers:chromium-native.
    """
    if chromium:
        org_chromium_chromedriver()
        org_chromium_chromium()
    if firefox:
        org_mozilla_firefox()
        org_mozilla_geckodriver()

def bazel_skylib():
    native.http_archive(
        name = "bazel_skylib",
        sha256 = "d7cffbed034d1203858ca19ff2e88d241781f45652a4c719ed48eedc74bc82a9",
        strip_prefix = "bazel-skylib-0.3.1",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/archive/0.3.1.tar.gz",
            "https://github.com/bazelbuild/bazel-skylib/archive/0.3.1.tar.gz",
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
        sha256 = "12a849b4e9a08619233d4490a281aa2d34a69f9eaf85c2295f5357927e4d1763",
        strip_prefix = "context-1.1",
        urls = [
            "https://mirror.bazel.build/github.com/gorilla/context/archive/v1.1.tar.gz",
            "https://github.com/gorilla/context/archive/v1.1.tar.gz",
        ],
    )

def com_github_gorilla_mux():
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sha256 = "fe4d6909570b53121eb0d5e6f933ef7c49d5de094705af6ba07fab9c299df0f9",
        strip_prefix = "mux-1.6.1",
        urls = [
            "https://mirror.bazel.build/github.com/gorilla/mux/archive/v1.6.1.tar.gz",
            "https://github.com/gorilla/mux/archive/v1.6.1.tar.gz",
        ],
    )

def com_github_tebeka_selenium():
    go_repository(
        name = "com_github_tebeka_selenium",
        importpath = "github.com/tebeka/selenium",
        sha256 = "c731518d8f4724b1c581bdca90215978bf4d2658a5dc49b720d7c61294888396",
        strip_prefix = "selenium-a789e65b0e7f126888873e84f528c1c8537dff3e",
        urls = [
            "https://mirror.bazel.build/github.com/tebeka/selenium/archive/a789e65b0e7f126888873e84f528c1c8537dff3e.tar.gz",
            "https://github.com/tebeka/selenium/archive/a789e65b0e7f126888873e84f528c1c8537dff3e.tar.gz",
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
            "b7134929f7cc7c04021ec1cc27ef63ab907e410cf0588e397b8851181eb91092",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/code/gson/gson/2.8.2/gson-2.8.2.jar",
            "https://repo1.maven.org/maven2/com/google/code/gson/gson/2.8.2/gson-2.8.2.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
    )

def com_google_errorprone_error_prone_annotations():
    java_import_external(
        name = "com_google_errorprone_error_prone_annotations",
        jar_sha256 =
            "6ebd22ca1b9d8ec06d41de8d64e0596981d9607b42035f9ed374f9de271a481a",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.2.0/error_prone_annotations-2.2.0.jar",
            "https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.2.0/error_prone_annotations-2.2.0.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
    )

def com_google_guava():
    java_import_external(
        name = "com_google_guava",
        jar_sha256 =
            "31bfe27bdf9cba00cb4f3691136d3bc7847dfc87bfe772ca7a9eb68ff31d79f5",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/guava/guava/24.1-jre/guava-24.1-jre.jar",
            "https://repo1.maven.org/maven2/com/google/guava/guava/24.1-jre/guava-24.1-jre.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        exports = [
            "@com_google_code_findbugs_jsr305",
            "@com_google_errorprone_error_prone_annotations",
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
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/squareup/okio/okio/1.14.0/okio-1.14.0.jar",
            "https://repo1.maven.org/maven2/com/squareup/okio/okio/1.14.0/okio-1.14.0.jar",
        ],
        jar_sha256 =
            "4633c331f50642ebe795dc089d6a5928aff43071c9d17e7840a009eea2fe95a3",
        licenses = ["notice"],  # Apache 2.0
        deps = ["@com_google_code_findbugs_jsr305"],
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
        jar_sha256 =
            "1c7c222d5c317481538117e54029c289c5a1605a3cdcadf4e7f7cc1fe7469277",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.8.3/byte-buddy-1.8.3.jar",
            "https://repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.8.3/byte-buddy-1.8.3.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
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
            "7e97724443ad2a25ad8c73183431d47cc7946271bcbbdfa91a8a17522a566573",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.5/httpclient-4.5.5.jar",
            "https://repo1.maven.org/maven2/org/apache/httpcomponents/httpclient/4.5.5/httpclient-4.5.5.jar",
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
            "67fad24c4a85e3f33f51c97924a98b619722db15ce92dcd27484fb748af93e8e",
        amd64_urls = [
            "https://chromedriver.storage.googleapis.com/2.35/chromedriver_linux64.zip",
        ],
        macos_sha256 =
            "c11521bdc991874a0a29cf36beea6a4b8d73616ce6c8d1e6b90067d85718aa87",
        macos_urls = [
            "https://chromedriver.storage.googleapis.com/2.35/chromedriver_mac64.zip",
        ],
        windows_sha256 =
            "b32fdcc1c19bb829032f0447a8aac4f5436565ec1d0f105b63c4451ba4e6ae8a",
        windows_urls = [
            "https://chromedriver.storage.googleapis.com/2.35/chromedriver_win32.zip",
        ],
    )

def org_chromium_chromium():
    platform_http_file(
        name = "org_chromium_chromium",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        amd64_sha256 =
            "edb9807d40a57d235d8477beabe1dfa3d98e275312e7a48bc0cb9b44adb68236",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/564817/chrome-linux.zip"
        ],
        macos_sha256 =
            "b5a8641b187c623fad11ddccaa7f3053ac469a4a568ef8a593341846406ac965",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/564814/chrome-mac.zip"
        ],
        windows_sha256 =
            "1684732c817ce037fb22866ad579347c1eeebfb9a404155a78c02b783bcb1d06",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win_x64/564812/chrome-win32.zip"
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

def org_json():
    java_import_external(
        name = "org_json",
        jar_sha256 =
            "3eddf6d9d50e770650e62abe62885f4393aa911430ecde73ebafb1ffd2cfad16",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/json/json/20180130/json-20180130.jar",
            "https://repo1.maven.org/maven2/org/json/json/20180130/json-20180130.jar",
        ],
        licenses = ["notice"],  # MIT-style license
    )

def org_mozilla_firefox():
    platform_http_file(
        name = "org_mozilla_firefox",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "134fec04819eb56fa7b644cdd6d89623b21f4020bbedc3bd122db2a2caa4e434",
        amd64_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/58.0/linux-x86_64/en-US/firefox-58.0.tar.bz2",
            "https://ftp.mozilla.org/pub/firefox/releases/58.0/linux-x86_64/en-US/firefox-58.0.tar.bz2",
        ],
        macos_sha256 =
            "a853eb20821a21c0bedeb0263d7b5975e7704f20b78edfef129c73804b1fb962",
        macos_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/58.0/mac/en-US/Firefox%2058.0.dmg",
            "https://ftp.mozilla.org/pub/firefox/releases/58.0/mac/en-US/Firefox%2058.0.dmg",
        ],
    )

def org_mozilla_geckodriver():
    platform_http_file(
        name = "org_mozilla_geckodriver",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "7f55c4c89695fd1e6f8fc7372345acc1e2dbaa4a8003cee4bd282eed88145937",
        amd64_urls = [
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.19.1/geckodriver-v0.19.1-linux64.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.19.1/geckodriver-v0.19.1-linux64.tar.gz",
        ],
        macos_sha256 =
            "eb5a2971e5eb4a2fe74a3b8089f0f2cc96eed548c28526b8351f0f459c080836",
        macos_urls = [
            # TODO(fisherii): v0.19.1 is mirrored and ready to go.
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.16.1/geckodriver-v0.16.1-macos.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.16.1/geckodriver-v0.16.1-macos.tar.gz",
        ],
    )

def org_seleniumhq_py():
    native.new_http_archive(
        name = "org_seleniumhq_py",
        build_file = str(Label("//build_files:org_seleniumhq_py.BUILD")),
        sha256 = "5841fb30c3965866220c34d16de8e3d091e2833fcac385160a63db0c3522a297",
        strip_prefix = "selenium-3.11.0",
        urls = [
            "https://mirror.bazel.build/pypi.python.org/packages/d4/28/8124d32415bd3d67fabea52480395427576b582771283e89ce10a56d9e5b/selenium-3.11.0.tar.gz",
            "https://pypi.python.org/packages/d4/28/8124d32415bd3d67fabea52480395427576b582771283e89ce10a56d9e5b/selenium-3.11.0.tar.gz",
        ],
    )

def org_seleniumhq_selenium_api():
    java_import_external(
        name = "org_seleniumhq_selenium_api",
        jar_sha256 =
            "3e81810f61a1930d4ca868475cefdbe10b7260e352d09c4bfeda5e9ede5dd538",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.11.0/selenium-api-3.11.0.jar",
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.11.0/selenium-api-3.11.0.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
    )

def org_seleniumhq_selenium_remote_driver():
    java_import_external(
        name = "org_seleniumhq_selenium_remote_driver",
        jar_sha256 =
            "7a9c23a6a304bdcaec5a642f6641e07f798adc2f213cce456e4a829ad4454deb",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.11.0/selenium-remote-driver-3.11.0.jar",
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.11.0/selenium-remote-driver-3.11.0.jar",
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
