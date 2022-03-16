#!/bin/bash -eu
#
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

# --- begin runfiles.bash initialization v2 ---
# Copy-pasted from the Bazel Bash runfiles library v2. We need to copy the runfile
# helper code as we want to resolve Bazel targets through the runfiles (in order to
# make the test work on windows where runfiles are not symlinked). The runfile
# helpers expose a bash function called "rlocation" that can be used to resolve targets.
# https://github.com/bazelbuild/bazel/blob/master/tools/bash/runfiles/runfiles.bash.
set -uo pipefail; f=bazel_tools/tools/bash/runfiles/runfiles.bash
source "${RUNFILES_DIR:-/dev/null}/$f" 2>/dev/null || \
source "$(grep -sm1 "^$f " "${RUNFILES_MANIFEST_FILE:-/dev/null}" | cut -f2- -d' ')" 2>/dev/null || \
source "$0.runfiles/$f" 2>/dev/null || \
source "$(grep -sm1 "^$f " "$0.runfiles_manifest" | cut -f2- -d' ')" 2>/dev/null || \
source "$(grep -sm1 "^$f " "$0.exe.runfiles_manifest" | cut -f2- -d' ')" 2>/dev/null || \
{ echo>&2 "ERROR: cannot find $f"; exit 1; }; f=; set -e
# --- end runfiles.bash initialization v2 ---

set +e

error=0

# The path to the merger is passed as `sh_test` attribute. This
# allows us to access the merger in a platform-agnostic way.
merger_root_path=${1}
merger=$(rlocation $TEST_WORKSPACE/${merger_root_path})

$merger --output $TEST_TMPDIR/out.json \
    $(rlocation $TEST_WORKSPACE/testdata/chrome-linux.json) \
    $(rlocation $TEST_WORKSPACE/testdata/android-browser-gingerbread-nexus-s.json)

diff -b $TEST_TMPDIR/out.json \
    $(rlocation $TEST_WORKSPACE/testdata/merger-result.json)
if [[ $? != 0 ]]; then
  echo "Merge of chrome-linux.json with android-browser-gingerbread-nexus-s.json didn't equal merger-result.json."
  error=1
fi

$merger --output $TEST_TMPDIR/out2.json \
    $(rlocation $TEST_WORKSPACE/testdata/named-files1.json) \
    $(rlocation $TEST_WORKSPACE/testdata/named-files1.json)
if [[ $? != 0 ]]; then
  echo "Merge of named-files1.json with itself failed."
  error=1
fi

$merger --output $TEST_TMPDIR/out2.json \
    $(rlocation $TEST_WORKSPACE/testdata/named-files1.json) \
    $(rlocation $TEST_WORKSPACE/testdata/named-files2.json)
if [[ $? == 0 ]]; then
  echo "Expected merge of named-files1.json with named-files2.json to fail."
  error=1
fi

$merger --output $TEST_TMPDIR/out2.json \
    $(rlocation $TEST_WORKSPACE/testdata/bad-named-files.json)
if [[ $? == 0 ]]; then
  echo "Expected load of bad-named-files.json to fail."
  error=1
fi

exit $error
