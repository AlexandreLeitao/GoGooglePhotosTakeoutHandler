package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	//To be replaced with prompt
	isCopy := true

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentDirectory)
	//dirToIterate := "B:\\GoogleTakeout\\Google Photos\\Extracted"
	// rootToProcessTo := "B:\\GoogleTakeout\\Google Photos\\Extracted\\Processed"
	dirToIterate := "C:\\Users\\alexandre.leitao\\OneDrive - Havas\\Documents\\TestFolder"
	rootToProcessTo := "C:\\Users\\alexandre.leitao\\OneDrive - Havas\\Documents\\TestFolder\\Processed"
	// iterate(dirToIterate)

	prepareCommonStructure(dirToIterate, rootToProcessTo)

	of, _ := os.Lstat(dirToIterate)
	nf, _ := os.Lstat(rootToProcessTo)

	mapperObj := mapper{
		orginalFolder: of,
		originalPath:  dirToIterate,
		newFolder:     nf,
		newPath:       rootToProcessTo,
	}

	x, y := mapperObj.getCorrespondentFolder("Important Shit")

	moveFile("The-Dark-Angel-HD-Wallpaper-HD-1080p3.jpg", "C:\\Users\\alexandre.leitao\\OneDrive - Havas\\Documents\\TestFolder\\takeout-20230205T163750Z-001\\Google Fotos\\Photos from 2007", "C:\\Users\\alexandre.leitao\\OneDrive - Havas\\Documents\\TestFolder\\Processed\\Photos from 2007", isCopy)

	// fmt.Print(createDummyStructure())
	fmt.Printf("%+v \r\r\r", x)
	fmt.Println(y)

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

func prepareCommonStructure(dirToIterate string, rootToProcessTo string) {
	allFolders := getFolders(dirToIterate)
	allFoldersStrings := getFoldersStrings(allFolders)
	parentFolders := getParentFolders(allFolders)
	parentFoldersStrings := getFoldersStrings(parentFolders)

	distinctFolders := make([]string, 0)

	for _, x := range allFoldersStrings {
		if !sliceContains(distinctFolders, x) && !sliceContains(parentFoldersStrings, x) {
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
