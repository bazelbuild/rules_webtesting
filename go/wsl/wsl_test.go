package wsl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
)

func TestHandleHealthz(t *testing.T) {
	handler := createHandler(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}), "", func() {})

	w := newFakeResponseWriter()
	r, err := http.NewRequest(http.MethodGet, "http://localhost/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(w, r)

	if w.status != http.StatusOK {
		t.Errorf(`Got status %d, want %d`, w.status, http.StatusOK)
	}

	if !strings.Contains(string(w.Bytes()), "ok") {
		t.Errorf(`Got %q, want to contain "ok"`, string(w.Bytes()))
	}
}

func TestHandleQuitQuitQuit(t *testing.T) {
	cancelCalled := 0

	cancel := func() {
		cancelCalled = cancelCalled + 1
	}

	handler := createHandler(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}), "", cancel)

	w := newFakeResponseWriter()
	r, err := http.NewRequest(http.MethodGet, "http://localhost/quitquitquit", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(w, r)

	if w.status != http.StatusOK {
		t.Errorf(`Got status %d, want %d`, w.status, http.StatusOK)
	}

	if !strings.Contains(string(w.Bytes()), "shutting down") {
		t.Errorf(`Got %q, want to contain "shutting down"`, string(w.Bytes()))
	}

	if cancelCalled != 1 {
		t.Errorf("Cancel was called %d times, want 1 time", cancelCalled)
	}
}

func TestHandleSession(t *testing.T) {
	hubCalled := 0

	hub := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		hubCalled = hubCalled + 1
	})

	handler := createHandler(hub, "", func() {})

	r, err := http.NewRequest(http.MethodGet, "http://localhost/session", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(newFakeResponseWriter(), r)

	if hubCalled != 1 {
		t.Errorf("Hub was called %d times, want 1 time", hubCalled)
	}

	r, err = http.NewRequest(http.MethodGet, "http://localhost/session/id", nil)
	handler.ServeHTTP(newFakeResponseWriter(), r)

	if hubCalled != 2 {
		t.Errorf("Hub was called %d times, want 2 time", hubCalled)
	}
}

func TestHandleGoogleStaticFile(t *testing.T) {
	testData, err := bazel.Runfile("testdata")
	if err != nil {
		t.Fatal(err)
	}

	handler := createHandler(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}), testData, func() {})

	t.Run("exists", func(t *testing.T) {
		w := newFakeResponseWriter()
		r, _ := http.NewRequest(http.MethodGet, "http://localhost/google/staticfile/testpage.html", nil)

		handler.ServeHTTP(w, r)

		if w.status != http.StatusOK {
			t.Errorf(`Got status %d, want %d`, w.status, http.StatusOK)
		}

		if !strings.Contains(string(w.Bytes()), "Test Page") {
			t.Errorf(`Got %q, want to contain "Test Page"`, string(w.Bytes()))
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		w := newFakeResponseWriter()
		r, _ := http.NewRequest(http.MethodGet, "http://localhost/google/staticfile/does-not-exist.txt", nil)

		handler.ServeHTTP(w, r)

		if w.status != http.StatusNotFound {
			t.Errorf(`Got status %d, want %d`, w.status, http.StatusNotFound)
		}
	})
}

func TestHandleStatus(t *testing.T) {
	handler := createHandler(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}), "", func() {})

	r, err := http.NewRequest(http.MethodGet, "http://localhost/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	frw := newFakeResponseWriter()

	handler.ServeHTTP(frw, r)

	if frw.status != http.StatusOK {
		t.Errorf("Got status %d, want %d", frw.status, http.StatusOK)
	}

	value := struct{
		Status *int
		Value *struct {
			Build map[string]string
			OS map[string]string
			Ready *bool
			Message *string
		}
	}{}

	if err := json.NewDecoder(frw).Decode(&value); err != nil {
		t.Fatal(err)
	}

	if value.Status == nil || *value.Status != 0{
		t.Errorf("Got status %v, want 0", value.Status)
	} 

	if value.Value == nil {
		t.Fatal("Got nil value, want non-nil")
	}

	if value.Value.Ready == nil || !*value.Value.Ready {
		t.Errorf("Got %v Value.Ready, want true", value.Value.Ready)
	}

	if value.Value.Message == nil  {
		t.Error("Got nil Value.Message, want non-nil")
	}

	if value.Value.Build == nil  {
		t.Error("Got nil Value.Build, want non-nil")
	}

	if value.Value.OS == nil  {
		t.Error("Got nil Value.OS, want non-nil")
	}
}

func TestHandleRoot(t *testing.T) {
	handler := createHandler(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}), "", func() {})

	w := newFakeResponseWriter()
	r, err := http.NewRequest(http.MethodGet, "http://localhost/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(w, r)

	if w.status != http.StatusNotFound {
		t.Errorf(`Got status %d, want %d`, w.status, http.StatusNotFound)
	}
}

func TestStartAndShutdownServer(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "ok")
	})

	port, err := portpicker.PickUnusedPort()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	errChan := make(chan error)

	go func() {
		errChan <- startServer(ctx, port, handler)
	}()

	// Give server plenty of time to get started
	time.Sleep(100 * time.Millisecond)

	addr := fmt.Sprintf("http://localhost:%d", port)

	resp, err := http.Get(addr)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Got %d status, expected 200", resp.StatusCode)
	}

	cancel()

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeoutCancel()

	select {
	case <-timeoutCtx.Done():
		t.Fatal("expected server to stop with 5 seconds")
	case err := <-errChan:
		if err != nil {
			t.Error(err)
		}
	}

	_, err = http.Get(addr)
	if err == nil {
		t.Fatal("Got nil err, want err")
	}
}

func newFakeResponseWriter() *fakeResponseWriter {
	return &fakeResponseWriter{
		Buffer: &bytes.Buffer{},
		status: http.StatusOK,
		header: http.Header(map[string][]string{}),
	}
}

type fakeResponseWriter struct {
	*bytes.Buffer
	status int
	header http.Header
}

func (frw *fakeResponseWriter) Header() http.Header {
	return frw.header
}

func (frw *fakeResponseWriter) WriteHeader(status int) {
	frw.status = status
}
