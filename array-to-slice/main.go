package main

import "fmt"

func main() {
	var a [4]int
	src := []int{1, 2, 3, 4}
	dst := a[:] // convert array to slice
	copy(dst, src)
	fmt.Println(a, dst) // [1 2 3 4] [1 2 3 4]

	src = []int{5, 6, 7, 8}
	dst = a[:]
	dst = append(dst, 9, 10) // a underlying array of `dst` changes
	copy(dst, src)
	fmt.Println(a, dst) // [1 2 3 4] [5 6 7 8 9 10]

	/*
		CONCLUSION: `a[:]` generates a slice that has `a` as a underlying array
	*/
}
