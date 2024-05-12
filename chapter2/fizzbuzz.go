package main

import "fmt"

const FIZZ_BUZZ_LIMIT = 100

func main() {
	for i := 1; i <= FIZZ_BUZZ_LIMIT; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
