package main

import "practice/pkg/shape"

func main() {
	t := shape.Triangle{Base: 10, Height: 5}
	r := shape.Rectangle{Length: 10, Width: 5}
	s := shape.Square{Side: 5}

	shape.PrintShapeDetauls(t, r, s)
}
