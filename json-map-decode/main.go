package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Endpoint struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {
	f, err := os.Open("./sample.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var m map[string]*Endpoint
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range m {
		if v == nil {
			fmt.Println("<nil>")
			continue
		}
		fmt.Printf("key: %s\n", k)
		fmt.Printf("host: %s\n", v.Host)
		fmt.Printf("port: %d\n", v.Port)
	}
}
