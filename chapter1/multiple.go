package main

import (
	"fmt"
	"time"
)

func main() {
	Debug := true
	LogLevel := "info"
	StartUpTime := time.Now()

	fmt.Println(Debug, LogLevel, StartUpTime)

}
