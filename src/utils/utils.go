package utils

// import pkg
import (
	"fmt"

	types "jobgosh/src/typesAndVars"
	"os"
	"strconv"
	"time"
)

// file exist?
func IsExist(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}



// Up Down display thingy
func PrintProgress(_group string, directoryName string, progress string, cAmount int64) {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now()
	t := time.Unix(cAmount, 0).In(loc)
	_, min, sec := t.Clock()

	hour := int((cAmount - int64(sec) - int64(min)*60) / 3600)
	
	fmt.Printf(string(types.ColorGreen))
	fmt.Print("\n")
	fmt.Println("["+directoryName+" ditails]", string(types.ColorReset))
	fmt.Println("> group : " + _group)
	fmt.Println("> progress : " + progress + "")
	fmt.Printf(string(types.ColorGreen))
	fmt.Print("\n", string(types.ColorCyan))

	if progress == "up" {
		fmt.Println("> timestamped at " + now.Format("2006/01/02 15:04:05") + "\n")
	} else if progress == "down" {
		fmt.Print("> timestamped at " + now.Format("2006/01/02 15:04:05"))
		fmt.Print("\n")
		fmt.Println("> " + strconv.Itoa(hour) + "h " + strconv.Itoa(min) + "m " + strconv.Itoa(sec) + "s added\n")
	}

}


