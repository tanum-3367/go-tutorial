package main

import (
	"fmt"
	"strconv"
	"time"
)

func elapsedTime(start time.Time, end time.Time) string {
	elapsed := end.Sub(start)
	hours := strconv.Itoa(int(elapsed.Hours()))
	minutes := strconv.Itoa(int(elapsed.Minutes()))
	seconds := strconv.Itoa(int(elapsed.Seconds()))

	return "The total time was " + hours + " hours, " + minutes + " minutes, and " + seconds + " seconds."
}

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println(elapsedTime(start, end))
}
