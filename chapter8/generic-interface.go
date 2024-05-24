package main

import "fmt"

type Number interface {
	int | float64
}

func findMaxGeneric[Num Number](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	maxGenericInt := findMaxGeneric([]int{1, 2, 3, 4, 5})
	maxGenericFloat := findMaxGeneric([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	fmt.Printf("Maximum value of int is %d\n", maxGenericInt)
	fmt.Printf("Maximum value of float64 is %f\n", maxGenericFloat)
}
