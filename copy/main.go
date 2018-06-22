package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	d := make([]int, 2, 4)
	copy(d, s)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
}
