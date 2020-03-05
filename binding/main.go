package main

import "fmt"

func main() {
	ch := make(chan struct{}, 10)

	fmt.Println("======")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("without binding: %d\n", i)
			ch <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-ch
	}

	fmt.Println("======")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("with binding: %d\n", i)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
