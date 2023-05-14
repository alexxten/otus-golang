package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
)

func TestRunCmd(t *testing.T) {
	t.Run("testdata, success", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)
		err = os.Chmod("test.sh", 0o777)
		require.NoError(t, err)

		var returnCode int
		result := capturer.CaptureStdout(func() {
			returnCode = RunCmd([]string{"sh", "test.sh"}, env)
		})
		require.Equal(t, 0, returnCode)
		require.Equal(t, "PASS\n", result)
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
