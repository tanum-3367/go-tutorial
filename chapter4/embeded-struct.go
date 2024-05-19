package main

type name string
type location struct {
	x int
	y int
}
type size struct {
	width  int
	height int
}

type dot struct {
	name
	location
	size
}

func getDots() []dot {
	var dot1 dot
	dot2 := dot{}
	dot2.name = "dot2"
	dot2.x = 5
	dot2.y = 5
	dot2.width = 10
	dot2.height = 10

	dot3 := dot{
		name: "dot3",
		location: location{
			x: 10,
			y: 10,
		},
		size: size{
			width:  10,
			height: 10,
		},
	}

	dot4 := dot{}
	dot4.name = "dot4"
	dot4.location.x = 15
	dot4.location.y = 15
	dot4.size.width = 10
	dot4.size.height = 10

	return []dot{dot1, dot2, dot3, dot4}
}

func main() {
	dots := getDots()
	for _, dot := range dots {
		println(dot.name)
		println(dot.x)
		println(dot.y)
		println(dot.width)
		println(dot.height)
		println("------")
	}
}
