package main

import "log"

type MyInt int

func main() {
	a := MyInt(1)
	b := MyInt(1)
	if a == b {
		log.Println("Equal")
	}
}
