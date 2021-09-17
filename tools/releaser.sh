#!/bin/bash

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

# Shell script for creating a release directory.

copybaraBin=$(which copybara)

# If Copybara is not available in the `PATH`, we try using Copybara
# from a path where it is supposed to exist (Googlers-only).
if [[ -z "${copybaraBin}" ]]; then
  copybaraBin=/google/data/ro/teams/copybara/copybara
fi

cd "$(dirname $0)/.."
pwd

bazel build -c opt --stamp \
    --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 \
    //go/metadata/main \
    //go/wsl/main \
    //go/wtl/main

bazel build -c opt --stamp \
    --platforms=@io_bazel_rules_go//go/toolchain:darwin_arm64 \
    //go/metadata/main \
    //go/wsl/main \
    //go/wtl/main

bazel build -c opt --stamp \
    --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 \
    //go/metadata/main \
    //go/wsl/main \
    //go/wtl/main

bazel build -c opt --stamp \
    --platforms=@io_bazel_rules_go//go/toolchain:windows_amd64 \
    //go/metadata/main \
    //go/wsl/main \
    //go/wtl/main

tmpDir="$PWD/dist"
outputDir="${tmpDir}/release-artifact"
archiveDir="${tmpDir}/rules_webtesting.tar.gz"

# Create a temporary directory for storing the release artifacts.
rm -Rf ${tmpDir}
mkdir -p ${tmpDir}

# Build the release artifact directory using the Copybara `release` workflow.
${copybaraBin} --folder-dir="${outputDir}" migrate tools/copy.bara.sky "release" $PWD

# Create the release output tarball.
tar -cvzf ${archiveDir} -C ${outputDir} .

echo "Release tarball has been stored in: ${archiveDir}"
echo "Attach this tarball to the Github release entry."
