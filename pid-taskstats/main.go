package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mdlayher/taskstats"
)

func main() {
	c, err := taskstats.New()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 1 {
		log.Fatal("please specify PID")
	}
	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	stats, err := c.PID(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", stats)
}
