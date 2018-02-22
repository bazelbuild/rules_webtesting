package upload

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
)

func TestWrongMethod(t *testing.T) {
	uploader := &Uploader{}

	request, err := http.NewRequest(http.MethodGet, "http://localhost/", nil)
	if err != nil {
		t.Fatal(err)
	}

	frw := newFakeResponseWriter()

	uploader.ServeHTTP(frw, request)

	if frw.status != http.StatusMethodNotAllowed {
		t.Errorf("Got status  %d, want %d", frw.status, http.StatusMethodNotAllowed)
	}

	respJSON := struct {
		Status int
		Value  struct {
			Error   string
			Message string
		}
	}{}
	if err := json.NewDecoder(frw).Decode(&respJSON); err != nil {
		t.Fatal(err)
	}

	if respJSON.Status == 0 {
		t.Error("Want non-zero status, got 0")
	}

	if respJSON.Value.Error == "" {
		t.Error(`Want non-empty value.error, got ""`)
	}

	if respJSON.Value.Message == "" {
		t.Error(`Want non-empty value.message, got ""`)
	}
}

func TestBadArgs(t *testing.T) {
	uploader := &Uploader{}

	request, err := http.NewRequest(http.MethodPost, "http://localhost/", strings.NewReader("{}"))
	if err != nil {
		t.Fatal(err)
	}

	frw := newFakeResponseWriter()

	uploader.ServeHTTP(frw, request)

	if frw.status != http.StatusBadRequest {
		t.Errorf("Got status  %d, want %d", frw.status, http.StatusBadRequest)
	}

	respJSON := struct {
		Status int
		Value  struct {
			Error   string
			Message string
		}
	}{}
	if err := json.NewDecoder(frw).Decode(&respJSON); err != nil {
		t.Fatal(err)
	}

	if respJSON.Status == 0 {
		t.Error("Want non-zero status, got 0")
	}

	if respJSON.Value.Error == "" {
		t.Error(`Want non-empty value.error, got ""`)
	}

	if respJSON.Value.Message == "" {
		t.Error(`Want non-empty value.message, got ""`)
	}
}

func TestNotZippedUpload(t *testing.T) {
	tmpDir, err := bazel.NewTmpDir("TestNotZippedUpload")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	uploader := &Uploader{Root: tmpDir}

	file := "this is my file contents."

	encodedFile := base64.StdEncoding.EncodeToString([]byte(file))

	reqBytes, err := json.Marshal(map[string]interface{}{
		"file": encodedFile,
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost/", bytes.NewReader(reqBytes))
	if err != nil {
		t.Fatal(err)
	}

	frw := newFakeResponseWriter()

	uploader.ServeHTTP(frw, request)

	if frw.status != http.StatusOK {
		t.Errorf("Got status  %d, want %d", frw.status, http.StatusOK)
	}

	respJSON := struct {
		Status int
		Value  string
	}{}
	if err := json.NewDecoder(frw).Decode(&respJSON); err != nil {
		t.Fatal(err)
	}

	if respJSON.Status != 0 {
		t.Fatalf("Got %d status, want 0", respJSON.Status)
	}

	if respJSON.Value == "" {
		t.Fatal("Got empty file name, want non-empty filename")
	}

	savedFile, err := ioutil.ReadFile(respJSON.Value)

	if string(savedFile) != file {
		t.Errorf("Got %q, want %q", string(savedFile), file)
	}
}

func TestZippedSingleFile(t *testing.T) {
	tmpDir, err := bazel.NewTmpDir("TestNotZippedUpload")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	uploader := &Uploader{Root: tmpDir}

	file := "this is my file contents."

	buffer := bytes.NewBuffer(nil)
	zipWriter := zip.NewWriter(buffer)

	fileWriter, err := zipWriter.Create("file.txt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.WriteString(fileWriter, file); err != nil {
		t.Fatal(err)
	}

	if err := zipWriter.Close(); err != nil {
		t.Fatal(err)
	}

	encodedFile := base64.StdEncoding.EncodeToString(buffer.Bytes())

	reqBytes, err := json.Marshal(map[string]interface{}{
		"file": encodedFile,
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost/", bytes.NewReader(reqBytes))
	if err != nil {
		t.Fatal(err)
	}

	frw := newFakeResponseWriter()

	uploader.ServeHTTP(frw, request)

	if frw.status != http.StatusOK {
		t.Errorf("Got status  %d, want %d", frw.status, http.StatusOK)
	}

	respJSON := struct {
		Status int
		Value  string
	}{}
	if err := json.NewDecoder(frw).Decode(&respJSON); err != nil {
		t.Fatal(err)
	}

	if respJSON.Status != 0 {
		t.Fatalf("Got %d status, want 0", respJSON.Status)
	}

	if respJSON.Value == "" {
		t.Fatal("Got empty file name, want non-empty filename")
	}

	savedFile, err := ioutil.ReadFile(respJSON.Value)

	if string(savedFile) != file {
		t.Errorf("Got %q, want %q", string(savedFile), file)
	}
}

func TestZippedTwoFiles(t *testing.T) {
	tmpDir, err := bazel.NewTmpDir("TestNotZippedUpload")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	uploader := &Uploader{Root: tmpDir}

	file1 := "this is my first file contents."
	file2 := "this is my second file contents."

	buffer := bytes.NewBuffer(nil)
	zipWriter := zip.NewWriter(buffer)

	fileWriter, err := zipWriter.Create("file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.WriteString(fileWriter, file1); err != nil {
		t.Fatal(err)
	}

	fileWriter, err = zipWriter.Create("file2.txt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.WriteString(fileWriter, file2); err != nil {
		t.Fatal(err)
	}

	if err := zipWriter.Close(); err != nil {
		t.Fatal(err)
	}

	encodedFile := base64.StdEncoding.EncodeToString(buffer.Bytes())

	reqBytes, err := json.Marshal(map[string]interface{}{
		"file": encodedFile,
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost/", bytes.NewReader(reqBytes))
	if err != nil {
		t.Fatal(err)
	}

	frw := newFakeResponseWriter()

	uploader.ServeHTTP(frw, request)

	if frw.status != http.StatusOK {
		t.Errorf("Got status  %d, want %d", frw.status, http.StatusOK)
	}

	respJSON := struct {
		Status int
		Value  []string
	}{}
	if err := json.NewDecoder(frw).Decode(&respJSON); err != nil {
		t.Fatal(err)
	}

	if respJSON.Status != 0 {
		t.Fatalf("Got %d status, want 0", respJSON.Status)
	}

	found1 := false
	found2 := false

	for _, file := range respJSON.Value {
		savedFile, err := ioutil.ReadFile(file)
		if err != nil {
			t.Error(err)
			continue
		}

		if strings.HasSuffix(file, "file1.txt") {
			found1 = true
			if string(savedFile) != file1 {
				t.Errorf("Got %q, want %q", string(savedFile), file1)
			}
		} else if strings.HasSuffix(file, "file2.txt") {
			found2 = true
			if string(savedFile) != file2 {
				t.Errorf("Got %q, want %q", string(savedFile), file2)
			}
		} else {
			t.Error("Got extra file %q", file)
		}
	}

	if !found1 {
		t.Error("Missing file1.txt")
	}

	if !found2 {
		t.Error("Missing file2.txt")
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
