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
set +e
printenv

error=0

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/merger --output $TEST_TMPDIR/out.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/chrome-linux.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/android-browser-gingerbread-nexus-s.json

diff $TEST_TMPDIR/out.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/merger-result.json
if [ $? != 0 ]; then
	echo "Merge result didn't match."
	error=1
fi

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/merger --output $TEST_TMPDIR/out2.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files1.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files1.json
if [ $? != 0 ]; then
	echo "Expected successful named files merge failed."
	error=1
fi

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/merger --output $TEST_TMPDIR/out2.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files1.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files2.json
if [ $? == 0 ]; then
	echo "Expected failing named files merge succeeded."
	error=1
fi

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/merger --output $TEST_TMPDIR/out2.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/bad-named-files.json
if [ $? == 0 ]; then
	echo "Expected failing metdata file load succeeded."
	error=1
fi

exit $error
