package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}

	name := args[1]
	idontknow := args[0]
	greeting := "Hello, " + name + "!" + " I don't know who you are, " + idontknow
	fmt.Println(greeting)
}
