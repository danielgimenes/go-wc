// TODO organize error code, buffered read and counts into separate packages?
// TODO add unit tests

package main

import (
	"fmt"
	"os"
	"strings"
)

type ErrorMessage string

const (
	NO_ARGS_SUPPLIED_ERR   ErrorMessage = "no arguments were supplied."
	INVALID_ARGUMENTS_ERR  ErrorMessage = "invalid arguments."
	CANT_OPEN_FILE_ERR     ErrorMessage = "can't open file."
	CANT_GET_FILE_INFO_ERR ErrorMessage = "can't get file info."
	CLOSE_FILE_ERR         ErrorMessage = "can't close file"
)

type Operation string

const (
	BYTE_COUNT_OPERATION    Operation = "-c"
	CHAR_COUNT_OPERATION    Operation = "-m"
	NEWLINE_COUNT_OPERATION Operation = "-l"
	WORD_COUNT_OPERATION    Operation = "-w"
)

const FILE_READ_BUFFER_SIZE = 4096 // 4kb

func exitWithError(message ErrorMessage) {
	fmt.Println("Error:", message)
	os.Exit(1)
}

func readCommandLineArgs() (Operation, string) {
	if len(os.Args) == 0 {
		exitWithError(NO_ARGS_SUPPLIED_ERR)
	}
	if len(os.Args) != 3 {
		exitWithError(INVALID_ARGUMENTS_ERR)
	}
	operationArg := os.Args[1]
	filePath := os.Args[2]
	return Operation(operationArg), filePath
}

func openFile(filePath string) (os.FileInfo, *os.File) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		exitWithError(CANT_GET_FILE_INFO_ERR)
	}
	file, err := os.Open(filePath)
	if err != nil {
		exitWithError(CANT_OPEN_FILE_ERR)
	}
	return fileInfo, file
}

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

func closeFile(file *os.File) {
	closeErr := file.Close()
	if closeErr != nil {
		exitWithError(CLOSE_FILE_ERR)
	}
}

func main() {
	operationArg, filePath := readCommandLineArgs()
	fileInfo, file := openFile(filePath)
	switch operationArg {
	case BYTE_COUNT_OPERATION, CHAR_COUNT_OPERATION:
		printFileByteCount(fileInfo)
	case NEWLINE_COUNT_OPERATION:
		printFileNewlineCount(file)
	case WORD_COUNT_OPERATION:
		printFileWordCount(fileInfo, file)
	default:
		exitWithError(INVALID_ARGUMENTS_ERR)
	}
	closeFile(file)
}
