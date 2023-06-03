package main

import (
	"fmt"
	"github.com/ipaqsa/go-ns/namespace"
	"runtime"
)

func uts_test() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	err := namespace.NewWrapper([]string{"uts"}, func() {
		fmt.Println("--------------------------IN-----------------------------------")
		run("hostname", "test")
		run("hostname", "")
	})
	check(err)

	fmt.Println("--------------------------OUT-----------------------------------")
	run("hostname", "")
}
