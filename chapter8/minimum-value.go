package main

import "fmt"

func findMinGeneric[Num int | float64](list []Num) Num {
	if len(list) == 0 {
		return 0
	}
	min := list[0]
	for _, v := range list {
		if v < min {
			v = min
		}
	}
	return min
}

func main() {
	min := findMinGeneric([]int{, 2, 3, 4, 5})
	fmt.Printf("Minimum value is %d\n", min)
}
