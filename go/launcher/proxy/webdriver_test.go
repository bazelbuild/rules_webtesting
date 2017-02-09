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

package webdriver

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestCreateSessionAndQuit(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	if err := d.Quit(ctx); err != nil {
		t.Error(err)
	}
}

func TestCreateSessionFails(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, map[string]interface{}{
		"chromeOptions": map[string]interface{}{
			"binary": "a-binary",
		},
		"moz:firefoxOptions": map[string]interface{}{
			"binary": "a-binary",
		},
		"commandTimeout": "10000000",
	})
	if err == nil {
		t.Error("got nil err from CreateSession with bad capabilities")
		d.Quit(ctx)
	}
}

func TestHealthy(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 1, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	if err := d.Healthy(ctx); err != nil {
		t.Error(err)
	}
	if err := d.Quit(ctx); err != nil {
		t.Error(err)
	}
	if err := d.Healthy(ctx); err == nil {
		t.Error("got nil error from Healthy after quit")
	}
}

func TestExecuteScript(t *testing.T) {
	testCases := []struct {
		script   string
		args     []interface{}
		value    int
		expected int
		err      bool
	}{
		{
			"return 1 + 2;",
			[]interface{}{},
			0,
			3,
			false,
		},
		{
			"return arguments[0] + arguments[1];",
			[]interface{}{1, 2},
			0,
			3,
			false,
		},
		{
			"return argument[0] + arguments[1];",
			[]interface{}{1, 2},
			0,
			0,
			true,
		},
	}

	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 1, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)
	for _, tc := range testCases {
		t.Run(tc.script, func(t *testing.T) {
			if err := d.ExecuteScript(ctx, tc.script, tc.args, &tc.value); err != nil {
				if !tc.err {
					t.Errorf("got unexpected err %v for ExecuteScript(%q, %v)", err, tc.script, tc.args)
				}
				return
			}
			if tc.err {
				t.Errorf("got nil err for ExecuteScript(%q, %v)", tc.script, tc.args)
				return
			}
			if tc.value != tc.expected {
				t.Errorf("got %v, expected %v for ExecuteScript(%q, %v)", tc.value, tc.expected, tc.script, tc.args)
			}
		})
	}
}

func TestExecuteScriptAsync(t *testing.T) {
	testCases := []struct {
		script   string
		args     []interface{}
		value    int
		expected int
		err      bool
	}{
		{
			"arguments[0](1 + 2);",
			[]interface{}{},
			0,
			3,
			false,
		},
		{
			"arguments[2](arguments[0] + arguments[1]);",
			[]interface{}{1, 2},
			0,
			3,
			false,
		},
		{
			"argument[2](argument[0] + argument[1]);",
			[]interface{}{1, 2},
			0,
			0,
			true,
		},
		{
			"return 1 + 2;",
			[]interface{}{1, 2},
			0,
			0,
			true,
		},
	}

	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 1, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	for _, tc := range testCases {
		t.Run(tc.script, func(t *testing.T) {
			if err := d.ExecuteScriptAsync(ctx, tc.script, tc.args, &tc.value); err != nil {
				if !tc.err {
					t.Errorf("got unexpected err %v for ExecuteScriptAsync(%q, %v)", err, tc.script, tc.args)
				}
				return
			}
			if tc.err {
				t.Errorf("got nil err for ExecuteScriptAsync(%q, %v)", tc.script, tc.args)
				return
			}
			if tc.value != tc.expected {
				t.Errorf("got %v, expected %v for ExecuteScriptAsync(%q, %v)", tc.value, tc.expected, tc.script, tc.args)
			}
		})
	}
}

func TestScreenshot(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	img, err := d.Screenshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("got nil, expected an image.Image")
	}
}

func TestWindowHandles(t *testing.T) {
	ctx := context.Background()

	driver, err := CreateSession(ctx, wdAddress(), 3, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Quit(ctx)

	if windows, err := driver.WindowHandles(ctx); err != nil {
		t.Fatal(err)
	} else if len(windows) != 1 {
		t.Fatalf("Got %d handles, expected 1", len(windows))
	}
}

func TestQuit(t *testing.T) {
	ctx := context.Background()

	driver, err := CreateSession(ctx, wdAddress(), 3, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	driver.Quit(ctx)

	if _, err := driver.WindowHandles(ctx); err == nil {
		t.Fatal("Got nil err, expected unknown session err")
	}
}

func wdAddress() string {
	addr := os.Getenv("WEB_TEST_WEBDRIVER_SERVER")
	if !strings.HasSuffix(addr, "/") {
		addr = addr + "/"
	}
	return addr
}
