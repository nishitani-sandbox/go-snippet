package main

import "fmt"

func main() {
	i := 1

	f := func() {
		fmt.Printf("before num: %d\n", i)

		i := i + 1
		fmt.Printf("after num: %d\n", i)
	}

	f()
	f()
}
