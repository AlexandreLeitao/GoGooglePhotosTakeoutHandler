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
	//dirToIterate := "B:\\GoogleTakeout\\Google Photos\\Extracted"
	dirToIterate := "C:\\Users\\alexandre.leitao\\OneDrive - Havas\\Documents\\TestFolder"
	rootToProcessTo := "C:\\Users\\alexandre.leitao\\OneDrive - Havas\\Documents\\TestFolder\\Processed"
	// iterate(dirToIterate)
	prepareCommonStructure(dirToIterate, rootToProcessTo)
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
			//fmt.Printf("Folder: %s\n", info.Name())
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

func getFoldersStrings(folders []os.FileInfo) []string {
	strings := make([]string, 0)
	for _, i := range folders {
		strings = append(strings, i.Name())
	}
	return strings
}

func SliceContains(slice []string, value string) bool {
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

func prepareCommonStructure(dirToIterate string, rootToProcessTo string) {
	allFolders := getFolders(dirToIterate)
	allFoldersStrings := getFoldersStrings(allFolders)
	parentFolders := getParentFolders(allFolders)
	parentFoldersStrings := getFoldersStrings(parentFolders)

	distinctFolders := make([]string, 0)

	for _, x := range allFoldersStrings {
		if !SliceContains(distinctFolders, x) && !SliceContains(parentFoldersStrings, x) {
			distinctFolders = append(distinctFolders, x)
		}
	}

	if !exists(rootToProcessTo) {
		os.Mkdir(rootToProcessTo, 0700)
	}

	//Create all new Structure
	for _, i := range distinctFolders {
		tempPath := rootToProcessTo + "\\" + i
		if !exists(tempPath) {
			os.Mkdir(tempPath, 0700)
		}
		fmt.Println(i)
	}
}
