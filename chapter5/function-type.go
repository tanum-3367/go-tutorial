package main

import "fmt"

type calc func(int, int) string

func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}

func add(i, j int) string {
	result := i + j
	return fmt.Sprintf("Added %d + %d = %d", i, j, result)
}

func substract(i, j int) string {
	result := i - j
	return fmt.Sprintf("Substracted %d - %d = %d", i, j, result)
}

func main() {
	calculator(add, 5, 10)
	calculator(substract, 10, 5)
}
