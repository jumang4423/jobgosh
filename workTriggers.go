package main

// import pkg
import (
	"os"
	"path/filepath"
)

func WorkTriggers(w string) {

	type Group struct {
		Group     string   `json:"group"`
		WorkSpace []string `json:"workSpace"`
	}

	type Log struct {
		Up   string `json:"up"`
		Down string `json:"down"`
	}

	// get a current directory
	_currentDir, _ := os.Getwd()
	directoryName := filepath.Base(_currentDir)

	var _group string = ""

	// if directoryName isnt there in group.json, make new group
	// which means that, we dont watch log folder at this point.
	_group = FindGroup(directoryName)
	if _group == "" {
		_group = MakeNewGroup() // make a new group
	}

	PrintProgress(_group, directoryName, "up")

}
