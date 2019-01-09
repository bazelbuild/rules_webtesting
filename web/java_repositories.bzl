# Copyright 2019 Google Inc.
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

load(":repositories.bzl", "should_create_repository")
load("@bazel_tools//tools/build_defs/repo:java.bzl", "java_import_external")

def java_repositories(**kwargs):
    """Defines external repositories required to use Java webtest and screenshotter APIs.

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
    if should_create_repository("org_seleniumhq_selenium_api", kwargs):
        org_seleniumhq_selenium_api()
    if should_create_repository("org_seleniumhq_selenium_remote_driver", kwargs):
        org_seleniumhq_selenium_remote_driver()
    if kwargs.keys():
        print("The following parameters are unknown: " + str(kwargs.keys()))

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
