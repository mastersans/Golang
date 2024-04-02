package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("using the time in golang")
	currtime := time.Now()
	fmt.Println(currtime.Format("Monday 15:04 02-01-2006"))
	madetime := time.Date(2001, time.October, 31, 05, 30, 00, 00, time.UTC)
	fmt.Println(madetime.Format("Monday 15:04 02-01-2006"))

	duration := currtime.Sub(madetime)
	fmt.Println(duration.Hours(), "hours")
	_, _ = fmt.Scanln()
}
