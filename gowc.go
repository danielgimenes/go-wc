package main

import (
	"fmt"
	"os"
	"strings"
)

const NO_ARGS_SUPPLIED_ERR = "no arguments were supplied."
const INVALID_ARGUMENTS_ERR = "invalid arguments."
const CANT_OPEN_FILE_ERR = "can't open file."
const CANT_GET_FILE_INFO_ERR = "can't get file info."
const CLOSE_FILE_ERR = "can't close file"
const FILE_READ_BUFFER_SIZE = 4096 // 4kb

// TODO organize error code, buffered read and counts into modules
// TODO add unit tests

func exitWithError(errorMessage string) {
	fmt.Println("Error:", errorMessage)
	os.Exit(1)
}

func readCommandLineArgs() (string, string) {
	if len(os.Args) == 0 {
		exitWithError(NO_ARGS_SUPPLIED_ERR)
	}

	if len(os.Args) != 3 {
		exitWithError(INVALID_ARGUMENTS_ERR)
	}
	operationArg := os.Args[1]
	filePath := os.Args[2]
	return operationArg, filePath
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
	data := make([]byte, fileInfo.Size()) // read the whole file
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
	case "-c", "-m":
		printFileByteCount(fileInfo)
	case "-l":
		printFileNewlineCount(file)
	case "-w":
		printFileWordCount(fileInfo, file)
	default:
		exitWithError(INVALID_ARGUMENTS_ERR)
	}
	closeFile(file)
}
