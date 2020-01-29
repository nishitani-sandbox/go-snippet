package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const (
	format = `"2006-01-02 15:04(MST)"`
)

type MyTime struct {
	t *time.Time
}

func (mt *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == `"--"` {
		*mt = MyTime{}
		return nil
	}

	t, err := time.Parse(format, string(data))
	if err != nil {
		return err
	}
	*mt = MyTime{&t}
	return nil
}

func (mt *MyTime) MarshalJSON() ([]byte, error) {
	if mt.t == nil {
		return []byte(`"--"`), nil
	}
	return json.Marshal(mt.t.Format(format))
}

func (mt *MyTime) IsEmpty() bool {
	return mt.t == nil
}

func (mt *MyTime) String() string {
	if mt.t == nil {
		return "--"
	}
	return mt.t.Format(format)
}

type Test struct {
	Num         int64  `json:"num,omitempty"`
	InvalidTime MyTime `json:"invalidTime,string"`
	ValidTime   MyTime `json:"validTime,string"`
}

func main() {
	b, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	var t Test
	err = json.Unmarshal(b, &t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Num)
	fmt.Println(t.InvalidTime.IsEmpty())
	fmt.Println(t.InvalidTime.String())
	fmt.Println(t.ValidTime.IsEmpty())
	fmt.Println(t.ValidTime.String())
}
