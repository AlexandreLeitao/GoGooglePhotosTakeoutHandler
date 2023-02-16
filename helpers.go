package main

import (
	"fmt"
	"log"
	"os"
)

func sliceContains(slice []string, value string) bool {
	for _, x := range slice {
		if x == value {
			return true
		}
	}
	return false
}

// exists returns whether the given file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return false
}

func moveFile(file string, oldPath string, newPath string) bool {

	err := os.Rename(oldPath+"\\"+file, newPath+"\\"+file)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
