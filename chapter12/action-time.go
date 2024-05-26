package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	appName := "HTTPChecker"
	action := "Basic"

	date := time.Now()

	logFileName := appName + "_" + action + "_" + strconv.Itoa(date.Year()) + "_" + date.Month().String() + "_" + strconv.Itoa(date.Day()) + ".log"
	fmt.Println("The name of the logfile is ", logFileName)

}
