package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyPositive(t *testing.T) {
	tests := []struct {
		src            string
		dst            string
		offset         int64
		limit          int64
		expectedOutput string
	}{
		{src: "testdata/input.txt", dst: "testdata/case_copy_all.txt", offset: 0, limit: 0, expectedOutput: "testdata/out_offset0_limit0.txt"},
		{src: "testdata/input.txt", dst: "testdata/case_offset.txt", offset: 6000, limit: 0, expectedOutput: "testdata/out_offset6000_limit0.txt"},
		{src: "testdata/input.txt", dst: "testdata/case_limit.txt", offset: 0, limit: 1000, expectedOutput: "testdata/out_offset0_limit1000.txt"},
		{src: "testdata/input.txt", dst: "testdata/case_offset_limit.txt", offset: 100, limit: 1000, expectedOutput: "testdata/out_offset100_limit1000.txt"},
		{src: "testdata/input_empty.txt", dst: "testdata/case_empty.txt", offset: 0, limit: 0, expectedOutput: "testdata/out_empty.txt"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.expectedOutput, func(t *testing.T) {
			defer os.Remove(tc.dst)
			err := Copy(tc.src, tc.dst, tc.offset, tc.limit)
			require.Nil(t, err)
			expectedFile, err := os.ReadFile(tc.expectedOutput)
			require.Nil(t, err)
			actualFile, err := os.ReadFile(tc.dst)
			require.Nil(t, err)
			require.Equal(t, expectedFile, actualFile)
		})
	}
}

func TestCopyNegative(t *testing.T) {
	tests := []struct {
		src           string
		dst           string
		offset        int64
		limit         int64
		expectedError error
	}{
		{src: "/dev/urandom", dst: "testdata/case_unsupported.txt", offset: 0, limit: 0, expectedError: ErrUnsupportedFile},
		{src: "testdata/input.txt", dst: "testdata/case_offset_more_than_file.txt", offset: 10000, limit: 0, expectedError: ErrOffsetExceedsFileSize},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.expectedError.Error(), func(t *testing.T) {
			err := Copy(tc.src, tc.dst, tc.offset, tc.limit)
			require.Equal(t, err, tc.expectedError)
		})
	}
}
