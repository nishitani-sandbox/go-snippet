package main

import "fmt"

func main() {
	s1 := make([]int, 50)
	s2 := make(chan int, 50)

	fmt.Println(len(s1))
	fmt.Println(len(s2))
}
