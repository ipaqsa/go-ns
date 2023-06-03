package main

import (
	"fmt"
	"github.com/ipaqsa/go-ns/namespace"
	"runtime"
)

func ipc_test() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	err := namespace.NewWrapper([]string{"ipc"}, func() {
		fmt.Println("--------------------------IN-----------------------------------")
		run("ipcs", "")
	})
	check(err)

	fmt.Println("--------------------------OUT-----------------------------------")
	run("ipcs", "")
}
