package main

import (
	"os"
	"strconv"
)

type structure struct {
	name     string
	isRoot   bool
	fileInfo os.FileInfo
	children []structure
	parent   *structure
}

func createDummyStructure() structure {
	mainStructure := structure{
		name:     "2023",
		children: []structure{},
	}

	for i := 1; i <= 12; i++ {
		mainStructure.children = append(mainStructure.children, structure{name: strconv.Itoa(i)})
	}
	return mainStructure
}

func (s structure) makeYearFolderStructure(year string) {
	//root := s.getRoot()

}

func (s structure) makeMonthFolderStructure(month string) {

}

func (s structure) getRoot() structure {
	if s.isRoot {
		return s
	}

	tempStructure := *s.parent

	for tempStructure.isRoot {
		if tempStructure.isRoot {
			return tempStructure
		}
		tempStructure = *tempStructure.parent
	}

	return tempStructure
}
