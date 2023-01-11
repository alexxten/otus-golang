package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func getNewFileSize(fileSize int64, offset int64, limit int64) int64 {
	n := fileSize
	if offset > 0 {
		n -= offset
	}
	if limit > 0 && n > limit {
		n = limit
	}
	return n
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	// get file info to check it
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}
	fileSize := fileInfo.Size()

	// check offset val
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer src.Close()

	n := getNewFileSize(fileSize, offset, limit)
	if offset > 0 {
		_, err = src.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	// copy to dst
	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	bar := pb.Full.Start64(n)
	barReader := bar.NewProxyReader(src)
	_, err = io.CopyN(dst, barReader, n)
	bar.Finish()
	return err
}
