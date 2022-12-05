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
    this one (e.g. com_google_guava_guava.) Alternatively, an allowlist model may be
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
    if should_create_repository("com_google_guava_guava", kwargs):
        com_google_guava_guava()
    if should_create_repository("com_squareup_okhttp3_okhttp", kwargs):
        com_squareup_okhttp3_okhttp()
    if should_create_repository("com_squareup_okio_okio", kwargs):
        com_squareup_okio_okio()
    if should_create_repository("junit_junit", kwargs):
        junit()
    if should_create_repository("net_bytebuddy_byte_buddy", kwargs):
        net_bytebuddy_byte_buddy()
    if should_create_repository("org_apache_commons_commons_exec", kwargs):
        org_apache_commons_commons_exec()
    if should_create_repository("org_hamcrest_hamcrest_core", kwargs):
        org_hamcrest_hamcrest_core()
    if should_create_repository("org_jetbrains_kotlin_kotlin_stdlib", kwargs):
        org_jetbrains_kotlin_kotlin_stdlib()
    if should_create_repository("org_json_json", kwargs):
        org_json_json()
    if should_create_repository("org_seleniumhq_selenium_selenium_api", kwargs):
        org_seleniumhq_selenium_selenium_api()
    if should_create_repository("org_seleniumhq_selenium_selenium_remote_driver", kwargs):
        org_seleniumhq_selenium_selenium_remote_driver()
    if should_create_repository("org_seleniumhq_selenium_selenium_support", kwargs):
        org_seleniumhq_selenium_selenium_support()
    if kwargs.keys():
        print("The following parameters are unknown: " + str(kwargs.keys()))

def com_google_code_findbugs_jsr305():
    java_import_external(
        name = "com_google_code_findbugs_jsr305",
        jar_sha256 = "766ad2a0783f2687962c8ad74ceecc38a28b9f72a2d085ee438b7813e928d0c7",
        jar_urls = [
            "https://repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar",
        ],
        licenses = ["notice"],  # BSD 3-clause
    )

def com_google_errorprone_error_prone_annotations():
    java_import_external(
        name = "com_google_errorprone_error_prone_annotations",
        jar_sha256 = "ec59f1b702d9afc09e8c3929f5c42777dec623a6ea2731ac694332c7d7680f5a",
        jar_urls = [
            "https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.3.3/error_prone_annotations-2.3.3.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
    )

def com_google_guava_guava():
    java_import_external(
        name = "com_google_guava_guava",
        jar_sha256 = "73e4d6ae5f0e8f9d292a4db83a2479b5468f83d972ac1ff36d6d0b43943b4f91",
        jar_urls = [
            "https://repo1.maven.org/maven2/com/google/guava/guava/28.0-jre/guava-28.0-jre.jar",
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
        jar_sha256 = "20f483a62087faa1dc8240150fa500b0a42c822735a12481ae32c5238d9922cc",
        jar_urls = [
            "https://repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/4.1.0/okhttp-4.1.0.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        deps = [
            "@com_squareup_okio_okio",
            "@com_google_code_findbugs_jsr305",
        ],
    )

def com_squareup_okio_okio():
    java_import_external(
        name = "com_squareup_okio_okio",
        jar_sha256 = "1c52079b6159b096181a2fad4df7f15423ee6c66266d1dcb0264bf37c58178b0",
        jar_urls = [
            "https://repo1.maven.org/maven2/com/squareup/okio/okio/2.3.0/okio-2.3.0.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        deps = [
            "@com_google_code_findbugs_jsr305",
            "@org_jetbrains_kotlin_kotlin_stdlib",
        ],
    )

def junit():
    java_import_external(
        name = "junit_junit",
        jar_sha256 = "59721f0805e223d84b90677887d9ff567dc534d7c502ca903c0c2b17f05c116a",
        jar_urls = [
            "https://repo1.maven.org/maven2/junit/junit/4.12/junit-4.12.jar",
        ],
        licenses = ["reciprocal"],  # Eclipse Public License 1.0
        testonly_ = 1,
        deps = ["@org_hamcrest_hamcrest_core"],
    )

def net_bytebuddy_byte_buddy():
    java_import_external(
        name = "net_bytebuddy_byte_buddy",
        jar_sha256 = "6b71e4f70c96b67d420f592148aa4fd1966aba458b35d11f491ff13de97dc862",
        jar_urls = [
            "https://repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.9.16/byte-buddy-1.9.16.jar",
        ],
        licenses = ["notice"],  # Apache 2.0
        deps = ["@com_google_code_findbugs_jsr305"],
    )

def org_apache_commons_commons_exec():
    java_import_external(
        name = "org_apache_commons_commons_exec",
        jar_sha256 =
            "cb49812dc1bfb0ea4f20f398bcae1a88c6406e213e67f7524fb10d4f8ad9347b",
        jar_urls = [
            "https://repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar",
        ],
        licenses = ["notice"],  # Apache License, Version 2.0
    )

def org_hamcrest_hamcrest_core():
    java_import_external(
        name = "org_hamcrest_hamcrest_core",
        jar_sha256 = "e09109e54a289d88506b9bfec987ddd199f4217c9464132668351b9a4f00bee9",
        jar_urls = [
            "https://repo1.maven.org/maven2/org/hamcrest/hamcrest-core/2.1/hamcrest-core-2.1.jar",
        ],
        licenses = ["notice"],  # New BSD License
        testonly_ = 1,
    )

def org_jetbrains_kotlin_kotlin_stdlib():
    java_import_external(
        name = "org_jetbrains_kotlin_kotlin_stdlib",
        jar_sha256 = "6ea3d0921b26919b286f05cbdb906266666a36f9a7c096197114f7495708ffbc",
        jar_urls = [
            "https://repo1.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.3.41/kotlin-stdlib-1.3.41.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
    )

def org_json_json():
    java_import_external(
        name = "org_json_json",
        jar_sha256 = "518080049ba83181914419d11a25d9bc9833a2d729b6a6e7469fa52851356da7",
        jar_urls = [
            "https://repo1.maven.org/maven2/org/json/json/20190722/json-20190722.jar",
        ],
        licenses = ["notice"],  # MIT-style license
    )

def org_seleniumhq_selenium_selenium_api():
    java_import_external(
        name = "org_seleniumhq_selenium_selenium_api",
        jar_sha256 = "8bfd5a736eccfc08866301ffc9b7f529e55976355c5799bed8392486df64dee5",
        jar_urls = [
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.141.59/selenium-api-3.141.59.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
    )

def org_seleniumhq_selenium_selenium_remote_driver():
    java_import_external(
        name = "org_seleniumhq_selenium_selenium_remote_driver",
        jar_sha256 = "9829fe57adf36743d785d0c2e7db504ba3ba0a3aacac652b8867cc854d2dfc45",
        jar_urls = [
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.141.59/selenium-remote-driver-3.141.59.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
        deps = [
            "@com_google_guava_guava",
            "@net_bytebuddy_byte_buddy",
            "@com_squareup_okhttp3_okhttp",
            "@com_squareup_okio_okio",
            "@org_apache_commons_commons_exec",
            "@org_seleniumhq_selenium_selenium_api",
        ],
    )

def org_seleniumhq_selenium_selenium_support():
    java_import_external(
        name = "org_seleniumhq_selenium_selenium_support",
        jar_sha256 = "2c74196d15277ce6003454d72fc3434091dbf3ba65060942719ba551509404d8",
        jar_urls = [
            "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-support/3.141.59/selenium-support-3.141.59.jar",
        ],
        licenses = ["notice"],  # The Apache Software License, Version 2.0
        testonly_ = 1,
        deps = [
            "@com_google_guava_guava",
            "@net_bytebuddy_byte_buddy",
            "@com_squareup_okhttp3_okhttp",
            "@com_squareup_okio_okio",
            "@org_apache_commons_commons_exec",
            "@org_seleniumhq_selenium_selenium_api",
            "@org_seleniumhq_selenium_selenium_remote_driver",
        ],
    )

# Artifacts for rules_jvm_external. See:
# https://github.com/bazelbuild/rules_jvm_external#exporting-and-consuming-artifacts-from-external-repositories
RULES_WEBTESTING_ARTIFACTS = [
    "org.seleniumhq.selenium:selenium-api:3.141.59",
    "org.seleniumhq.selenium:selenium-remote-driver:3.141.59",
    "com.google.guava:guava:28.0-jre",
]
