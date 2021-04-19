package main

// import pkg
import (
	"flag"
	"fmt"
	tv "jobgosh/src/timeViewer"
	types "jobgosh/src/typesAndVars"
	wTriggers "jobgosh/src/workTriggers"
	"os"
	"os/user"
	"time"
)



func main() {
	
	// disable when u develop
	if !types.DOCKER_DEVELOPMENT {
		usr, _ := user.Current()
		types.HomePath = usr.HomeDir + "/.jobgosh/"
	}

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
			wTriggers.WorkTriggers(*work) // w command trrigered
		} else {
			fmt.Println(string(types.ColorRed), "! ERROR #01, bad hook parameter", string(types.ColorReset))
			os.Exit(0)
		}

	} else if *times != "" {
		if *times == "all" {
			loc, _ := time.LoadLocation("UTC")
			_now := time.Now().AddDate(0, 0, 1).In(loc).Unix()
			tv.TimeViewer(0, _now) // t command trrigered
		} else {

			fmt.Println(string(types.ColorRed), "! ERROR #02, bad hook parameter", string(types.ColorReset))
			os.Exit(0)
		}
	} else if *from != "" && *to != "" {
		loc, _ := time.LoadLocation("UTC")
		_from, _ := time.Parse("2006/01/02", *from)
		_to, _ := time.Parse("2006/01/02", *to)
		_UnixedFrom := _from.In(loc).Unix()
		_UnixedTo := _to.In(loc).Unix()
		tv.TimeViewer(_UnixedFrom, _UnixedTo) // t command trrigered
	}
}
