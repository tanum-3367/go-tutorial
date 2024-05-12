package main

import "fmt"

func main() {
	numbers := []int{5, 3, 8, 2, 1, 4, 2, 23, 2, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Before: ", numbers)

	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(numbers); i++ {
			if numbers[i-1] > numbers[i] {
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
				swapped = true
			}
		}
	}

	fmt.Println("After: ", numbers)
}
