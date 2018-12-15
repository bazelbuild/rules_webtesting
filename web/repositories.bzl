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
    if should_create_repository("junit", kwargs):
        junit()
    if should_create_repository("net_bytebuddy", kwargs):
        net_bytebuddy()
    if should_create_repository("org_apache_commons_exec", kwargs):
        org_apache_commons_exec()
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
        sha256 = "68ef2998919a92c2c9553f7a6b00a1d0615b57720a13239c0e51d0ded5aa452a",
        strip_prefix = "bazel-skylib-8cecf885c8bf4c51e82fd6b50b9dd68d2c98f757",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/archive/8cecf885c8bf4c51e82fd6b50b9dd68d2c98f757.tar.gz",
            "https://github.com/bazelbuild/bazel-skylib/archive/8cecf885c8bf4c51e82fd6b50b9dd68d2c98f757.tar.gz",
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
        sha256 = "4f0e1b9f54825580707fca637a4b174872dce5da6685c010b98eff45c1b3064a",
        strip_prefix = "selenium-edf31bb7fd715ad505d9190f8d65d13f39a7c825",
        urls = [
            "https://mirror.bazel.build/github.com/tebeka/selenium/archive/edf31bb7fd715ad505d9190f8d65d13f39a7c825.tar.gz",
            "https://github.com/tebeka/selenium/archive/edf31bb7fd715ad505d9190f8d65d13f39a7c825.tar.gz",
        ],
    )

def com_github_urllib3():
    http_archive(
        name = "com_github_urllib3",
        build_file = str(Label("//build_files:com_github_urllib3.BUILD")),
        sha256 = "de9529817c93f27c8ccbfead6985011db27bd0ddfcdb2d86f3f663385c6a9c22",
        strip_prefix = "urllib3-1.24.1",
        urls = [
            "https://files.pythonhosted.org/packages/b1/53/37d82ab391393565f2f831b8eedbffd57db5a718216f82f1a8b4d381a1c1/urllib3-1.24.1.tar.gz",
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

def com_google_errorprone_error_prone_annotations():
    java_import_external(
        name = "com_google_errorprone_error_prone_annotations",
        jar_sha256 =
            "357cd6cfb067c969226c442451502aee13800a24e950fdfde77bcdb4565a668d",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.3.2/error_prone_annotations-2.3.2.jar",
            "https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.3.2/error_prone_annotations-2.3.2.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
    )

def com_google_guava():
    java_import_external(
        name = "com_google_guava",
        jar_sha256 = "63b09db6861011e7fb2481be7790c7fd4b03f0bb884b3de2ecba8823ad19bf3f",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/google/guava/guava/27.0-jre/guava-27.0-jre.jar",
            "https://repo1.maven.org/maven2/com/google/guava/guava/27.0-jre/guava-27.0-jre.jar",
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
        windows_sha256 =
            "ec11b4ee029c9f0cba316820995df6ab5a4f394053102e1871b9f9589d0a9eb5",
        windows_urls = [
            "https://saucelabs.com/downloads/sc-4.4.12-win32.zip",
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
        jar_sha256 = "6773b78e61ed63b9dbb36c87f11873883f2cc3539c8be2a9568091248d83b2a2",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/com/squareup/okio/okio/2.1.0/okio-2.1.0.jar",
            "https://repo1.maven.org/maven2/com/squareup/okio/okio/2.1.0/okio-2.1.0.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        deps = [
            "@com_google_code_findbugs_jsr305",
            "@org_jetbrains_kotlin_stdlib",
        ],
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

def org_chromium_chromedriver():
    platform_http_file(
        name = "org_chromium_chromedriver",
        licenses = ["reciprocal"],  # BSD 3-clause, ICU, MPL 1.1, libpng (BSD/MIT-like), Academic Free License v. 2.0, BSD 2-clause, MIT
        amd64_sha256 =
            "d4a5eec0a3b7fec9bcb71353233dde38630e51b29fa7b218cdd196e2e4487da7",
        amd64_urls = [
            "https://chromedriver.storage.googleapis.com/2.45/chromedriver_linux64.zip",
        ],
        macos_sha256 =
            "aa0f416a48e20185da62525869c1f98f994bf99f241d6ce2eb1af6ceb517c425",
        macos_urls = [
            "https://chromedriver.storage.googleapis.com/2.45/chromedriver_mac64.zip",
        ],
        windows_sha256 =
            "8f3373a260a524410e25ea643ecb0175ed49a078088c7ab4d88db323ee19a230",
        windows_urls = [
            "https://chromedriver.storage.googleapis.com/2.45/chromedriver_win32.zip",
        ],
    )

def org_chromium_chromium():
    platform_http_file(
        name = "org_chromium_chromium",
        licenses = ["notice"],  # BSD 3-clause (maybe more?)
        amd64_sha256 =
            "a83baae72d70d4c337c7af6fb190ec3fbb082b94af1e16c5f28390dfd7771612",
        amd64_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/587811/chrome-linux.zip",
        ],
        macos_sha256 =
            "2878816a49e2eeaac4f9e689d05668eb51d0e5e0d1c071756d3a4dfcf35cfcd0",
        macos_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac/587811/chrome-mac.zip",
        ],
        windows_sha256 =
            "fe5c6fc33db8d69c3270a22d6569fe5e9f2afb24cb0b7a0a451793c521ab6eef",
        windows_urls = [
            "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win_x64/587811/chrome-win32.zip",
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
        jar_sha256 = "4ff0fcb97f4983b4aaba12668c24ad21b08460915db1b021d8f1d8bee687f21c",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.3.0/kotlin-stdlib-1.3.0.jar",
            "https://repo1.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.3.0/kotlin-stdlib-1.3.0.jar",
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
            "d4c696d77823bb3a3ea24d67d0f1075899fbaa4a0893d069f606015708ad1fca",
        amd64_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/63.0.1/linux-x86_64/en-US/firefox-63.0.1.tar.bz2",
            "https://ftp.mozilla.org/pub/firefox/releases/63.0.1/linux-x86_64/en-US/firefox-63.0.1.tar.bz2",
        ],
        macos_sha256 =
            "73e810389dcbf3b0ac74d0997121d498789d5cd205da16175c4ca5fca2ca8f79",
        macos_urls = [
            "https://mirror.bazel.build/ftp.mozilla.org/pub/firefox/releases/63.0.1/mac/en-US/Firefox%2063.0.1.dmg",
            "https://ftp.mozilla.org/pub/firefox/releases/63.0.1/mac/en-US/Firefox%2063.0.1.dmg",
        ],
    )

def org_mozilla_geckodriver():
    platform_http_file(
        name = "org_mozilla_geckodriver",
        licenses = ["reciprocal"],  # MPL 2.0
        amd64_sha256 =
            "2abf02cb69b48f2ba65ea344b752ff547e5431659aad80b03bf68cdb4f8df14b",
        amd64_urls = [
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-linux64.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-linux64.tar.gz",
        ],
        macos_sha256 =
            "006e206cc4c93ad9ef857aa5b7efc5a9084fa239381a7afaaa4acbba6f00bac9",
        macos_urls = [
            "https://mirror.bazel.build/github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-macos.tar.gz",
            "https://github.com/mozilla/geckodriver/releases/download/v0.23.0/geckodriver-v0.23.0-macos.tar.gz",
        ],
    )

def org_seleniumhq_py():
    http_archive(
        name = "org_seleniumhq_py",
        build_file = str(Label("//build_files:org_seleniumhq_py.BUILD")),
        sha256 = "deaf32b60ad91a4611b98d8002757f29e6f2c2d5fcaf202e1c9ad06d6772300d",
        strip_prefix = "selenium-3.141.0",
        urls = [
            "https://files.pythonhosted.org/packages/ed/9c/9030520bf6ff0b4c98988448a93c04fcbd5b13cd9520074d8ed53569ccfe/selenium-3.141.0.tar.gz",
        ],
    )

def org_seleniumhq_selenium_api():
    java_import_external(
        name = "org_seleniumhq_selenium_api",
        jar_sha256 = "c9d8907216650cffac1526fa40caab840ae6fdbe901ce3d4576a435d54dd41fa",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.141.5/selenium-api-3.141.5.jar",
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.141.5/selenium-api-3.141.5.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
    )

def org_seleniumhq_selenium_remote_driver():
    java_import_external(
        name = "org_seleniumhq_selenium_remote_driver",
        jar_sha256 =
            "fe144c413fba8dcf5dc490cfd063588758b2ac8a960e3d2117b06f28a16d04b5",
        jar_urls = [
            "https://mirror.bazel.build/repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.141.5/selenium-remote-driver-3.141.5.jar",
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.141.5/selenium-remote-driver-3.141.5.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
        deps = [
            "@com_google_guava",
            "@net_bytebuddy",
            "@com_squareup_okhttp3_okhttp",
            "@com_squareup_okio",
            "@org_apache_commons_exec",
            "@org_seleniumhq_selenium_api",
        ],
    )
