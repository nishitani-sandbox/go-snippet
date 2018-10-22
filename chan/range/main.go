package main

import "fmt"

var (
	ints = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func sendVal(s []int, ch chan<- int) {
	for _, i := range s {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go sendVal(ints, ch)

	for i := range ch {
		fmt.Println(i)
	}
}
