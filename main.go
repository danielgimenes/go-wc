// TODO add unit tests

package main

import (
	"os"
)

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
