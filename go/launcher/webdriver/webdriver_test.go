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
	"time"
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

func TestExecuteScriptAsyncWithTimeout(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	if err := d.SetScriptTimeout(ctx, 5*time.Second); err != nil {
		t.Fatal(err)
	}

	start := time.Now()
	if err := d.ExecuteScriptAsyncWithTimeout(ctx, time.Second, "return;", nil, nil); err == nil {
		t.Error("Got nil err, expected timeout err")
	}
	if run := time.Now().Sub(start); run < time.Second || run > 5*time.Second {
		t.Errorf("Got runtime %s, expected < 1 and < 5 seconds", run)
	}

	start = time.Now()
	if err := d.ExecuteScriptAsync(ctx, "return;", nil, nil); err == nil {
		t.Error("Got nil err, expected timeout err")
	}
	if run := time.Now().Sub(start); run < 5*time.Second {
		t.Errorf("Got runtime %s, expected > 5 seconds", run)
	}
}

func TestExecuteScriptAsyncWithTimeoutWithCaps(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, map[string]interface{}{
		"timeouts": map[string]interface{}{
			"script": 5000,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	start := time.Now()
	if err := d.ExecuteScriptAsyncWithTimeout(ctx, time.Second, "return;", nil, nil); err == nil {
		t.Error("Got nil err, expected timeout err")
	}
	if run := time.Now().Sub(start); run < time.Second || run > 5*time.Second {
		t.Errorf("Got runtime %s, expected < 1 and < 5 seconds", run)
	}

	start = time.Now()
	if err := d.ExecuteScriptAsync(ctx, "return;", nil, nil); err == nil {
		t.Error("Got nil err, expected timeout err")
	}
	if run := time.Now().Sub(start); run < 5*time.Second {
		t.Errorf("Got runtime %s, expected > 5 seconds", run)
	}
}

func wdAddress() string {
	addr := os.Getenv("WEB_TEST_WEBDRIVER_SERVER")
	if !strings.HasSuffix(addr, "/") {
		addr = addr + "/"
	}
	return addr
}
