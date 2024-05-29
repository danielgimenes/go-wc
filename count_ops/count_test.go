package count_ops

import (
	"fmt"
	"os"
	"testing"
)

func createTempFile(fileContent string) (os.FileInfo, *os.File) {
	file, createErr := os.CreateTemp("", "test_*")
	if createErr != nil {
		panic(fmt.Sprintf("Couldn't create temp file. Details: %s", createErr))
	}
	writtenCount, writeErr := file.WriteString(fileContent)
	if writtenCount != len(fileContent) || writeErr != nil {
		panic(fmt.Sprintf("Error writing content to temp file. Details: %s", writeErr))
	}
	file.Seek(0, 0)
	fileInfo, statErr := file.Stat()
	if statErr != nil {
		panic(fmt.Sprintf("Couldn't get file stats of temp file. Details: %s", statErr))
	}
	return fileInfo, file
}

func tempFileCleanup(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(fmt.Sprintf("Unable to close temp file. Details: %s", err))
	}
	os.Remove(file.Name())
}

func TestFileByteCount(t *testing.T) {
	var tests = []struct {
		name        string
		fileContent string
		want        int64
	}{
		{"many chars and symbols", "abc 123\nSomething.!?\t\n@", 23},
		{"single letter", "a", 1},
		{"no content", "", 0},
		{"complex", "\n\t\b", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInfo, file := createTempFile(tt.fileContent)
			result := FileByteCount(fileInfo)
			tempFileCleanup(file)
			if result != tt.want {
				t.Errorf("Expected %d, received %v", tt.want, result)
			}
		})
	}
}

func TestFileNewlineCount(t *testing.T) {
	var tests = []struct {
		name        string
		fileContent string
		want        int
	}{
		{"long content, some newlines", "abc 123\nSomething.!?\t\n@", 2},
		{"long content, no newlines", "That is a long text, right? right? :)", 0},
		{"no content", "", 0},
		{"just newlines", "\n\n\n", 3},
		{"newline before and after", "\nabc\ndef\n", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, file := createTempFile(tt.fileContent)
			result := FileNewlineCount(file)
			tempFileCleanup(file)
			if result != tt.want {
				t.Errorf("Expected %d, received %v", tt.want, result)
			}
		})
	}
}
