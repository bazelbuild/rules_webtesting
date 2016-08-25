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

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")


def web_test_bindings(launcher="@io_bazel_rules_web//launcher:main",
                      merger="@io_bazel_rules_web//metadata:merger",
                      default_config="@io_bazel_rules_web//rules:default"):
  native.bind(
      name="web_test_launcher",
      actual=launcher,)

  native.bind(
      name="web_test_merger",
      actual=merger,)

  native.bind(
      name="web_test_default_config",
      actual=default_config,)

  go_repositories()
