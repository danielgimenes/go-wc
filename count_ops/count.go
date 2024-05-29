package count_ops

import (
	"os"
	"strings"
)

const FILE_READ_BUFFER_SIZE = 4096 // 4kb

func FileByteCount(fileInfo os.FileInfo) int64 {
	return fileInfo.Size()
}

func FileNewlineCount(file *os.File) int {
	newlines := 0
	data := make([]byte, FILE_READ_BUFFER_SIZE)
	readBytes, fileReadErr := file.Read(data)
	for fileReadErr == nil && readBytes != 0 {
		for _, c := range string(data[:readBytes]) {
			if c == '\n' {
				newlines++
			}
		}
		readBytes, fileReadErr = file.Read(data)
	}
	return newlines
}

func FileWordCount(fileInfo os.FileInfo, file *os.File) int {
	words := 0
	data := make([]byte, fileInfo.Size()) // read the whole file to avoid split words
	readBytes, fileReadErr := file.Read(data)
	for fileReadErr == nil && readBytes != 0 {
		strRead := string(data[:readBytes])
		words += len(strings.Fields(strRead))
		readBytes, fileReadErr = file.Read(data)
	}
	return words
}
