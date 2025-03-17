#!/bin/bash
# Copyright 2019 Google Inc. All Rights Reserved
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS-IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -ex

# Build and test web java and related tests in web_java module
(cd web_java && bazel test ... --java_runtime_version=17)

# Build and test web python and related tests in web_python module
(cd web_python && bazel test ...)

# Build and test web scala and related tests in web_scala module
(cd web_scala && bazel test ...)
