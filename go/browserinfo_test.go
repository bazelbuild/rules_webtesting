// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////
//
package webtest

import "testing"

func TestBrowserInfo(t *testing.T) {
	i, err := newInfo("io_bazel_rules_webtesting/go/metadata/testdata/all-fields.json")
	if err != nil {
		t.Fatal(err)
	}
	if i.BrowserLabel != "//browsers:figaro" {
		t.Errorf(`got BrowserLabel = %q, expected "//browsers:figaro"`, i.BrowserLabel)
	}
	if i.Environment != "chromeos" {
		t.Errorf(`got Environment = %q, expected "chromeos"`, i.Environment)
	}
}
