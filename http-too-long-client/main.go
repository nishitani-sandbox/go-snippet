package main

import (
	"context"
	"os/exec"
	"time"

	"github.dev.cybozu.co.jp/hazama/infra/go/src/github.com/cybozu-go/log"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "curl", "localhost:8888")
	err := cmd.Run()
	if err != nil {
		log.ErrorExit(err)
	}
}
