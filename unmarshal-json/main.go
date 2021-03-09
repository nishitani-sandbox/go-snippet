package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

var (
	IDPat = regexp.MustCompile("^n[0-9]+$")
)

type ID string

func (i *ID) UnmarshalText(data []byte) error {
	if !IDPat.Match(data) {
		return errors.New("invalid id")
	}
	*i = ID(data)
	return nil
}

type User struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

func main() {
	normal := []byte(`"n3"`)
	invalid := []byte(`"invalid"`)
	var i1, i2 ID
	if err := json.Unmarshal(normal, &i1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i1)
	}
	if err := json.Unmarshal(invalid, &i2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i2)
	}

	normal = []byte(`{"id": "n3", "name": "test"}`)
	invalid = []byte(`{"name": "test"}`)
	var u1, u2 User
	if err := json.Unmarshal(normal, &u1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u1)
	}
	if err := json.Unmarshal(invalid, &u2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u2)
	}
}
