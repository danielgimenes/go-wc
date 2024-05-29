package main

import (
	"fmt"
	"os"

	"github.com/danielgimenes/gowc/count_ops"
)

type Operation string

const (
	BYTE_COUNT_OPERATION    Operation = "-c"
	CHAR_COUNT_OPERATION    Operation = "-m"
	NEWLINE_COUNT_OPERATION Operation = "-l"
	WORD_COUNT_OPERATION    Operation = "-w"
	ALL_OPERATION           Operation = ""
)

func readCommandLineArgs() (Operation, string) {
	var operationArg Operation
	var filePath string
	switch len(os.Args) {
	case 1:
		exitWithError(NO_ARGS_SUPPLIED_ERR, nil)
	case 2:
		operationArg = ALL_OPERATION
		filePath = os.Args[1]
	case 3:
		operationArg = Operation(os.Args[1])
		filePath = os.Args[2]
	default:
		exitWithError(INVALID_ARGUMENTS_ERR, nil)
	}
	return operationArg, filePath
}

func main() {
	operationArg, filePath := readCommandLineArgs()
	fileInfo, file := openFile(filePath)
	switch operationArg {
	case ALL_OPERATION:
		newlineCount := count_ops.FileNewlineCount(file)
		file.Seek(0, 0)
		wordCount := count_ops.FileWordCount(fileInfo, file)
		byteCount := count_ops.FileByteCount(fileInfo)
		fmt.Println(newlineCount, wordCount, byteCount, fileInfo.Name())
	case BYTE_COUNT_OPERATION, CHAR_COUNT_OPERATION:
		byteCount := count_ops.FileByteCount(fileInfo)
		fmt.Println(byteCount, fileInfo.Name())
	case NEWLINE_COUNT_OPERATION:
		newlineCount := count_ops.FileNewlineCount(file)
		fmt.Println(newlineCount, fileInfo.Name())
	case WORD_COUNT_OPERATION:
		wordCount := count_ops.FileWordCount(fileInfo, file)
		fmt.Println(wordCount, fileInfo.Name())
	default:
		exitWithError(INVALID_ARGUMENTS_ERR, nil)
	}
	closeFile(file)
}
