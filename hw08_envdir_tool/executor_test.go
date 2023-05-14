package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
)

func TestRunCmd(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dir, err := os.MkdirTemp("testdata", "tmp")
		require.NoError(t, err)
		defer os.RemoveAll(dir)
		err = os.Mkdir(filepath.Join(dir, "vars"), 0o777)
		require.NoError(t, err)
		err = ioutil.WriteFile(filepath.Join(dir, "vars", "FOO"), []byte("value from file"), 0o666)
		require.NoError(t, err)
		err = ioutil.WriteFile(filepath.Join(dir, "t.sh"), []byte("#!/usr/bin/env bash\necho $1\necho $FOO\n"), 0o666)
		require.NoError(t, err)
		err = os.Chmod(filepath.Join(dir, "t.sh"), 0o777)
		require.NoError(t, err)

		env, err := ReadDir(filepath.Join(dir, "vars"))
		require.NoError(t, err)

		var returnCode int
		result := capturer.CaptureStdout(func() {
			returnCode = RunCmd([]string{filepath.Join(dir, "t.sh"), "value from cmd"}, env)
		})
		require.Equal(t, 0, returnCode)
		require.Equal(t, "value from cmd\nvalue from file\n", result)
	})

	t.Run("error executing command", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)

		var returnCode int
		result := capturer.CaptureStderr(func() {
			returnCode = RunCmd([]string{"cat", "."}, env)
		})
		require.Equal(t, 1, returnCode)
		require.Equal(t, "cat: .: Is a directory\n", result)
	})
}
