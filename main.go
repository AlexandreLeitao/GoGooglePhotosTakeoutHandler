package main

import (
	"bufio"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input path to operate in:")
	// dirToIterateInput, _ := reader.ReadString('\n')
	dirToIterateInput := "C:\\Users\\alexs\\Downloads\\GoogleTakeout\\Extracted"
	fmt.Println("Do you want to copy files instead of moving them? (Y/N)")
	isCopyInput, _ := reader.ReadString('\n')

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentDirectory)

	//Configurations
	isCopy := strings.EqualFold(isCopyInput, "Y")
	dirToIterate := filepath.Clean(dirToIterateInput)
	rootToProcessTo := filepath.Join(dirToIterateInput, "Processed")

	if !exists(rootToProcessTo) {
		fmt.Println("Creating folder:", rootToProcessTo)
		os.MkdirAll(rootToProcessTo, 0755)
	} else {
		fmt.Println("Folder already exists:", rootToProcessTo)
	}

	totalFiles, totalFolders := iteratePreProcessing(rootToProcessTo)
	fmt.Printf("Total files: %d, Total folders: %d\n", totalFiles, totalFolders)

	//Setup
	prepareCommonStructure(dirToIterate, rootToProcessTo)

	//Moving Files
	moveAllFilesToCommonStructure(dirToIterate, rootToProcessTo, isCopy)

	//Finish File Date Handling
	iteratePostProcessing(rootToProcessTo)

	fmt.Println("The end!")
}

// Calculates Total Files and Folders to be processed by the application
func iteratePreProcessing(dirToIterate string) (int, int) {

	totalFiles, totalFolders := 0, 0

	filepath.Walk(dirToIterate, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if info.IsDir() {
			totalFolders += 1
			//fmt.Printf("Folder: %s\n", info.Name())
		} else {
			totalFiles += 1
			//fmt.Printf("File : %s\n", info.Name())
			//fmt.Printf("%+v \r", info)
		}

		return nil
	})
	return totalFiles, totalFolders
}

// Returns All folders in path in a slice of fileInfos
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

// Returns all "Parent Folders" that obey the regex rule
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

// Converts slice of FileInfos to slice of Strings with each fileinfo name
func getFoldersStrings(folders []os.FileInfo) []string {
	strings := make([]string, 0)
	for _, i := range folders {
		strings = append(strings, i.Name())
	}
	return strings
}

// Creates or updates the Common Structure
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

// Moves all files to correspondent Folder
func moveAllFilesToCommonStructure(dirToIterate string, rootToProcessTo string, isCopy bool) {
	//Maps Correspondent folders in the unified strructure
	of, _ := os.Lstat(dirToIterate)
	nf, _ := os.Lstat(rootToProcessTo)
	mapperObj := mapper{
		orginalFolder: of,
		originalPath:  dirToIterate,
		newFolder:     nf,
		newPath:       rootToProcessTo,
	}

	filepath.Walk(dirToIterate, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if !info.IsDir() {
			fileName := info.Name()
			parentFolder := filepath.Dir(path)

			parentFolderName := filepath.Base(parentFolder)
			fmt.Println(parentFolderName)
			fmt.Printf("File: %s\n", fileName)

			_, newFolderPath := mapperObj.getCorrespondentFolder(parentFolderName)

			if !FileExists(fileName, newFolderPath) {
				moveFile(fileName, parentFolder, newFolderPath, isCopy)
			} else {
				fmt.Printf("%s Skipped\n", fileName)
			}
		}
		return nil
	})

}

// Updates all files in Final Folder with their correspondent dates from json
func iteratePostProcessing(dirToIterate string) {

	filepath.Walk(dirToIterate, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if !info.IsDir() {
			fileName := info.Name()
			fileExtension := filepath.Ext(fileName)
			if fileExtension != ".json" {
				//GetJson and update file date
				finalJsonPath := findMetadataFile(path)

				if _, err := os.Stat(finalJsonPath); !os.IsNotExist(err) {
					resultJsonObj := GetFileJson(finalJsonPath)
					if resultJsonObj.PhotoTakenTime.Timestamp != "" {
						date, _ := strconv.ParseInt(resultJsonObj.PhotoTakenTime.Timestamp, 10, 64)
						UpdateFileDateWithTimeStamp(path, date)
						fmt.Print("- PhotoTakenTime.Timestamp")
					} else {
						if resultJsonObj.CreationTime.Timestamp != "" {
							date, _ := strconv.ParseInt(resultJsonObj.PhotoTakenTime.Timestamp, 10, 64)
							UpdateFileDateWithTimeStamp(path, date)
							fmt.Print("- CreationTime.Timestamp")
						} else {
							fmt.Println(path, " - Json Not found")
						}
					}
				}
			}
		} else {
			fmt.Println("Folder: " + info.Name())
		}
		return nil
	})

}

func findMetadataFile(basePath string) string {
	// Check for any file ending in ".metadata.json" in the same directory
	dir := filepath.Dir(basePath)
	baseFileName := filepath.Base(basePath)
	pattern := filepath.Join(dir, baseFileName+".*.json")

	// Use glob to find all matching patterns
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Println("Error matching pattern:", err)
		return ""
	}
	// Return the first match found, if any
	if len(matches) > 0 {
		return matches[0]
	}

	// If no matches found, look for files with a similar base name and `.json` suffix
	pattern = filepath.Join(dir, baseFileName+"*.json")
	matches, err = filepath.Glob(pattern)
	if err != nil {
		log.Println("Error matching pattern:", err)
		return ""
	}
	if len(matches) > 0 {
		// Assuming we take the first match based on variations
		return matches[0]
	}
	return ""
}
