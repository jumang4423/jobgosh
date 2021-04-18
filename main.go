package main

// import pkg
import (
	"flag"
	"fmt"
	"os"
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

	// init flags
	t := flag.String("t", "", "see how long u spend times for each group by month, year or all")
	w := flag.String("w", "", "start work or finish work by up or work")

	// parse flags
	flag.Parse()

	// cmd selector
	if *t == "" && *w == "" {
		Help() // no command
	} else if *w != "" {

		if *w == "up" || *w == "down" {
			WorkTriggers(*w) // w command trrigered
		} else {
			fmt.Println(string(colorRed), "! ERROR #01, bad hook parameter", string(colorReset))
			os.Exit(0)
		}

	} else if *t != "" {
		if *w == "month" || *w == "year" || *w == "all" {
			// TimeViewer(*t) // t command trrigered
		} else {
			fmt.Println(string(colorRed), "! ERROR #02, bad hook parameter", string(colorReset))
			os.Exit(0)
		}
	}
}
