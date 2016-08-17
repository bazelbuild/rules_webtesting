package bazel

import (
	"fmt"
	"os"
	"path/filepath"
)

const runfileEnv = "TEST_SRCDIR"

// Runfile returns an absolute path to the specified file in the runfiles directory of the running target.
func Runfile(path string) (string, error) {
	runfileDir, ok := os.LookupEnv(runfileEnv)
	if !ok {
		return "", fmt.Errorf("environment variable %q is not defined, are you running with bazel test", runfileEnv)
	}
	return filepath.Join(runfileDir, path), nil
}
