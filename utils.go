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

	fmt.Println(homePath + groupPath + "group.json")

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
	loc, _ := time.LoadLocation("UTC")
	now := time.Now()
	t := time.Unix(cAmount, 0).In(loc)
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
	loc, _ := time.LoadLocation("UTC")
	_up, _ := time.Parse("2006/01/02 15:04:05", up)
	_down, _ := time.Parse("2006/01/02 15:04:05", down)
	_UnixedDown := _down.In(loc).Unix()
	_UnixedUp := _up.In(loc).Unix()

	result := _UnixedDown - _UnixedUp
	return result
}

func SpaceSumCalc(logJson []Log, from int64, to int64) int64 {
	loc, _ := time.LoadLocation("UTC")
	var sum int64 = 0

	// one space
	for _, v := range logJson {
		_up, _ := time.Parse("2006/01/02 15:04:05", v.Up)
		_UnixedUp := _up.In(loc).Unix()
		if _UnixedUp > from && _UnixedUp < to && v.IsVisible {
			sum += UnixTimeCalc(v.Up, v.Down)
		}
	}
	return sum
}

func TimeCalc(from int64, to int64, listOfData []Result) []Result {
	for i, v := range listOfData {

		var spaceSum int64 = 0
		for j, r := range v.SumGroup {
			logJson := LoadLogJson(r.WorkSpace)
			// result sum of workSpace time
			var workSpaceSum int64 = SpaceSumCalc(logJson, from, to)
			listOfData[i].SumGroup[j].Sum = workSpaceSum
			spaceSum += workSpaceSum
		}

		listOfData[i].Sum = spaceSum
	}

	return listOfData
}

func UnixToHMS(cAmount int64) (int, int, int) {
	loc, _ := time.LoadLocation("UTC")
	t := time.Unix(cAmount, 0).In(loc)
	hour, min, sec := t.Clock()

	return hour, min, sec
}
