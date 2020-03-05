package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var (
	ErrInvalidTag = errors.New("invalid tag")
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func ParseFromGitTag(t string) (*Version, error) {
	splitted := strings.Fields(t)
	if len(splitted) < 2 {
		return nil, ErrInvalidTag
	}

	var major, minor, patch int
	_, err := fmt.Sscanf(splitted[1], "refs/tags/go%d.%d.%d\n", &major, &minor, &patch)
	if err == nil {
		return &Version{major, minor, patch}, nil
	}
	_, err = fmt.Sscanf(splitted[1], "refs/tags/go%d.%d\n", &major, &minor)
	if err == nil {
		return &Version{Major: major, Minor: minor}, nil
	}

	return nil, ErrInvalidTag
}

func main() {
	var stdout bytes.Buffer
	cmd := exec.Command("git", "ls-remote", "-t", "https://go.googlesource.com/go")
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(&stdout)
	for s.Scan() {
		v, err := ParseFromGitTag(s.Text())
		if err != nil {
			continue
		}
		fmt.Printf("Version: %d.%d.%d\n", v.Major, v.Minor, v.Patch)
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}
