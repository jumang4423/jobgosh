package main

// import pkg
import (
	"fmt"
)

func Help() {
	// say welcome when nothing entered except jobgosh command
	fmt.Println("\n> jobgosh@" + Version)

	fmt.Println(string(colorGreen), "\n Usage: jobgosh -command\n", string(colorReset))
	fmt.Println("	jobgosh -times all                              see how long u spend times for each group")
	fmt.Println("	jobgosh -from [YYYY/MM/DD] -to [YYYY/MM/DD]     specify duration")
	fmt.Println("	jobgosh -work [up | down]                       up to start work and down to finish\n")

	fmt.Println(string(colorGreen), "All commands:\n", string(colorReset))
	fmt.Println("	-times, -from, -to, -w\n")
}
