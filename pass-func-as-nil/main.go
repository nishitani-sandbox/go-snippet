package main

import "fmt"

func test(f func()) {
	if f != nil {
		f()
	}
}

func main() {
	test(func() { fmt.Println("hoge") })
	test(nil)
}
