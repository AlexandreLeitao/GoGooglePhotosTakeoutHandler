package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type mapper struct {
	orginalFolder os.FileInfo
	originalPath  string
	newFolder     os.FileInfo
	newPath       string
}

func Test() {
	fmt.Println("Test")
}

func (m mapper) getCorrespondentFolder(folderName string) (os.FileInfo, string) {

	var tempDir os.FileInfo
	tempPath := ""

	filepath.Walk(m.newPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if info.IsDir() {
			if info.Name() == folderName {
				tempPath = m.newPath + "\\" + info.Name()
				tempDir = info
			}
		} else {
			return nil
		}
		return nil
	})

	return tempDir, tempPath
}
