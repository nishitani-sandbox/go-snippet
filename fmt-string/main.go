package main

import "fmt"

type MyInt int

func (i MyInt) String() string {
	return "myint"
}

func main() {
	fmt.Printf("%s\n", MyInt(1))
}
