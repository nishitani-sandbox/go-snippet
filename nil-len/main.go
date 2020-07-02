package main

import "fmt"

func nilSlice() []int {
	return nil
}

func main() {
	fmt.Println(len(nilSlice()))
}
