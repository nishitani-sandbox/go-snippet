package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ifaces {
		fmt.Println(i.Name)
		addresses, err := i.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, a := range addresses {
			switch a.(type) {
			case *net.IPNet:
				fmt.Println("IPNet")
				fmt.Println(a.Network())
				fmt.Println(a.String())
			case *net.IPAddr:
				fmt.Println("IPAddr")
				fmt.Println(a.Network())
				fmt.Println(a.String())
			}
		}
	}
}
