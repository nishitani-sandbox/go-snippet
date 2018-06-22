package main

import (
	"fmt"
	"strings"

	"github.com/cybozu-go/cmd"
	"github.com/pborman/uuid"
)

func main() {
	fmt.Println(strings.ToUpper(
		strings.Replace(uuid.NewRandom().String(), "-", "", -1),
	))

	fmt.Println(strings.ToUpper(
		strings.Replace(cmd.GenerateID(), "-", "", -1),
	))
}
