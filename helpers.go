package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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

// Move file based on paramenters either moves a file from a src to a dest or copies the file from the src to the dest
func moveFile(file string, oldPath string, newPath string, isCopy bool) bool {

	if isCopy {
		bytesRead, err := ioutil.ReadFile(oldPath + "\\" + file)

		if err != nil {
			log.Fatal(err)
			return false
		}

		err = ioutil.WriteFile(newPath+"\\"+file, bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
			return false
		}
	} else {
		err := os.Rename(oldPath+"\\"+file, newPath+"\\"+file)

		if err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}

// Updates the file passed as path with the date provided as a parameter
func UpdateFileDate(filePath string, updateDate time.Time) {
	// Update the file's creation time
	err := os.Chtimes(filePath, updateDate, updateDate)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Updates the file passed as path with the unix timestamp provided as a parameter
func UpdateFileDateWithTimeStamp(filePath string, timestamp int64) {

	updateDate := time.Unix(timestamp, 0)

	// Update the file's creation time
	err := os.Chtimes(filePath, updateDate, updateDate)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func FileExists(filename string, path string) bool {
	if _, err := os.Stat(path + "\\" + filename); err == nil {
		return true
	} else {
		return false
	}
}
