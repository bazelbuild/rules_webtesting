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


def web_test_repositories(java=False,
                          go=False,
                          python=False,
                          omit_com_github_gorilla_mux=False,
                          omit_org_seleniumhq_server=False,
                          omit_org_seleniumhq_java=False,
                          omit_org_json_json=False,
                          omit_com_google_code_findbugs_jsr305=False,
                          omit_com_google_guava_guava=False,
                          omit_com_github_tebeka_selenium=False,
                          omit_org_seleniumhq_py=False):
  if not omit_com_github_gorilla_mux:
    native.new_http_archive(
        name="com_github_gorilla_mux",
        build_file=str(Label("//build_files:BUILD.gorilla_mux")),
        url="https://github.com/gorilla/mux/archive/cf79e51a62d8219d52060dfc1b4e810414ba2d15.tar.gz",
        sha256="80077e14b2f0f8f2796b6bfcf5c8e41e148e3c8c45b4c20d1e6856b348d5efb7",
        strip_prefix="mux-cf79e51a62d8219d52060dfc1b4e810414ba2d15")

  if not omit_org_seleniumhq_server:
    native.http_jar(
        name="org_seleniumhq_server",
        sha256="f5ada04a651ba7ec70fcbc68bd4a59342a928ef7dce858ec594a8d5c49576ace",
        url="http://selenium-release.storage.googleapis.com/3.0-beta3/selenium-server-standalone-3.0.0-beta3.jar"
    )

  if java:
    if not omit_org_seleniumhq_java:
      native.new_http_archive(
          name="org_seleniumhq_java",
          build_file=str(Label("//build_files:BUILD.selenium_java")),
          sha256="a26a449388abd46d1e152771e3641859ac4acee9c0ea24a101ca369048a81ecb",
          url="http://selenium-release.storage.googleapis.com/3.0-beta3/selenium-java-3.0.0-beta3.zip"
      )

    if not omit_org_json_json:
      native.maven_jar(
          name="org_json_json",
          artifact="org.json:json:20160810",
          sha1="aca5eb39e2a12fddd6c472b240afe9ebea3a6733")

    if not omit_com_google_code_findbugs_jsr305:
      native.maven_jar(
          name="com_google_code_findbugs_jsr305",
          artifact="com.google.code.findbugs:jsr305:3.0.1",
          sha1="f7be08ec23c21485b9b5a1cf1654c2ec8c58168d")

    if not omit_com_google_guava_guava:
      native.maven_jar(
          name="com_google_guava_guava",
          artifact="com.google.guava:guava:19.0",
          sha1="6ce200f6b23222af3d8abb6b6459e6c44f4bb0e9")

  if go:
    if not omit_com_github_tebeka_selenium:
      native.new_http_archive(
          name="com_github_tebeka_selenium",
          build_file=str(Label("//build_files:BUILD.selenium_go")),
          url="https://github.com/tebeka/selenium/archive/v0.9.2.tar.gz",
          sha256="c5f21652eda6230ee8bb5f9f02b740fa8d8b22c0cddc832ec666a7654bb0d1a4",
          strip_prefix="selenium-0.9.2")

  if python:
    if not omit_org_seleniumhq_py:
      native.new_http_archive(
          name="org_seleniumhq_py",
          build_file=str(Label("//build_files:BUILD.selenium_py")),
          sha256="0705803349964c7a2a144f1796a5d29905fe2a09931b2bb945ee0cb4deab75d7",
          strip_prefix="selenium-3.0.1/py",
          url="https://pypi.python.org/packages/3a/a3/e4ab60a0229a85f468a36367bc0672a4bca2720f24391eda33704a5f0ad5/selenium-3.0.1.tar.gz"
      )
