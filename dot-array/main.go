package main

import "fmt"

func main() {
	var a = [...]int{
		1, 2, 3, 4, 5,
	}
	// a = append(a, 6) // this causes panic cause a is array not slice
	fmt.Println(a)
}
