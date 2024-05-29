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

func exitWithError(errorMessage string) {
	fmt.Println("Error:", errorMessage)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 0 {
		exitWithError(NO_ARGS_SUPPLIED_ERR)
	}

	if len(os.Args) != 3 {
		exitWithError(INVALID_ARGUMENTS_ERR)
	}
	operationArg := os.Args[1]
	filePath := os.Args[2]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		exitWithError(CANT_GET_FILE_INFO_ERR)
	}

	switch operationArg {
	case "-c":
		fmt.Println(fileInfo.Size(), fileInfo.Name())
	case "-l":
		file, err := os.Open(filePath)
		if err != nil {
			exitWithError(CANT_OPEN_FILE_ERR)
		}
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
		closeErr := file.Close()
		fmt.Println(newlines, file.Name())
		if closeErr != nil {
			exitWithError(CLOSE_FILE_ERR)
		}
	case "-w":
		file, err := os.Open(filePath)
		if err != nil {
			exitWithError(CANT_OPEN_FILE_ERR)
		}
		words := 0
		data := make([]byte, fileInfo.Size()) // read the whole file
		readBytes, fileReadErr := file.Read(data)
		for fileReadErr == nil && readBytes != 0 {
			strRead := string(data[:readBytes])
			words += len(strings.Fields(strRead))
			readBytes, fileReadErr = file.Read(data)
		}

		closeErr := file.Close()
		fmt.Println(words, file.Name())
		if closeErr != nil {
			exitWithError(CLOSE_FILE_ERR)
		}
	default:
		exitWithError(INVALID_ARGUMENTS_ERR)
	}

}
