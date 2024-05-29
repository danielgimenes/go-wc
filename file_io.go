package main

import "os"

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

func closeFile(file *os.File) {
	closeErr := file.Close()
	if closeErr != nil {
		exitWithError(CLOSE_FILE_ERR)
	}
}
