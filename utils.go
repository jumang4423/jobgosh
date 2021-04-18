package main

// import pkg
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Group struct {
	Group     string   `json:"group"`
	WorkSpace []string `json:"workSpace"`
}

type Log struct {
	Up   string `json:"up"`
	Down string `json:"down"`
}

// seach files in the assigned directory and return an array of the files
func Dirsearch(dir string) []string {

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(string(colorRed), "! error #03, log directory error", string(colorReset))
		os.Exit(0)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirsearch(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

// file exist?
func IsExist(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

func LoadGroupJson() []Group {
	// load json file
	bytes, err := ioutil.ReadFile(homePath + groupPath + "group.json")
	if err != nil {
		fmt.Println(string(colorRed), "! error #04, group json not found", string(colorReset))
		os.Exit(0)
	}
	var allGroups []Group
	if err := json.Unmarshal(bytes, &allGroups); err != nil {
		fmt.Println(string(colorRed), "! error #04, group json not found", string(colorReset))
		os.Exit(0)
	}

	return allGroups
}

// save group json
func SaveGroupJson(_allGroups []Group) {
	str, _ := json.MarshalIndent(_allGroups, "", "   ")
	ioutil.WriteFile(homePath+groupPath+"group.json", str, 0644)
}

// save group json
func SaveLogJson(directoryName string) {
	var log []Log
	str, _ := json.MarshalIndent(log, "", "   ")
	ioutil.WriteFile(homePath+logPath+directoryName+".json", str, 0644)
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

// new group object into existing array
func SaveNewGroup(_newGroup string, directoryName string, _allGroups []Group) []Group {
	var _newG Group

	_newG.Group = _newGroup
	_allGroups = append(_allGroups, _newG)
	return _allGroups
}

// add group workspace
func AddAsGroupMember(_Group string, directoryName string, _allGroups []Group) {

	for i, v := range _allGroups {

		if v.Group == _Group {
			// add
			_allGroups[i].WorkSpace = append(v.WorkSpace, directoryName)
		}
	}

	SaveGroupJson(_allGroups)
}

// find group in group.json
func FindGroup(directoryName string) string {

	var group string = ""
	var _allGroups []Group
	_allGroups = LoadGroupJson()

	for _, v := range _allGroups {

		// find workspace
		for _, r := range v.WorkSpace {
			if r == directoryName {
				group = v.Group
			}
		}

	}

	return group
}

func PrintProgress(_group string, directoryName string, progress string) {
	fmt.Printf(string(colorGreen))
	fmt.Print("\n")
	fmt.Println("["+directoryName+"] ditails", string(colorReset))
	fmt.Println("> group : " + _group)
	fmt.Println("> progress : " + progress + "\n")
}
