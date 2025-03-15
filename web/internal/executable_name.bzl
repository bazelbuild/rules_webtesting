# Copyright 2021 Google LLC.
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

def get_platform_executable_name():
    """
    Retrieves the binary name for an executable based on the current platform.

    This is helpful when shipping binaries within the release package as otherwise
    Go binaries for example might have the same name for various platforms. A consistent
    way of naming executables for specific platforms allows them to be shipped in the
    release output, where the binaries can be referenced using the known file names.
    """
    return select({
        "@rules_webtesting//common/conditions:linux_x64": "main_linux_x64",
        "@rules_webtesting//common/conditions:macos_x64": "main_darwin_x64",
        "@rules_webtesting//common/conditions:macos_arm64": "main_darwin_arm64",
        "@rules_webtesting//common/conditions:windows_x64": "main_windows_x64.exe",
    })
