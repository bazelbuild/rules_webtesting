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
#
################################################################################
#
package(default_testonly = True)

licenses(["notice"])  # Apache 2.0

alias(
    name = "main",
    actual  = select({
        "//common/conditions:linux": "linux_amd64_stripped/main",
        "//common/conditions:mac": "darwin_amd64_pure_stripped/main",
        "//common/conditions:windows": "windows_amd64_pure_stripped/main.exe",
    }),
    visibility = ["//visibility:public"],
)
