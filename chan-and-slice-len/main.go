package main

import "fmt"

func main() {
	s1 := make([]int, 50)
	fmt.Println(len(s1)) // 50

	s2 := make(chan int, 50)
	fmt.Println(len(s2)) // 0
	s2 <- 0
	fmt.Println(len(s2)) // 1
}
