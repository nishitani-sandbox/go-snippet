package main

import "fmt"

func numSlices() (s []int) {
	s = append(s, 1, 2)
	return
}

func nilSlices() []int {
	return nil
}

func main() {
	fmt.Println(numSlices())

	s := nilSlices()
	fmt.Println(s)
	if s != nil {
		fmt.Println("not nil")
	}

	s = append(s, 1)
	fmt.Println(s)

	s = make([]int, 0)
	fmt.Println(s)
	if s != nil {
		fmt.Println("not nil")
	}
}
