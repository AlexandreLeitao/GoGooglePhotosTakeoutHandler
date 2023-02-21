package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type takeoutJson struct {
	Title                 string   `json:"title"`
	CreationTime          jsonTime `json:"creationTime"`
	PhotoTakenTime        jsonTime `json:"photoTakenTime"`
	PhotoLastModifiedTime jsonTime `json:"photoLastModifiedTime"`
}

type jsonTime struct {
	Timestamp string `json:"timestamp"`
	Formatted string `json:"formatted"`
}

func GetFileJson(filepath string) takeoutJson {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(-1)
	}

	// Unmarshal the JSON data into a struct
	var returnJson takeoutJson
	err = json.Unmarshal(file, &returnJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(-2)
	}
	return returnJson
}
