package main

import (
	"log"
	"sync"
)

type aptutil struct {
	wg       *sync.WaitGroup
	packages []string

	recieved chan string
	err      chan error
	done     chan bool
}

func newAptutil() *aptutil {
	packages := make([]string, 0)
	wg := &sync.WaitGroup{}
	recieved := make(chan string)
	err := make(chan error)
	done := make(chan bool)

	return &aptutil{
		packages: packages,
		wg:       wg,
		recieved: recieved,
		err:      err,
		done:     done,
	}
}

func (a *aptutil) fetchPackages(packages []string) {
	for _, name := range packages {
		a.wg.Add(1)
		go func(s string) {
			a.recieved <- s
		}(name)
	}
	a.wg.Wait()
	a.done <- true
}

func (a *aptutil) wait() error {
	for {
		select {
		case recieved := <-a.recieved:
			a.packages = append(a.packages, recieved)
			a.wg.Done()
		case err := <-a.err:
			return err
		case <-a.done:
			return nil
		}
	}
}

func (a *aptutil) log() {
	for _, name := range a.packages {
		log.Println(name)
	}
}

func main() {
	a := newAptutil()
	packages := []string{
		"hoge",
		"foo",
		"bar",
		"baz",
	}
	go a.fetchPackages(packages)
	err := a.wait()
	if err != nil {
		log.Fatal(err)
	}
	a.log()
}
