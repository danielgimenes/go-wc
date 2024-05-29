package main

import "os"

func openFile(filePath string) (os.FileInfo, *os.File) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		exitWithError(CANT_GET_FILE_INFO_ERR, err)
	}
	file, err := os.Open(filePath)
	if err != nil {
		exitWithError(CANT_OPEN_FILE_ERR, err)
	}
	return fileInfo, file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		exitWithError(CLOSE_FILE_ERR, err)
	}
}
