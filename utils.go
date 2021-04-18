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

type Group struct {
	Group     string   `json:"group"`
	WorkSpace []string `json:"workSpace"`
}

type Log struct {
	Up        string `json:"up"`
	Down      string `json:"down"`
	IsVisible bool   `json:"isVisible"`
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

func LoadLogJson(directoryName string) []Log {
	// load json file
	bytes, err := ioutil.ReadFile(homePath + logPath + directoryName + ".json")
	if err != nil {
		fmt.Println(string(colorRed), "! error #04, log json not found", string(colorReset))
		os.Exit(0)
	}
	var allLog []Log
	if err := json.Unmarshal(bytes, &allLog); err != nil {
		fmt.Println(string(colorRed), "! error #04, log json not found", string(colorReset))
		os.Exit(0)
	}

	return allLog
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

// Up Down display thingy
func PrintProgress(_group string, directoryName string, progress string, cAmount int64) {
	now := time.Now()
	t := time.Unix(cAmount, 0)
	hour, min, sec := t.Clock()

	fmt.Printf(string(colorGreen))
	fmt.Print("\n")
	fmt.Println("["+directoryName+" ditails]", string(colorReset))
	fmt.Println("> group : " + _group)
	fmt.Println("> progress : " + progress + "")
	fmt.Printf(string(colorGreen))
	fmt.Print("\n", string(colorCyan))
	if progress == "up" {
		fmt.Println("> timestamped at " + now.Format("2006/01/02 15:04:05") + "\n")
	} else if progress == "down" {
		fmt.Print("> timestamped at " + now.Format("2006/01/02 15:04:05"))
		fmt.Print("\n")
		fmt.Println("> " + strconv.Itoa(hour) + "h " + strconv.Itoa(min) + "m " + strconv.Itoa(sec) + "s added\n")
	}

}

func UnixTimeCalc(up string, down string) int64 {
	_up, _ := time.Parse("2006/01/02 15:04:05", up)
	_down, _ := time.Parse("2006/01/02 15:04:05", down)
	_UnixedDown := _up.Unix()
	_UnixedUp := _down.Unix()

	result := _UnixedUp - _UnixedDown
	fmt.Println(_UnixedDown)
	fmt.Println(_UnixedUp)

	return result

}
