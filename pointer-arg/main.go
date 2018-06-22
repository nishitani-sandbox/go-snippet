package main

import "fmt"

func addOne(s *[]int) {
	*s = append(*s, 1)
}

func replaceByOne(i *int) {
	*i = 1
}

func main() {
	s := []int{0, 1, 2, 3}
	addOne(&s)
	fmt.Println(s)

	i := 0
	replaceByOne(&i)
	fmt.Println(i)
}
