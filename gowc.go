package main

import (
	"fmt"
	"os"
)

const NO_ARGS_SUPPLIED_ERR = "Error: no arguments were supplied."
const INVALID_ARGUMENTS_ERR = "Error: invalid arguments."
const CANT_OPEN_FILE_ERR = "Error: can't open file."
const CANT_GET_FILE_INFO_ERR = "Error: can't get file info."

func main() {
	if len(os.Args) == 0 {
		fmt.Println(NO_ARGS_SUPPLIED_ERR)
		return
	}

	if len(os.Args) == 3 && os.Args[1] == "-c" {
		filePath := os.Args[2]
		// file, err := os.Open(filePath)
		// if err != nil {
		// 	fmt.Println(CANT_OPEN_FILE_ERR)
		// }
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			fmt.Println(CANT_GET_FILE_INFO_ERR)
		}
		fmt.Println(fileInfo.Size(), fileInfo.Name())
		return
	} else {
		fmt.Println(INVALID_ARGUMENTS_ERR)
		return
	}

}
