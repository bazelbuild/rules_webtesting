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
#
################################################################################
#
package(
    default_testonly = True,
    default_visibility = ["//:__subpackages__"],
)

licenses(["notice"])  # Apache 2.0

filegroup(
    name = "testdata",
    srcs = glob(["*"]),
)

load("//web:web.bzl", "web_test_named_file", "web_test_config")

web_test_named_file(
    name = "PROXY_TLS_KEY",
    file = "key.pem",
)

web_test_named_file(
    name = "PROXY_TLS_CERT",
    file = "cert.pem",
)

web_test_config(
    name = "https",
    deps = [
        ":PROXY_TLS_CERT",
        ":PROXY_TLS_KEY",
        "//web:default_config",
    ],
)

web_test_config(
    name = "mobileemulation",
    metadata = "mobileemulation.json",
    deps = ["//web:default_config"],
)
