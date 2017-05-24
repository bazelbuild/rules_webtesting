// Copyright 2017 Google Inc.
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

package scripttimeout

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "testing"
  "time"

  "github.com/bazelbuild/rules_webtesting/go/bazel"
  "github.com/bazelbuild/rules_webtesting/go/portpicker"
  "github.com/bazelbuild/rules_webtesting/go/webtest"
  "github.com/tebeka/selenium"
)

var testpage = ""

func TestMain(m *testing.M) {
  port, err := portpicker.PickUnusedPort()
  if err != nil {
    log.Fatal(err)
  }

  dir, err := bazel.Runfile("testdata/")
  if err != nil {
    log.Fatal(err)
  }

  go func() {
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir))))
  }()

  testpage = fmt.Sprintf("http://localhost:%d/testpage.html", port)

  os.Exit(m.Run())
}

func TestSetScriptTimeout(t *testing.T) {
  // This test only superficially tests script timeout functionality (e.g. that the timeout still gets set).
  driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
  if err != nil {
    t.Fatal(err)
  }

  defer driver.Quit()

  if err := driver.Get(testpage); err != nil {
    t.Fatal(err)
  }

  if err := driver.SetAsyncScriptTimeout(1 * time.Second); err != nil {
    t.Fatal(err)
  }

  start := time.Now()
  if _, err := driver.ExecuteScriptAsync("return;", []interface{}{}); err == nil {
    t.Fatal("got nil err, expected timeout err")
  }
  if run := time.Now().Sub(start); run < 1*time.Second {
    t.Fatalf("got runtime %v, expected to be at least 1 seconds", run)
  }

  if err := driver.SetAsyncScriptTimeout(5 * time.Second); err != nil {
    t.Fatal(err)
  }

  start = time.Now()
  if _, err := driver.ExecuteScriptAsync("return;", []interface{}{}); err == nil {
    t.Fatal("got nil err, expected timeout err")
  }
  if run := time.Now().Sub(start); run < 5*time.Second {
    t.Fatalf("got runtime %v, expected to be at least 1 seconds", run)
  }

}
