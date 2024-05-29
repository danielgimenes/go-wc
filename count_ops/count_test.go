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
