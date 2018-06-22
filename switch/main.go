package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("error 1")
var err2 = errors.New("error 2")

func printed() {
	fmt.Println("PRINTED CASE")
	e := err1
	switch e {
	case nil:
		fmt.Println("nil")
	case err1:
		fmt.Println(e)
	case err2:
	default:
		fmt.Println(e)
	}
}

func notPrinted() {
	fmt.Println("NOT PRINTED CASE")
	e := err2
	switch e {
	case nil:
		fmt.Println("nil")
	case err1:
		fmt.Println(e)
	case err2:
	default:
		fmt.Println(e)
	}
}

func main() {
	printed()
	notPrinted()
}
