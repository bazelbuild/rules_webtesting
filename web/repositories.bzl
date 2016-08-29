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

def web_test_repositories(
    prefix="@io_bazel_rules_web",
    java=False,
    go=False,
    launcher="@io_bazel_rules_web//launcher:main",
    merger="@io_bazel_rules_web//metadata:merger",
    default_config="@io_bazel_rules_web//rules:default"):
  native.new_git_repository(
      name="com_github_gorilla_mux",
      build_file=prefix + "//:BUILD.gorilla_mux",
      commit="cf79e51a62d8219d52060dfc1b4e810414ba2d15",
      remote="https://github.com/gorilla/mux.git",)

  native.http_jar(
      name="org_seleniumhq_server",
      sha256="df874ce5b9508ac9f4ee0a3f50290836915c837b68975066a3841e839bc39804",
      url="http://selenium-release.storage.googleapis.com/3.0-beta2/selenium-server-standalone-3.0.0-beta2.jar",
  )

  native.bind(
      name="web_test_launcher",
      actual=launcher,)

  native.bind(
      name="web_test_merger",
      actual=merger,)

  native.bind(
      name="web_test_default_config",
      actual=default_config,)

  if java:
    native.new_http_archive(
        name="org_seleniumhq_java",
        build_file=prefix + "//:BUILD.selenium_java",
        url="http://selenium-release.storage.googleapis.com/3.0-beta2/selenium-java-3.0.0-beta2.zip",
        sha256="3ee5d714c18e7bbbd3c112961712a825da057854c8f5f7ca12af368ac3270b29",
    )

    native.maven_jar(
        name="org_json_json",
        artifact="org.json:json:20160810",
        sha1="aca5eb39e2a12fddd6c472b240afe9ebea3a6733",)

    native.maven_jar(
        name="com_google_code_findbugs_jsr305",
        artifact="com.google.code.findbugs:jsr305:3.0.1",
        sha1="f7be08ec23c21485b9b5a1cf1654c2ec8c58168d",)

    native.maven_jar(
        name="com_google_guava_guava",
        artifact="com.google.guava:guava:19.0",
        sha1="6ce200f6b23222af3d8abb6b6459e6c44f4bb0e9",)

  if go:
    native.new_git_repository(
        name="com_github_tebeka_selenium",
        build_file=prefix + "//:BUILD.selenium_go",
        remote="https://github.com/tebeka/selenium.git",
        tag="v0.8.5",)
