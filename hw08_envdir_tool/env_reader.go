package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadFile(dir, fileName string) (string, error) {
	f, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	if !s.Scan() {
		return "", nil
	}
	line := s.Text()
	line = strings.ReplaceAll(line, "\x00", "\n")
	line = strings.TrimRightFunc(line, unicode.IsSpace)
	return line, nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	result := make(Environment)
	for _, file := range files {
		if strings.Contains(file.Name(), "=") {
			continue
		}
		if !file.Mode().IsRegular() {
			continue
		}

		if file.Size() == 0 {
			result[file.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		value, err := ReadFile(dir, file.Name())
		if err != nil {
			return nil, err
		}
		result[file.Name()] = EnvValue{Value: value}
	}

	return result, nil
}
