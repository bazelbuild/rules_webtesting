# Copyright 2016 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

workspace(name = "io_bazel_rules_web")

load("//web:repositories.bzl", "web_test_repositories")

web_test_repositories(
    go = True,
    java = True,
    prefix = "",
)

load("//web:bindings.bzl", "web_test_bindings")

web_test_bindings(
    default_config = "//rules:default",
    launcher = "//launcher:main",
    merger = "//metadata:merger",
)

maven_jar(
    name = "junit_junit",
    artifact = "junit:junit:4.12",
    sha1 = "2973d150c0dc1fefe998f834810d68f278ea58ec",
)

maven_jar(
    name = "com_google_truth_truth",
    artifact = "com.google.truth:truth:0.29",
    sha1 = "b6ad12d98295a7d17b3fe4b8969d0f7905626b30",
)
