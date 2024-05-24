package main

import "fmt"

func findLargestRanchInStock[K comparable, V int | float64](ranches map[K]V) K {
	var stock V
	var name K

	for k, v := range ranches {
		if v > stock {
			stock = v
			name = k
		}
	}
	return name
}

func main() {
	animalStock := map[string]int{
		"Chicken": 100,
		"Cow":     50,
		"Goat":    75,
	}
	miscStock := map[string]float64{
		"Hay":   100.5,
		"Grass": 50.25,
		"Water": 75.75,
	}
	largestStockInt := findLargestRanchInStock(animalStock)
	largestStockFloat := findLargestRanchInStock(miscStock)
	fmt.Printf("Largest stock of animals is %s\n", largestStockInt)
	fmt.Printf("Largest stock of misc items is %s\n", largestStockFloat)
}
