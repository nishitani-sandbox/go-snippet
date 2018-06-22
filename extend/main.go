package main

import "fmt"

type Base struct {
	s string
}

func (b *Base) Overrided() {
	fmt.Println("this is overrided")
}

func (b *Base) NotOverrided() {
	fmt.Println("this is not overrided")
}

type Sub struct {
	*Base
	t string
}

func (s *Sub) Overrided() {
	fmt.Println("successfully overrided")
}

func main() {
	s := Sub{
		&Base{"s"}, "t", // cannot mix field:value and value initializers
	}
	s.Overrided()    // successfully overrided
	s.NotOverrided() // this is not overried
}
