package utils

import (
	"encoding/json"
	"io/ioutil"
	types "jobgosh/src/typesAndVars"
)

// save group json
func SaveGroupJson(_allGroups []types.Group) {
	str, _ := json.MarshalIndent(_allGroups, "", "   ")
	ioutil.WriteFile(types.HomePath+types.GroupPath+"group.json", str, 0644)
}

// save group json
func SaveLogJson(directoryName string) {
	var log []types.Log
	str, _ := json.MarshalIndent(log, "", "   ")
	ioutil.WriteFile(types.HomePath+types.LogPath+directoryName+".json", str, 0644)
}

// new group object into existing array
func SaveNewGroup(_newGroup string, directoryName string, _allGroups []types.Group) []types.Group {
	var _newG types.Group

	_newG.Group = _newGroup
	_allGroups = append(_allGroups, _newG)
	return _allGroups
}