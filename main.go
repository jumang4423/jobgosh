package main

// import pkg
import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"time"
)

// variables
var (
	Version   string = "v0.1"
	Revision  string = "17/04/2021"
	logPath   string = "log/"
	groupPath string = "group/"
	homePath  string = "/go/src/work/" // usr.HomeDir

	colorReset string = "\033[0m"
	colorRed   string = "\033[31m"
	colorGreen string = "\033[32m"
	colorCyan  string = "\033[36m"
)

func main() {

	// disable when u develop
	usr, _ := user.Current()
	homePath = usr.HomeDir + "/.jobgosh/"

	// init flags
	times := flag.String("times", "", "see how long u spend times for each group")
	from := flag.String("from", "", "duration flag like YYYY/MM/DD")
	to := flag.String("to", "", "duration flag like YYYY/MM/DD")
	work := flag.String("work", "", "start work or finish work by up or work")

	// parse flags
	flag.Parse()

	// cmd selector
	if *times == "" && *work == "" && *to == "" && *from == "" {
		Help() // no command
	} else if *work != "" {

		if *work == "up" || *work == "down" {
			WorkTriggers(*work) // w command trrigered
		} else {
			fmt.Println(string(colorRed), "! ERROR #01, bad hook parameter", string(colorReset))
			os.Exit(0)
		}

	} else if *times != "" {
		if *times == "all" {
			loc, _ := time.LoadLocation("UTC")
			_now := time.Now().AddDate(0, 0, 1).In(loc).Unix()
			TimeViewer(0, _now) // t command trrigered
		} else {

			fmt.Println(string(colorRed), "! ERROR #02, bad hook parameter", string(colorReset))
			os.Exit(0)
		}
	} else if *from != "" && *to != "" {
		loc, _ := time.LoadLocation("UTC")
		_from, _ := time.Parse("2006/01/02", *from)
		_to, _ := time.Parse("2006/01/02", *to)
		_UnixedFrom := _from.In(loc).Unix()
		_UnixedTo := _to.In(loc).Unix()
		TimeViewer(_UnixedFrom, _UnixedTo) // t command trrigered
	}
}
