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
	t := flag.String("t", "", "same as -times")
	from := flag.String("from", "", "duration flag like YYYY/MM/DD")
	to := flag.String("to", "", "duration flag like YYYY/MM/DD")
	work := flag.String("work", "", "start work or finish work by up or work")
	w := flag.String("w", "", "same as -work")

	// parse flags
	flag.Parse()

	// factors
	// commands
	var _times string = *times + *t
	var _to string = *to
	var _from string = *from
	var _work string = *work + *w

	if _work == "u" {
		_work = "up"
	}
	if _work == "d" {
		_work = "down"
	}
	if _times == "a" {
		_times = "all"
	}

	// second command filter

	// cmd selector
	if _times == "" && _work == "" && _to == "" && _from == "" {
		Help() // no command
	} else if _work != "" {

		if _work == "up" ||
			_work == "down" {
			wTriggers.WorkTriggers(_work) // w command trrigered
		} else {
			fmt.Println(string(types.ColorRed), "! ERROR #01, bad hook parameter", string(types.ColorReset))
			os.Exit(0)
		}

	} else if _times != "" {
		if _times == "all" {
			loc, _ := time.LoadLocation("UTC")
			_now := time.Now().AddDate(0, 0, 1).In(loc).Unix()
			tv.TimeViewer(0, _now) // t command trrigered
		} else {
			fmt.Println(string(types.ColorRed), "! ERROR #02, bad hook parameter", string(types.ColorReset))
			os.Exit(0)
		}
	} else if _from != "" && _to != "" {
		loc, _ := time.LoadLocation("UTC")
		_from, _ := time.Parse("2006/01/02", _from)
		_to, _ := time.Parse("2006/01/02", _to)
		_UnixedFrom := _from.In(loc).Unix()
		_UnixedTo := _to.In(loc).Unix()
		tv.TimeViewer(_UnixedFrom, _UnixedTo) // t command trrigered
	}
}
