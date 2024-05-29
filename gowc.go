package main

import (
	"fmt"
	"os"
)

const NO_ARGS_SUPPLIED_ERR = "no arguments were supplied."
const INVALID_ARGUMENTS_ERR = "invalid arguments."
const CANT_OPEN_FILE_ERR = "can't open file."
const CANT_GET_FILE_INFO_ERR = "can't get file info."
const CLOSE_FILE_ERR = "can't close file"

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

	switch operationArg {
	case "-c":
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			exitWithError(CANT_GET_FILE_INFO_ERR)
		}
		fmt.Println(fileInfo.Size(), fileInfo.Name())
	case "-l":
		file, err := os.Open(filePath)
		if err != nil {
			exitWithError(CANT_OPEN_FILE_ERR)
		}
		closeErr := file.Close()
		if closeErr != nil {
			exitWithError(CLOSE_FILE_ERR)
		}
	}

}
