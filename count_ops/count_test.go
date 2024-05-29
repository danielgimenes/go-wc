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
	fileContent := "abc123\nSomething\t\n"
	fileInfo, file := createTempFile(fileContent)
	expected := int64(len(fileContent))
	result := FileByteCount(fileInfo)
	tempFileCleanup(file)
	if expected != result {
		t.Errorf("Expected %d, received %v", expected, result)
	}
}
