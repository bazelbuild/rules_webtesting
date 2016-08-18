#!/bin/sh

printenv

$TEST_SRCDIR/__main__/metadata/merger --output $TEST_TMPDIR/out.json \
	$TEST_SRCDIR/__main__/metadata/testdata/chrome-linux.json \
	$TEST_SRCDIR/__main__/metadata/testdata/android-browser-gingerbread-nexus-s.json

diff $TEST_TMPDIR/out.json \
	$TEST_SRCDIR/__main__/metadata/testdata/merger-result.json