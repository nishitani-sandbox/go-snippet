package main

import (
	"context"
	"log"
	"os/exec"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := exec.CommandContext(ctx, "dpkg", "-l", "hoge")
	out, err := c.Output()
	if exitErr, ok := err.(*exec.ExitError); ok {
		log.Println(string(exitErr.Stderr))
		return
	}
	log.Println(out)
	log.Println(err)
}
