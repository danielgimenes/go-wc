package main

import (
	"fmt"
	"os"
	"strings"
)

type Operation string

const (
	BYTE_COUNT_OPERATION    Operation = "-c"
	CHAR_COUNT_OPERATION    Operation = "-m"
	NEWLINE_COUNT_OPERATION Operation = "-l"
	WORD_COUNT_OPERATION    Operation = "-w"
)

const FILE_READ_BUFFER_SIZE = 4096 // 4kb

func printFileByteCount(fileInfo os.FileInfo) {
	fmt.Println(fileInfo.Size(), fileInfo.Name())
}

func printFileNewlineCount(file *os.File) {
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

func printFileWordCount(fileInfo os.FileInfo, file *os.File) {
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
