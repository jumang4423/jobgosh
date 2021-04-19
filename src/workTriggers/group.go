package WorkTriggers

import (
	"fmt"
	types "jobgosh/src/typesAndVars"
	"jobgosh/src/utils"
	"os"
	"path/filepath"
	"strconv"
)

// new group maker
func MakeNewGroup() string {

	// get a current directory
	_currentDir, _ := os.Getwd()
	directoryName := filepath.Base(_currentDir)

	// non-exist
	fmt.Print("\n")
	fmt.Println("> seems like [" + directoryName + "] isnt organised by any groups currently")
	fmt.Print("\n")
	fmt.Println(string(types.ColorCyan), "? select a group:\n", string(types.ColorReset))

	allGroups := utils.LoadGroupJson()

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

	switch _group {
	case "":
		fmt.Println("no group entered, current process will be no longer exist")
		os.Exit(0)
	
		

	case "*":
		fmt.Print("\nenter a new group name :")

		//input number
		var _input string

		fmt.Scan(&_input)
		if _group == "" {
			fmt.Println("no group entered, current process will be no longer exist")
			os.Exit(0)
		}

		// add new group into the json
		utils.AddAsGroupMember(_input, directoryName, utils.SaveNewGroup(_input, directoryName, allGroups))
		utils.SaveLogJson(directoryName)
		return _input
		

	default:
		utils.AddAsGroupMember(_group, directoryName, allGroups)
		utils.SaveLogJson(directoryName)
		return _group
		
	}
	return _group
}
