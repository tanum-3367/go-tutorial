package main

func genSlices() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)
	return s1, s2, s3
}

func main() {
	s1, s2, s3 := genSlices()
	println(s1 == nil, len(s1), cap(s1))
	println(s2 == nil, len(s2), cap(s2))
	println(s3 == nil, len(s3), cap(s3))
}
