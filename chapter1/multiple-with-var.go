package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	var start, middle, end float32
	fmt.Println(start, middle, end)
	var name, left, right, top, bottom = "hello", 1, 2, 3, 4
	fmt.Println(name, left, right, top, bottom)

	var Debug, LogLevel, StartUpTime = getConfig()
	fmt.Println(Debug, LogLevel, StartUpTime)
}
