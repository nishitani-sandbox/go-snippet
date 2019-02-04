package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var (
		b      bytes.Buffer
		actual int64
	)
	b.Write([]byte("This is test for io.EOF\n"))
	expected := b.Len()
	for {
		n, err := io.CopyN(ioutil.Discard, &b, 5)
		fmt.Printf("written: %d bytes, err: %v\n", n, err)
		actual += n
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy failed: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("expected: %d, actual: %d\n", expected, actual)
}
