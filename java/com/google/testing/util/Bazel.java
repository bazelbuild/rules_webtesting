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
// //////////////////////////////////////////////////////////////////////////////
//
package com.google.testing.util;

import com.google.common.annotations.VisibleForTesting;
import com.google.common.base.Preconditions;
import java.io.IOException;
import java.nio.file.FileSystems;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Optional;

public class Bazel {

  private static Bazel instance;

  private static final String TEST_SRCDIR = "TEST_SRCDIR";
  private static final String TEST_TMPDIR = "TEST_TMPDIR";
  private static final String TEST_WORKSPACE = "TEST_WORKSPACE";
  private static final String DEFAULT_WORKSPACE = "io_bazel_rules_webtesting";

  private final Path runfilesDir;
  private final Optional<Path> testTmpDir;
  private final String testWorkspace;

  @VisibleForTesting
  Bazel(Path runfilesDir, Optional<Path> testTmpDir, String testWorkspace) {
    this.runfilesDir = Preconditions.checkNotNull(runfilesDir);
    this.testTmpDir = Preconditions.checkNotNull(testTmpDir);
    this.testWorkspace = Preconditions.checkNotNull(testWorkspace);
  }

  public static synchronized Bazel getInstance() {
    if (instance == null) {
      String tmpDirPath = System.getenv(TEST_TMPDIR);
      Optional<Path> tmpDir = Optional.empty();
      if (tmpDirPath != null && !tmpDirPath.equals("")) {
        tmpDir = Optional.of(FileSystems.getDefault().getPath(tmpDirPath));
      }
      String testWorkspace = System.getenv(TEST_WORKSPACE);
      if (testWorkspace == null || testWorkspace.equals("")) {
        testWorkspace = DEFAULT_WORKSPACE;
      }
      instance =
          new Bazel(
              FileSystems.getDefault().getPath(System.getenv(TEST_SRCDIR)), tmpDir, testWorkspace);
    }
    return instance;
  }

  // Returns the path of the runfile in the default workspace.
  public Path runfile(String path) throws IOException {
    Path candidate = runfilesDir.resolve(path);
    if (candidate.toFile().exists()) {
      return candidate;
    }
    candidate = runfilesDir.resolve(testWorkspace).resolve(path);
    if (candidate.toFile().exists()) {
      return candidate;
    }
    throw new IOException("Can not find runfile: " + path);
  }

  // Returns a new temporary subdirectory in the test temporary directory.
  public Path newTmpDir(String prefix) throws IOException {
    if (testTmpDir.isPresent()) {
      return Files.createTempDirectory(testTmpDir.get(), prefix);
    }
    return Files.createTempDirectory(prefix);
  }
}
