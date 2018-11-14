package main

import (
	"fmt"
	"reflect"
)

type Status string

const (
	ToDo Status = "TODO"
)

func isString(value interface{}) bool {
	switch value.(type) {
	case string:
		return true
	}
	return false
}

func likeString(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return true
	}
	return false
}

func main() {
	if isString(ToDo) {
		fmt.Println("string!!")
	}

	if likeString(ToDo) {
		fmt.Println("like string!!")
	}
}
