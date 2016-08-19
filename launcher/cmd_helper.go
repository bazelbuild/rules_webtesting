// Package cmdhelper provides functions to make working os/exec Command easier.
package cmdhelper

import (
	"os"
	"strings"
)

// UpdateEnv takes an environment array, an env var name, and an environment
// var value and updates the current definition of the env var or adds a new
// environment variable definition.
func UpdateEnv(env []string, name, value string) []string {
	prefix := name + "="

	for i := 0; i < len(env); {
		if strings.HasPrefix(env[i], prefix) {
			env = append(env[:i], env[i+1:]...)
		} else {
			i++
		}
	}

	return append(env, prefix+value)
}

// BulkUpdateEnv returns a new slice suitable for use with exec.Cmd.Env
// based on env but with environment variables for the keys in update
// added/changed to the values in update.
func BulkUpdateEnv(env []string, update map[string]string) []string {
	for name, value := range update {
		env = UpdateEnv(env, name, value)
	}
	return env
}

// IsTruthyEnv return true if name is set in the environment and has value other than
// 0 or false (case-insensitive).
func IsTruthyEnv(name string) bool {
	value := os.Getenv(name)
	return value != "" && value != "0" && strings.ToLower(value) != "false"
}
