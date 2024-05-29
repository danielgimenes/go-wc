package count_ops

import (
	"fmt"
	"os"
	"strings"
)

const FILE_READ_BUFFER_SIZE = 4096 // 4kb

func PrintFileByteCount(fileInfo os.FileInfo) {
	fmt.Println(fileInfo.Size(), fileInfo.Name())
}

func PrintFileNewlineCount(file *os.File) {
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
	fmt.Println(newlines, file.Name())
}

func PrintFileWordCount(fileInfo os.FileInfo, file *os.File) {
	words := 0
	data := make([]byte, fileInfo.Size()) // read the whole file to avoid split words
	readBytes, fileReadErr := file.Read(data)
	for fileReadErr == nil && readBytes != 0 {
		strRead := string(data[:readBytes])
		words += len(strings.Fields(strRead))
		readBytes, fileReadErr = file.Read(data)
	}
	fmt.Println(words, file.Name())
}
