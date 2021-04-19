package timeViewer

import (
	"fmt"
	types "jobgosh/src/typesAndVars"
	"jobgosh/src/utils"
	"strconv"
)

// import pkg

func TimeViewer(from int64, to int64) {

	var listOfData []types.Result
	// get a group Json
	groupJson := utils.LoadGroupJson()
	for _, v := range groupJson {
		// result time of each group
		var _indexedResult types.Result
		_indexedResult.GroupName = v.Group

		for _, p := range v.WorkSpace {
			// result time of each workSpace
			
			var _current types.SpaceNdCalc

			_current.WorkSpace = p
			_indexedResult.SumGroup = append(_indexedResult.SumGroup, _current)
		}

		listOfData = append(listOfData, _indexedResult)
	}

	listOfData = utils.TimeCalc(from, to, listOfData)

	PrintSumResult(listOfData)

	// _currentDir, _ := os.Getwd()
	// directoryName := filepath.Base(_currentDir)
	// logJson := LoadLogJson(directoryName)

}

// print calculated datas
func PrintSumResult(listOfData []types.Result) {

	// all time calclation
	var _allTime int64
	for _, v := range listOfData {
		_allTime += v.Sum
	}
	_, min, sec := utils.UnixToHMS(_allTime)
	hour := int((_allTime - int64(sec) - int64(min)*60) / 3600)

	fmt.Printf(string(types.ColorGreen))
	fmt.Print("\n")
	fmt.Print("[*]", string(types.ColorReset))
	fmt.Println(" " + strconv.Itoa(hour) + "h " + strconv.Itoa(min) + "m " + strconv.Itoa(sec) + "s\n")

	for _, v := range listOfData {

		fmt.Printf(string(types.ColorGreen))
		fmt.Print("["+v.GroupName+"]", string(types.ColorReset))
		_, min_g, sec_g := utils.UnixToHMS(v.Sum)
		hour_g := int((v.Sum - int64(sec) - int64(min)*60) / 3600)
		fmt.Println(" " + strconv.Itoa(hour_g) + "h " + strconv.Itoa(min_g) + "m " + strconv.Itoa(sec_g) + "s")

		for _, r := range v.SumGroup {
			_, min_sps, sec_sps := utils.UnixToHMS(r.Sum)
			hour_sps := int((r.Sum - int64(sec) - int64(min)*60) / 3600)
			fmt.Println("> " + r.WorkSpace + " : " + strconv.Itoa(hour_sps) + "h " + strconv.Itoa(min_sps) + "m " + strconv.Itoa(sec_sps) + "s")
		}
		fmt.Print("\n")
	}

}

// UnixToHMS
