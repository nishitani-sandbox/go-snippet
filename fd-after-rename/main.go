package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	tmp, err := ioutil.TempFile(".", "_tmp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmp.Name())
	fmt.Println(tmp.Fd())
	os.Rename(tmp.Name(), "test")
	fmt.Println(tmp.Name())
	fmt.Println(tmp.Fd())
	tmp.Close()
	fmt.Println(tmp.Name())
	fmt.Println(tmp.Fd())

	f1, _ := os.Open("test")
	fmt.Println(f1.Fd())

	f2, _ := os.Open("test")
	fmt.Println(f2.Fd())

	f1.Close()
	f2.Close()

	os.Remove("test")
}
