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
