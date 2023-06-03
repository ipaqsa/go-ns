package main

import (
	"fmt"
	"github.com/ipaqsa/go-ns/namespace"
	"runtime"
)

func user_test() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	err := namespace.NewWrapper([]string{"user"}, func() {
		fmt.Println("--------------------------IN-----------------------------------")
		run("id", "")
	})
	check(err)

	fmt.Println("--------------------------OUT-----------------------------------")
	run("id", "")
}
