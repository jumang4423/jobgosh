package main

// import pkg
import (
	"fmt"
)

func Help() {
	// say welcome when nothing entered except jobgosh command
	fmt.Println("\n> jobgosh@" + Version)

	fmt.Println(string(colorGreen), "\n Usage: jobgosh -command\n", string(colorReset))
	fmt.Println("	jobgosh -t [period of time]     see how long u spend times for each group, [period of time] should be month, year or all")
	fmt.Println("	jobgosh -w [up | down]          up to start work and down to finish work\n")

	fmt.Println(string(colorGreen), "All commands:\n", string(colorReset))
	fmt.Println("	-t, -w\n")
}
