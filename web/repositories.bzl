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
    launcher="@io_bazel_rules_web//go/launcher:main",
    merger="@io_bazel_rules_web//go/metadata:merger",
    default_config="@io_bazel_rules_web//web:default_config"):
  native.new_git_repository(
      name="com_github_gorilla_mux",
      build_file=prefix + "//build_files:BUILD.gorilla_mux",
      commit="cf79e51a62d8219d52060dfc1b4e810414ba2d15",
      remote="https://github.com/gorilla/mux.git",)

  native.http_jar(
      name="org_seleniumhq_server",
      sha256="f5ada04a651ba7ec70fcbc68bd4a59342a928ef7dce858ec594a8d5c49576ace",
      url="http://selenium-release.storage.googleapis.com/3.0-beta3/selenium-server-standalone-3.0.0-beta3.jar",
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
        build_file=prefix + "//build_files:BUILD.selenium_java",
        sha256="a26a449388abd46d1e152771e3641859ac4acee9c0ea24a101ca369048a81ecb",
        url="http://selenium-release.storage.googleapis.com/3.0-beta3/selenium-java-3.0.0-beta3.zip",
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
        build_file=prefix + "//build_files:BUILD.selenium_go",
        remote="https://github.com/tebeka/selenium.git",
        tag="v0.8.5",)
