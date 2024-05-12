package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)

	countTemp := 5
	// use & to create a pointer from existing variable
	count3 := &countTemp

	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("t: %#v\n", t)

}
