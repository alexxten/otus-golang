package main

import (
	"testing"

	"os"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("positive case, copy all file", func(t *testing.T) {
		exp := "testdata/out_offset0_limit0.txt"
		dst := "testdata/case_copy_all.txt"
		defer os.Remove(dst)
		err := Copy("testdata/input.txt", dst, 0, 0)
		require.Nil(t, err)
		expectedFile, err := os.ReadFile(exp)
		require.Nil(t, err)
		actualFile, err := os.ReadFile(dst)
		require.Nil(t, err)
		require.Equal(t, expectedFile, actualFile)
	})

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy(
			"/dev/urandom",
			"testdata/input2.txt",
			0,
			0,
		)
		require.Equal(t, err, ErrUnsupportedFile)
	})
	t.Run("empty file", func(t *testing.T) {
		err := Copy(
			"testdata/input_empty.txt",
			"testdata/input2.txt",
			0,
			0,
		)
		require.Equal(t, err, ErrEmptyFile)
	})

	t.Run("positive case, offset", func(t *testing.T) {
		dst := "testdata/case_offset.txt"
		exp := "testdata/out_offset6000_limit0.txt"
		defer os.Remove(dst)
		err := Copy("testdata/input.txt", dst, 6000, 0)
		require.Nil(t, err)
		expectedFile, err := os.ReadFile(exp)
		require.Nil(t, err)
		actualFile, err := os.ReadFile(dst)
		require.Nil(t, err)
		require.Equal(t, expectedFile, actualFile)
	})

	t.Run("positive case, limit", func(t *testing.T) {
		dst := "testdata/case_limit.txt"
		exp := "testdata/out_offset0_limit1000.txt"
		defer os.Remove(dst)
		err := Copy("testdata/input.txt", dst, 0, 1000)
		require.Nil(t, err)
		expectedFile, err := os.ReadFile(exp)
		require.Nil(t, err)
		actualFile, err := os.ReadFile(dst)
		require.Nil(t, err)
		require.Equal(t, expectedFile, actualFile)
	})
}
