package main

import (
	"fmt"
	"os"
)

type ErrorMessage string

const (
	NO_ARGS_SUPPLIED_ERR   ErrorMessage = "no arguments were supplied."
	INVALID_ARGUMENTS_ERR  ErrorMessage = "invalid arguments."
	CANT_OPEN_FILE_ERR     ErrorMessage = "can't open file."
	CANT_GET_FILE_INFO_ERR ErrorMessage = "can't get file info."
	CLOSE_FILE_ERR         ErrorMessage = "can't close file"
)

func exitWithError(message ErrorMessage, err error) {
	fmt.Printf("Error: %s. Details: %s\n", message, err)
	os.Exit(1)
}
