package httphelper

import (
  "os"
  "strings"
  "testing"
)

func TestFQDN(t *testing.T) {
  fqdn, err := FQDN()

  if err != nil {
    t.Error(err)
  }

  name, _ := os.Hostname()

  if !strings.HasPrefix(fqdn, name) {
    t.Errorf("Got %q, expected to start with %q", fqdn, name)
  }
}
