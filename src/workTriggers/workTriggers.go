package WorkTriggers

// import pkg
import (
	"jobgosh/src/utils"
	"os"
	"path/filepath"
)

func WorkTriggers(w string) {



	// get a current directory
	_currentDir, _ := os.Getwd()
	directoryName := filepath.Base(_currentDir)

	var _group string = ""

	// if directoryName isnt there in group.json, make new group
	// which means that, we dont watch log folder at this point.
	_group = utils.FindGroup(directoryName)
	if _group == "" {
		_group = MakeNewGroup() // make a new group
	}

	// show status and save timestamp
	utils.PrintProgress(_group, directoryName, w, SaveTimeStamp(directoryName, w))
}


