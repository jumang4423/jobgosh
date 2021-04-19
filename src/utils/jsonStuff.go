package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	types "jobgosh/src/typesAndVars"
	"os"
)

func LoadGroupJson() []types.Group {

	// load json file
	bytes, err := ioutil.ReadFile(types.HomePath + types.GroupPath + "group.json")
	if err != nil {
		fmt.Println(string(types.ColorRed), "! error #04, group json not found", string(types.ColorReset))
		os.Exit(0)
	}
	var allGroups []types.Group
	if err := json.Unmarshal(bytes, &allGroups); err != nil {
		fmt.Println(string(types.ColorRed), "! error #04, group json not found", string(types.ColorReset))
		os.Exit(0)
	}

	return allGroups
}

func LoadLogJson(directoryName string) []types.Log {

	// load json file
	bytes, err := ioutil.ReadFile(types.HomePath + types.LogPath + directoryName + ".json")
	if err != nil {
		fmt.Println(string(types.ColorRed), "! error #04, log json not found", string(types.ColorReset))
		os.Exit(0)
	}
	var allLog []types.Log
	if err := json.Unmarshal(bytes, &allLog); err != nil {
		fmt.Println(string(types.ColorRed), "! error #04, log json not found", string(types.ColorReset))
		os.Exit(0)
	}

	return allLog
}