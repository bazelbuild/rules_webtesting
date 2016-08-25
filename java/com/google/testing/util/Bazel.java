/* Copyright 2016 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package com.google.testing.util;

import com.google.common.annotations.VisibleForTesting;
import com.google.common.base.Preconditions;
import java.nio.file.FileSystems;
import java.nio.file.Path;

public class Bazel {

  private static Bazel instance;

  private static final String RUNFILES_VAR = "TEST_SRCDIR";

  private final Path runfilesDir;

  @VisibleForTesting
  Bazel(Path runfilesDir) {
    this.runfilesDir = Preconditions.checkNotNull(runfilesDir);
  }

  public static Bazel getInstance() {
    if (instance == null) {
      instance = new Bazel(FileSystems.getDefault().getPath(System.getenv(RUNFILES_VAR)));
    }
    return instance;
  }

  // Returns the path of the runfile in the default wworkspace.
  public Path runfile(String path) {
    return runfilesDir.resolve(path);
  }
}
