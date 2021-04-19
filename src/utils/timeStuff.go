package utils

import (
	types "jobgosh/src/typesAndVars"
	"time"
)

func UnixTimeCalc(up string, down string) int64 {
	loc, _ := time.LoadLocation("UTC")
	_up, _ := time.Parse("2006/01/02 15:04:05", up)
	_down, _ := time.Parse("2006/01/02 15:04:05", down)
	_UnixedDown := _down.In(loc).Unix()
	_UnixedUp := _up.In(loc).Unix()

	result := _UnixedDown - _UnixedUp
	return result
}

func SpaceSumCalc(logJson []types.Log, from int64, to int64) int64 {
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

func TimeCalc(from int64, to int64, listOfData []types.Result) []types.Result {
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
	_, min, sec := t.Clock()

	hour := int((cAmount - int64(sec) - int64(min)*60) / 3600)

	return hour, min, sec
}