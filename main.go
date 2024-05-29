// TODO add unit tests

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
)

func readCommandLineArgs() (Operation, string) {
	if len(os.Args) == 0 {
		exitWithError(NO_ARGS_SUPPLIED_ERR, nil)
	}
	if len(os.Args) != 3 {
		exitWithError(INVALID_ARGUMENTS_ERR, nil)
	}
	operationArg := os.Args[1]
	filePath := os.Args[2]
	return Operation(operationArg), filePath
}

func main() {
	operationArg, filePath := readCommandLineArgs()
	fileInfo, file := openFile(filePath)
	switch operationArg {
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
