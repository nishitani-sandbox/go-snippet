package main

import "fmt"

var (
	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func sendVal(i int, ch chan<- int) {
	ch <- i
}

func main() {
	ch := make(chan int, 5)

	for _, i := range s {
		go sendVal(i, ch)
	}

	for i := 0; i < len(s); i++ {
		select {
		case j := <-ch:
			fmt.Println(j)
		}
	}
}
