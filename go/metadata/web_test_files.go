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

package metadata

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
)

// WebTestFiles defines a set of namedFiles located either in the runfiles directory or
// in an archive file located in the runfiles directory of the test.
type WebTestFiles struct {
	// ArchiveFile is optional path to an archive file (.zip, .tar.gz, .tgz, .tar.bz2, .tbz2, .tar.Z)
	// file. If present, paths in NamedFiles are paths in the archive. If absent, paths in NamedFiles
	// are relative to the runfiles root. The archive will only be extracted if getFilePath is called
	// at least once with a name defined in NamedFiles. If so, the entire archive will be extracted
	// into subdirectory located test tmpdir.
	ArchiveFile string `json:"archiveFile,omitempty"`
	// StripPrefix is an optional prefix that will be stripped when an archive is extracted.
	StripPrefix string `json:"stripPrefix,omitempty"`
	// NamedFiles is a map of names to file paths. These file paths are relative to the runfiles
	// root if ArchiveFile is absent, otherwise they are paths inside the archive referred to by
	// ArchiveFile. The names are used by other parts of Web Test Launcher to refer to needed
	// resources. For example, if your environment needs to know where a chromedriver executable is
	// located, then there could be a name "CHROMEDRIVER" that refers to the path to the chromedriver
	// executable, and the part of you environment that needs to use the chromedriver executable
	// can call md.GetFilePath("CHROMEDRIVER") (where md is a *metadata.Metadata object) which will
	// search through all NamedFiles of all WebTestFiles structs in md to find that key and return
	// the path to the corresponding file (extracting an archive if necessary).
	NamedFiles map[string]string `json:"namedFiles"`

	// The mu field protects access to the extractedPath field.
	mu sync.Mutex
	// The extractedPath field refers to the location where this archive has been extracted to, if
	// has been extracted.
	extractedPath string
}

func normalizeWebTestFiles(in []*WebTestFiles) ([]*WebTestFiles, error) {
	merged := map[string]*WebTestFiles{}
	var archiveFiles []string

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
			archiveFiles = append(archiveFiles, a.ArchiveFile)
			merged[a.ArchiveFile] = a
		}
	}

	names := map[string]bool{}
	var result []*WebTestFiles
	for _, a := range archiveFiles {
		m := merged[a]
		for name := range m.NamedFiles {
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

func (w *WebTestFiles) getFilePath(name string, m *Metadata) (string, error) {
	filename, ok := w.NamedFiles[name]
	if !ok {
		return "", nil
	}

	if w.ArchiveFile == "" {
		return bazel.Runfile(filename)
	}

	if err := w.extract(m); err != nil {
		return "", err
	}

	path := filepath.Join(w.extractedPath, filename)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	return path, nil
}

func (w *WebTestFiles) extract(m *Metadata) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.extractedPath != "" {
		return nil
	}

	extractor, err := m.GetFilePath("EXTRACT_EXE")
	if err != nil {
		return err
	}

	filename, err := bazel.Runfile(w.ArchiveFile)
	if err != nil {
		return err
	}

	extractPath, err := bazel.NewTmpDir(filepath.Base(filename))
	if err != nil {
		return err
	}

	c := exec.Command(extractor, filename, extractPath, w.StripPrefix)

	if err := c.Run(); err != nil {
		return err
	}

	w.extractedPath = extractPath
	return nil
}

func (w *WebTestFiles) String() string {
	return fmt.Sprintf(
		`WebTestFiles{
	ArchiveFile: %q,
	StripPrefix: %q,
	NamedFiles: %+v,
}
`, w.ArchiveFile, w.StripPrefix, w.NamedFiles)

}
