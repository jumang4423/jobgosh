package main

// import pkg
import (
	"fmt"
	types "jobgosh/src/typesAndVars"
)

func Help() {
	// say welcome when nothing entered except jobgosh command
	fmt.Println("\n> jobgosh@" + types.Version)

	fmt.Println(string(types.ColorGreen), "\n Usage: jobgosh -command\n", string(types.ColorReset))
	fmt.Println("	jobgosh -times all                              see how long u spend times for each group")
	fmt.Println("	jobgosh -from [YYYY/MM/DD] -to [YYYY/MM/DD]     specify duration")
	fmt.Printf("	jobgosh -work [up | down]                       up to start work and down to finish\n\n")

	fmt.Println(string(types.ColorGreen), "All commands:\n", string(types.ColorReset))
	fmt.Printf("	-times, -from, -to, -w\n\n")
}
