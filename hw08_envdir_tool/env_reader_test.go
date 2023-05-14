package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("testdata, success", func(t *testing.T) {
		expectEnv := Environment{
			"BAR":   EnvValue{Value: "bar"},
			"EMPTY": EnvValue{NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line"},
			"HELLO": EnvValue{Value: `"hello"`},
			"UNSET": EnvValue{NeedRemove: true},
		}
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)
		require.Equal(t, env, expectEnv)
	})

	t.Run("= in filename", func(t *testing.T) {
		dir, err := ioutil.TempDir("testdata", "tmp")
		require.NoError(t, err)
		defer os.RemoveAll(dir)

		err = ioutil.WriteFile(filepath.Join(dir, "smth=="), []byte("bar"), 0666)
		require.NoError(t, err)

		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Len(t, env, 0)
	})

	t.Run("empty dir", func(t *testing.T) {
		dir, err := ioutil.TempDir("testdata", "tmp")
		require.NoError(t, err)
		defer os.RemoveAll(dir)

		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Len(t, env, 0)
	})

	t.Run("dir does not exist", func(t *testing.T) {
		_, err := ReadDir("some name")
		require.Error(t, err)
	})
}
