package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentDirectory)
	dirToIterate := "B:\\GoogleTakeout\\Google Photos\\Extracted"

	// iterate(dirToIterate)

	allFolders := getFolders(dirToIterate)
	parentFolders := getParentFolders(allFolders)

	for _, i := range parentFolders {
		fmt.Println(i.Name())
	}
}

func iterate(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if info.IsDir() {
			fmt.Printf("Folder: %s\n", info.Name())
		} else {
			fmt.Printf("File : %s\n", info.Name())
		}

		return nil
	})
}

func getFolders(path string) []os.FileInfo {

	folders := make([]os.FileInfo, 0)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if info.IsDir() {
			folders = append(folders, info)
			fmt.Printf("Folder: %s\n", info.Name())
		}
		return nil
	})
	return folders
}

func getParentFolders(folders []os.FileInfo) []os.FileInfo {
	parentFolders := make([]os.FileInfo, 0)

	for _, i := range folders {
		match, _ := regexp.MatchString("takeout-.{16}-.{3}", i.Name())

		if match {
			parentFolders = append(parentFolders, i)
		}
	}
	return parentFolders
}
