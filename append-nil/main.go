package main

import "fmt"

type a struct {
	f int
}

func main() {
	s := make([]*a, 0, 2)
	s = append(s, &a{1})
	s = append(s, nil)
	fmt.Println(s)
}
