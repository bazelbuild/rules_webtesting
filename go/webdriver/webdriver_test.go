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
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
	"github.com/bazelbuild/rules_webtesting/go/webtest"
)

var testHostURL *url.URL

func TestMain(m *testing.M) {
	port, err := portpicker.PickUnusedPort()
	if err != nil {
		log.Fatal(err)
	}

	// Note: We resolve to a testdata file first and then retrieve compute the directory.
	// This is necessary in order to support running this test on Windows. The runfile manifest
	// only contains mappings for files, so there is no mapping for just the `testdata/` directory.
	resolved_file, err := bazel.Runfile("rules_webtesting/testdata/BUILD.bazel")
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(resolved_file)

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir))))
	}()

	testHostURL, _ = url.Parse(fmt.Sprintf("http://localhost:%d", port))

	os.Exit(m.Run())
}

func TestCreateSessionAndQuit(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	if d.SessionID() == "" {
		t.Error("session ID should be set")
	}
	if name, _ := d.Capabilities()["browserName"].(string); name == "" {
		t.Error("capabilities browserName should be non-empty")
	}
	if sid, ok := d.Capabilities()["sessionId"]; ok {
		t.Errorf("capabilities should not contain session ID; has sessionId key with value %q", sid)
	}
	if err := d.Quit(ctx); err != nil {
		t.Error(err)
	}
}

func TestHealthy(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 1, nil)
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

	d, err := CreateSession(ctx, wdAddress(), 1, nil)
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

	d, err := CreateSession(ctx, wdAddress(), 1, nil)
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

func TestCurrentURL(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	u, err := d.CurrentURL(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if u == nil {
		t.Fatal("got nil, expected a url.URL ")
	}
}

func TestNavigateTo(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	u, err := url.Parse("https://www.google.com")

	if err := d.NavigateTo(ctx, u); err != nil {
		t.Fatal(err)
	}

	cu, err := d.CurrentURL(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(cu.Hostname(), "google.com") {
		t.Fatalf("got %v, expected to contain google.com", cu)
	}
}

func TestScreenshot(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
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

func TestActiveElement(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	testURL, _ := testURL("webdriver_autofocus.html")
	if err := d.NavigateTo(ctx, testURL); err != nil {
		t.Fatal(err)
	}

	e, err := d.ActiveElement(ctx)
	if err != nil {
		t.Fatal(err)
	}

	id, err := d.ElementGetAttribute(ctx, e, "id")
	if err != nil {
		t.Fatal(err)
	}

	if id != "input1" {
		t.Fatalf("Expected 'input1' but got '%v'", id)
	}
}

func TestKeyDownUp(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	testURL, _ := testURL("webdriver_autofocus.html")
	if err := d.NavigateTo(ctx, testURL); err != nil {
		t.Fatal(err)
	}

	d.KeyDown(ctx, TabKey)
	d.KeyUp(ctx, TabKey)

	e, err := d.ActiveElement(ctx)
	if err != nil {
		t.Fatal(err)
	}

	id, err := d.ElementGetAttribute(ctx, e, "id")
	if err != nil {
		t.Fatal(err)
	}

	if id != "input2" {
		t.Fatalf("Expected 'input2' but got '%v'", id)
	}
}

func TestWindowHandles(t *testing.T) {
	ctx := context.Background()

	driver, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Quit(ctx)

	if windows, err := driver.WindowHandles(ctx); err != nil {
		t.Fatal(err)
	} else if len(windows) != 1 {
		t.Fatalf("Got %v handles, expected 1", len(windows))
	}
}

func TestCurrentWindowHandle(t *testing.T) {
	ctx := context.Background()

	driver, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Quit(ctx)

	if handle, err := driver.CurrentWindowHandle(ctx); err != nil {
		t.Fatal(err)
	} else if handle == "" {
		t.Fatal(`Got "" handle, expected non-empty string`)
	}
}

func TestQuit(t *testing.T) {
	ctx := context.Background()

	driver, err := CreateSession(ctx, wdAddress(), 3, nil)
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

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
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
		t.Errorf("Got runtime %v, expected < 1 and < 5 seconds", run)
	}

	start = time.Now()
	if err := d.ExecuteScriptAsync(ctx, "return;", nil, nil); err == nil {
		t.Error("Got nil err, expected timeout err")
	}
	if run := time.Now().Sub(start); run < 5*time.Second {
		t.Errorf("Got runtime %v, expected > 5 seconds", run)
	}
}

func TestExecuteScriptAsyncWithTimeoutWithCaps(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, &capabilities.Capabilities{
		AlwaysMatch: map[string]interface{}{
			"timeouts": map[string]interface{}{
				"script": 5000,
			},
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
		t.Errorf("Got runtime %v, expected < 1 and < 5 seconds", run)
	}

	start = time.Now()
	if err := d.ExecuteScriptAsync(ctx, "return;", nil, nil); err == nil {
		t.Error("Got nil err, expected timeout err")
	}
	if run := time.Now().Sub(start); run < 5*time.Second {
		t.Errorf("Got runtime %v, expected > 5 seconds", run)
	}
}

func TestGetWindowRect(t *testing.T) {
	if bi, _ := webtest.GetBrowserInfo(); bi.Environment == "sauce" {
		t.Skip("fails on SauceLabs.")
	}

	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	rect, err := d.GetWindowRect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if rect.X < 0 {
		t.Errorf("got rect.X == %v, expected >= 0", rect.X)
	}

	if rect.Y < 0 {
		t.Errorf("got rect.Y == %v, expected >= 0", rect.Y)
	}

	if rect.Width <= 0 {
		t.Errorf("got rect.Width == %v, expected > 0", rect.Width)
	}

	if rect.Height <= 0 {
		t.Errorf("got rect.Height == %v, expected > 0", rect.Height)
	}
}

func TestSetWindowRect(t *testing.T) {
	testCases := []struct {
		name  string
		rect  Rectangle
		check bool
		err   bool
	}{
		{
			"valid",
			Rectangle{
				X:      200,
				Y:      200,
				Width:  600,
				Height: 400,
			},
			true,
			false,
		},
		{
			"zeroes",
			Rectangle{},
			false,
			false,
		},
		{
			"negative location",
			Rectangle{
				X:      -200,
				Y:      -200,
				Width:  500,
				Height: 400,
			},
			false, // what happens is os/wm dependent.
			false,
		},
	}

	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := d.SetWindowRect(ctx, tc.rect)
			if tc.err {
				if err == nil {
					t.Fatal("got nil err, expected non-nil err")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !tc.check {
				return
			}

			rect, err := d.GetWindowRect(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if rect != tc.rect {
				t.Errorf("got rect == %+v, expected %+v", rect, tc.rect)
			}
		})
	}
}

func TestSetWindowSize(t *testing.T) {
	if bi, _ := webtest.GetBrowserInfo(); bi.Environment == "sauce" {
		t.Skip("fails on SauceLabs.")
	}

	testCases := []struct {
		name   string
		width  float64
		height float64
		check  bool
		err    bool
	}{
		{
			"valid",
			600,
			400,
			true,
			false,
		},
		{
			"zeroes",
			0,
			0,
			false,
			false,
		},
	}

	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := d.SetWindowSize(ctx, tc.width, tc.height)
			if tc.err {
				if err == nil {
					t.Fatal("got nil err, expected non-nil err")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !tc.check {
				return
			}
			rect, err := d.GetWindowRect(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if tc.width != rect.Width || tc.height != rect.Height {
				t.Errorf("got (w, h) == (%v, %v), expected (%v, %v)", rect.Width, rect.Height, tc.width, tc.height)
			}
		})
	}
}

func TestSetWindowPosition(t *testing.T) {
	if bi, _ := webtest.GetBrowserInfo(); bi.Environment == "sauce" {
		t.Skip("fails on SauceLabs.")
	}

	testCases := []struct {
		name  string
		x     float64
		y     float64
		check bool
		err   bool
	}{
		{
			"valid",
			200,
			200,
			true,
			false,
		},
		{
			"zeroes",
			0,
			0,
			false,
			false,
		},
		{
			"negative",
			-200,
			-200,
			false, // what happens is os/wm dependent.
			false,
		},
	}

	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := d.SetWindowPosition(ctx, tc.x, tc.y)
			if tc.err {
				if err == nil {
					t.Fatal("got nil err, expected non-nil err")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !tc.check {
				return
			}

			rect, err := d.GetWindowRect(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if rect.X != tc.x || rect.Y != tc.y {
				t.Errorf("got rect == %+v, expected X: %v, Y: %v", rect, tc.x, tc.y)
			}
		})
	}
}

func TestSwitchToFrame(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	testURL, _ := testURL("webdriver.html")
	if err := d.NavigateTo(ctx, testURL); err != nil {
		t.Fatal(err)
	}

	// Initial context
	assertInFrame(ctx, t, d, false)

	// Switch to default content from initial context
	if err := d.SwitchToFrame(ctx, nil); err != nil {
		t.Fatal(err)
	}
	assertInFrame(ctx, t, d, false)

	// Switch to frame by index from initial context
	if err := d.SwitchToFrame(ctx, 0); err != nil {
		t.Fatal(err)
	}
	assertInFrame(ctx, t, d, true)

	// Switch to default content from frame
	if err := d.SwitchToFrame(ctx, nil); err != nil {
		t.Fatal(err)
	}
	assertInFrame(ctx, t, d, false)

	// Switch to frame by index (again) from initial context
	if err := d.SwitchToFrame(ctx, 0); err != nil {
		t.Fatal(err)
	}
	assertInFrame(ctx, t, d, true)

	// Switch to default content which is the parent of the current frame
	if err := d.SwitchToParentFrame(ctx); err != nil {
		t.Fatal(err)
	}
	assertInFrame(ctx, t, d, false)
}

func assertInFrame(ctx context.Context, t *testing.T, d WebDriver, want bool) {
	top := false
	if err := d.ExecuteScript(ctx, "return window.self === window.top;", nil, &top); err != nil {
		t.Fatal(err)
	}
	if top == want {
		t.Errorf("inFrame = %t, but want %t", !top, want)
	}
}

func TestSwitchToWindow(t *testing.T) {
	ctx := context.Background()

	d, err := CreateSession(ctx, wdAddress(), 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Quit(ctx)

	testURL, _ := testURL("webdriver.html")
	if err := d.NavigateTo(ctx, testURL); err != nil {
		t.Fatal(err)
	}

	handle, err := d.CurrentWindowHandle(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Open a new window.
	if err := d.ExecuteScript(ctx, "window.open();", nil, nil); err != nil {
		t.Fatal(err)
	}

	// Get the handle of the new window.
	handles, err := d.WindowHandles(ctx)
	if err != nil {
		t.Fatal(err)
	}
	newHandle := ""
	for _, h := range handles {
		if h != handle {
			newHandle = h
			break
		}
	}
	if newHandle == "" {
		t.Fatal("got no new handle, but want new handle")
	}

	// Switch the context to the new window.
	if err := d.SwitchToWindow(ctx, newHandle); err != nil {
		t.Fatal(err)
	}

	// Verify the expected context switch has occurred.
	handle, err = d.CurrentWindowHandle(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if handle != newHandle {
		t.Errorf("driver.CurrentWindowHandle = %s, but want %s", handle, newHandle)
	}
}

func wdAddress() string {
	addr := os.Getenv("WEB_TEST_WEBDRIVER_SERVER")
	if !strings.HasSuffix(addr, "/") {
		addr = addr + "/"
	}
	return addr
}

func testURL(filename string) (*url.URL, error) {
	testURL, err := url.Parse(fmt.Sprintf("%s/%s", testHostURL, filename))
	if err != nil {
		return nil, err
	}

	return testURL, nil
}
