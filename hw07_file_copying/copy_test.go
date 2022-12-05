package main

import "testing"

func TestCopy(t *testing.T) {
	t.Run("positive case", func(t *testing.T) {
		err := Copy(
			"testdata/input.txt",
			"testdata/input2.txt",
			0,
			100,
		)
		print(err)
	})

	t.Run("zero file size", func(t *testing.T) {
		err := Copy(
			"/dev/urandom",
			"testdata/input2.txt",
			0,
			100,
		)
		print(err)
	})
}
