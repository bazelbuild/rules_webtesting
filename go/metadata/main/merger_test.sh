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

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/main/main --output $TEST_TMPDIR/out.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/chrome-linux.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/android-browser-gingerbread-nexus-s.json

diff -b $TEST_TMPDIR/out.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/merger-result.json
if [[ $? != 0 ]]; then
	echo "Merge of chrome-linux.json with android-browser-gingerbread-nexus-s.json didn't equal merger-result.json."
	error=1
fi

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/main/main --output $TEST_TMPDIR/out2.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files1.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files1.json
if [[ $? != 0 ]]; then
	echo "Merge of named-files1.json with itself failed."
	error=1
fi

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/main/main --output $TEST_TMPDIR/out2.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files1.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/named-files2.json
if [[ $? == 0 ]]; then
	echo "Expected merge of named-files1.json with named-files2.json to fail."
	error=1
fi

$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/main/main --output $TEST_TMPDIR/out2.json \
	$TEST_SRCDIR/$TEST_WORKSPACE/go/metadata/testdata/bad-named-files.json
if [[ $? == 0 ]]; then
	echo "Expected load of bad-named-files.json to fail."
	error=1
fi

exit $error
