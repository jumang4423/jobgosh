package WorkTriggers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	types "jobgosh/src/typesAndVars"
	"jobgosh/src/utils"
	"os"
	"time"
)

func SaveTimeStamp(directoryName string, progress string) int64 {
	var allLog []types.Log
	var cAmount int64 = -1 // 出来高
	allLog = utils.LoadLogJson(directoryName)

	if progress == "up" {
		var _addLog types.Log
		_addLog.Up = time.Now().Format("2006/01/02 15:04:05")
		_addLog.IsVisible = false
		allLog = append(allLog, _addLog)

		str, _ := json.MarshalIndent(allLog, "", "   ")
		ioutil.WriteFile(types.HomePath+types.LogPath+directoryName+".json", str, 0644)
	} else if progress == "down" {

		if len(allLog) == 0 {
			os.Exit(0)
		}
		_now := time.Now().Format("2006/01/02 15:04:05")
		// if down isnt already declared
		if allLog[len(allLog)-1].Down == "" {
			allLog[len(allLog)-1].Down = _now
			allLog[len(allLog)-1].IsVisible = true
			str, _ := json.MarshalIndent(allLog, "", "   ")
			ioutil.WriteFile(types.HomePath+types.LogPath+directoryName+".json", str, 0644)

			cAmount = utils.UnixTimeCalc(allLog[len(allLog)-1].Up, allLog[len(allLog)-1].Down)
		} else {
			fmt.Println("\n> down is already timestamped\n")
			os.Exit(0)
		}
	}

	return cAmount
}
