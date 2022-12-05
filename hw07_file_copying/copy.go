package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrEmptyFile             = errors.New("zero file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// get source file size
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	// check file size
	if fileSize == 0 {
		return ErrEmptyFile
	}
	// check offset val
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	nBytes, err := io.Copy(dst, src)
	print(nBytes)
	return err
}
