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
package metadata

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/util/bazel"
)

// Archive is an archive file and the associated set of named file mappings as
// defined by a web_test_archive rule.
type WebTestFiles struct {
	ArchiveFile string            `json:"ArchiveFile",omitempty`
	NamedFiles  map[string]string `json:"namedFiles"`

	mu            sync.Mutex
	extractedPath string
}

// Only works correctly on normalized WeebTestFiles slices
func webTestFilesSliceEquals(a, b []*WebTestFiles) bool {
	if len(a) != len(b) {
		return false
	}
	first := map[string]*WebTestFiles{}
	for _, f := range a {
		first[f.ArchiveFile] = f
	}
	for _, s := range b {
		f, ok := first[s.ArchiveFile]
		if !ok || !webTestFilesEquals(f, s) {
			return false
		}
	}
	return true
}

func webTestFilesEquals(a, b *WebTestFiles) bool {
	return a.ArchiveFile == b.ArchiveFile && mapEquals(a.NamedFiles, b.NamedFiles)
}

func normalizeWebTestFiles(in []*WebTestFiles) ([]*WebTestFiles, error) {
	merged := map[string]*WebTestFiles{}

	for _, a := range in {
		// skip entries with no named files.
		if len(a.NamedFiles) == 0 {
			continue
		}
		if b := merged[a.ArchiveFile]; b != nil {
			m, err := mergeWebTestFiles(a, b)
			if err != nil {
				return nil, err
			}
			merged[m.ArchiveFile] = m
		} else {
			merged[a.ArchiveFile] = a
		}
	}

	names := map[string]bool{}
	result := []*WebTestFiles{}
	for _, m := range merged {
		for name, _ := range m.NamedFiles {
			if names[name] {
				return nil, fmt.Errorf("name %q exists in multiple WebTestFiles", name)
			}
			names[name] = true
		}
		result = append(result, m)
	}
	return result, nil
}

func mergeWebTestFiles(a1, a2 *WebTestFiles) (*WebTestFiles, error) {
	if a1.ArchiveFile != a2.ArchiveFile {
		return nil, fmt.Errorf("expected paths (%q, %q) to be equal", a1.ArchiveFile, a2.ArchiveFile)
	}
	nf, err := mergeNamedFiles(a1.NamedFiles, a2.NamedFiles)
	if err != nil {
		return nil, err
	}
	return &WebTestFiles{
		ArchiveFile: a1.ArchiveFile,
		NamedFiles:  nf,
	}, nil
}

func mergeNamedFiles(n1, n2 map[string]string) (map[string]string, error) {
	result := map[string]string{}

	for k, v := range n1 {
		result[k] = v
	}

	for k, v2 := range n2 {
		if v1, ok := result[k]; ok && v1 != v2 {
			return nil, fmt.Errorf("key %q exists in both NamedFiles with different values", k)
		}
		result[k] = v2
	}
	return result, nil
}

func (w *WebTestFiles) getFilePath(name string) (string, error) {
	filename, ok := w.NamedFiles[name]
	if !ok {
		return "", nil
	}

	if w.ArchiveFile == "" {
		return bazel.Runfile(filename)
	}

	if err := w.extract(); err != nil {
		return "", err
	}

	path := filepath.Join(w.extractedPath, filename)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	return path, nil
}

func (w *WebTestFiles) extract() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.extractedPath != "" {
		return nil
	}

	filename, err := bazel.Runfile(w.ArchiveFile)
	if err != nil {
		return err
	}

	extractPath, err := bazel.NewTmpDir(filepath.Base(filename))
	if err != nil {
		return err
	}

	var c *exec.Cmd
	switch {
	case strings.HasSuffix(filename, ".tar"):
		c = exec.Command("tar", "xf", filename, "-C", extractPath)
	case strings.HasSuffix(filename, ".tar.gz") || strings.HasSuffix(filename, ".tgz"):
		c = exec.Command("tar", "xzf", filename, "-C", extractPath)
	case strings.HasSuffix(filename, ".tar.bz2") || strings.HasSuffix(filename, ".tbz2"):
		c = exec.Command("tar", "xjf", filename, "-C", extractPath)
	case strings.HasSuffix(filename, ".tar.Z"):
		c = exec.Command("tar", "xZf", filename, "-C", extractPath)
	case strings.HasSuffix(filename, ".zip"):
		c = exec.Command("unzip", filename, "-d", extractPath)
	default:
		return fmt.Errorf("unknown archive type: %s", filename)
	}

	if err := c.Run(); err != nil {
		return err
	}

	w.extractedPath = extractPath
	return nil
}
