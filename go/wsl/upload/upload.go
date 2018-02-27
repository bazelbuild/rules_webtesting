// Copyright 2018 Google Inc.
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

// Package files handles file uploads and downloads.
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
	"path/filepath"
	"strings"
)

type Uploader struct {
	Root string
}

func (u *Uploader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet || r.Method == http.MethodDelete {
		errorResponse(w, http.StatusMethodNotAllowed, 9, "unknown method", "file deletion and fetching is not allowed")
		return
	}

	reqJSON := struct {
		File *string `json:"file"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&reqJSON); err != nil {
		errorResponse(w, http.StatusBadRequest, 13, "invalid argument", err.Error())
		return
	}

	if reqJSON.File == nil {
		errorResponse(w, http.StatusBadRequest, 13, "invalid argument", "file field is required in request")
		return
	}

	buffer := &bytes.Buffer{}

	if _, err := io.Copy(buffer, base64.NewDecoder(base64.StdEncoding, strings.NewReader(*reqJSON.File))); err != nil {
		errorResponse(w, http.StatusBadRequest, 13, "invalid argument", err.Error())
		return
	}

	dir, err := ioutil.TempDir(u.Root, "")
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, 13, "unknown error", err.Error())
		return
	}

	files, err := unzipFiles(dir, buffer.Bytes())
	if err != nil {
		f, err := uploadFile(dir, buffer.Bytes())
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, 13, "unknown error", err.Error())
			return
		}
		files = append(files, f)
	}

	respJSON := map[string]interface{}{}

	if len(files) == 1 {
		respJSON["value"] = files[0]
	} else {
		respJSON["value"] = files
	}

	if err := json.NewEncoder(w).Encode(respJSON); err != nil {
		errorResponse(w, http.StatusInternalServerError, 13, "unknown error", err.Error())
	}
}

func unzipFiles(dir string, contents []byte) ([]string, error) {
	unzipped, err := zip.NewReader(bytes.NewReader(contents), int64(len(contents)))
	if err != nil {
		return nil, err
	}

	var files []string

	for _, file := range unzipped.File {
		path, err := unzipFile(dir, file)
		if err != nil {
			return nil, err
		}
		files = append(files, path)
	}

	return files, nil
}

func unzipFile(dir string, file *zip.File) (string, error) {
	path := filepath.Join(dir, file.FileHeader.Name)
	in, err := file.Open()
	if err != nil {
		return "", err
	}
	defer in.Close()

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return "", err
	}

	return path, nil
}

func uploadFile(dir string, contents []byte) (string, error) {
	out, err := ioutil.TempFile(dir, "")
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := out.Write(contents); err != nil {
		return "", err
	}

	return out.Name(), nil
}

func errorResponse(w http.ResponseWriter, httpStatus, status int, err, message string) {
	w.WriteHeader(httpStatus)

	respJSON := map[string]interface{}{
		"status": status,
		"value": map[string]interface{}{
			"error":   err,
			"message": message,
		},
	}

	json.NewEncoder(w).Encode(respJSON)
}
