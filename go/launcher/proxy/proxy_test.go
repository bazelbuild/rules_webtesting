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

package proxy

import (
  "context"
  "fmt"
  "net/http"
  "os"
  "testing"

  "github.com/bazelbuild/rules_webtesting/go/httphelper"
)

func TestHTTPSProxy(t *testing.T) {
  address, ok := os.LookupEnv("WEB_TEST_HTTPS_SERVER")
  if !ok {
    t.Fatal("expected environment variable WEB_TEST_HTTPS_SERVER to be defined")
  }

  ctx := context.Background()

  url := fmt.Sprintf("%s/healthz", address)
  resp, err := httphelper.Get(ctx, url)
  if err != nil {
    t.Fatal(err)
  }
  resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
    t.Fatalf("Got status code %d expected %d", resp.StatusCode, http.StatusOK)
  }
}

func TestHTTPProxy(t *testing.T) {
  address, ok := os.LookupEnv("WEB_TEST_HTTP_SERVER")
  if !ok {
    t.Fatal("expected environment variable WEB_TEST_HTTP_SERVER to be defined")
  }

  ctx := context.Background()

  url := fmt.Sprintf("%s/healthz", address)
  resp, err := httphelper.Get(ctx, url)
  if err != nil {
    t.Fatal(err)
  }
  resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
    t.Fatalf("Got status code %d expected %d", resp.StatusCode, http.StatusOK)
  }
}
