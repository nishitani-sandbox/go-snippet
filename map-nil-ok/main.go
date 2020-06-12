package main

import "fmt"

func main() {
	s := "foo"
	m := map[string]*string{
		"hoge": &s,
		"foo":  nil,
	}

	if v, ok := m["hoge"]; ok {
		fmt.Println(v)
	}
	if v, ok := m["foo"]; ok {
		fmt.Println(v)
	}
}
