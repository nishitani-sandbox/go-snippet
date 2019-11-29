package main

import "fmt"

type State int

const (
	Uploading State = iota
	Converting
)

func printInt(i int) {
	fmt.Println(i)
}

func main() {
	printInt(Uploading)
	printInt(Converting)
}
