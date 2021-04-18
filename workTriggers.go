package main

// import pkg
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func WorkTriggers(w string) {

	type Group struct {
		Group     string   `json:"group"`
		WorkSpace []string `json:"workSpace"`
	}

	type Log struct {
		Up        string `json:"up"`
		Down      string `json:"down"`
		IsVisible bool   `json:"isVisible"`
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

	// show status and save timestamp
	PrintProgress(_group, directoryName, w, SaveTimeStamp(directoryName, w))
}

func SaveTimeStamp(directoryName string, progress string) int64 {
	var allLog []Log
	var cAmount int64 = -1 // 出来高
	allLog = LoadLogJson(directoryName)

	if progress == "up" {
		var _addLog Log
		_addLog.Up = time.Now().Format("2006/01/02 15:04:05")
		_addLog.IsVisible = false
		allLog = append(allLog, _addLog)

		str, _ := json.MarshalIndent(allLog, "", "   ")
		ioutil.WriteFile(homePath+logPath+directoryName+".json", str, 0644)
	} else if progress == "down" {
		_now := time.Now().Format("2006/01/02 15:04:05")
		allLog[len(allLog)-1].Down = _now
		allLog[len(allLog)-1].IsVisible = true
		str, _ := json.MarshalIndent(allLog, "", "   ")
		ioutil.WriteFile(homePath+logPath+directoryName+".json", str, 0644)

		cAmount = UnixTimeCalc(allLog[len(allLog)-1].Up, allLog[len(allLog)-1].Down)
	}

	return cAmount
}

// new group maker
func MakeNewGroup() string {

	// get a current directory
	_currentDir, _ := os.Getwd()
	directoryName := filepath.Base(_currentDir)

	// non-exist
	fmt.Print("\n")
	fmt.Println("> seems like [" + directoryName + "] isnt organised by any groups currently")
	fmt.Print("\n")
	fmt.Println(string(colorCyan), "? select a group:\n", string(colorReset))

	allGroups := LoadGroupJson()

	// list the groups
	for i, p := range allGroups {
		fmt.Printf("[%d] : %s\n", i+1, p.Group)
	}
	fmt.Println("[*] : add a new group")

	fmt.Print("\nenter the number :")

	//input number
	var _input string
	var _group string = ""

	fmt.Scan(&_input)
	// int input into group name
	for i, v := range allGroups {
		_string_to_int, _ := strconv.Atoi(_input)
		if i+1 == _string_to_int {
			_group = v.Group
			break
		}
	}
	if "*" == _input {
		_group = "*"
	}

	if _group == "" {
		fmt.Println("no group entered, current process will be no longer exist")
		os.Exit(0)
	}

	if _group == "*" {
		fmt.Print("\nenter a new group name :")

		//input number
		var _input string

		fmt.Scan(&_input)
		if _group == "" {
			fmt.Println("no group entered, current process will be no longer exist")
			os.Exit(0)
		}

		// add new group into the json
		AddAsGroupMember(_input, directoryName, SaveNewGroup(_input, directoryName, allGroups))
		SaveLogJson(directoryName)
		return _input

	} else {
		AddAsGroupMember(_group, directoryName, allGroups)
		SaveLogJson(directoryName)
		return _group
	}
}
