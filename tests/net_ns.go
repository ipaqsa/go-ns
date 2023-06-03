package main

import (
	"fmt"
	"github.com/ipaqsa/go-ns/namespace"
	"net"
	"runtime"
)

func net_test() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	err := namespace.NewWrapper([]string{"net"}, func() {
		ifcs, err := net.Interfaces()
		check(err)
		fmt.Printf("Interfaces count: %d\n", len(ifcs))
	})
	check(err)

	fmt.Println("--------------------------OUT-----------------------------------")
	ifcs, err := net.Interfaces()
	check(err)
	fmt.Printf("Interfaces count: %d\n", len(ifcs))
}
