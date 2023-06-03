package main

import (
	"fmt"
	"github.com/ipaqsa/go-ns/namespace"
	"os"
	"runtime"
)

func mnt_test() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	err := namespace.NewWrapper([]string{"mnt"}, func() {
		fmt.Println("--------------------------IN-----------------------------------")
		l, err := os.Readlink(fmt.Sprintf("/proc/%d/ns/mnt", os.Getpid()))
		check(err)
		fmt.Println(l)
	})
	check(err)

	fmt.Println("--------------------------OUT-----------------------------------")
	l, _ := os.Readlink(fmt.Sprintf("/proc/%d/ns/mnt", os.Getpid()))
	check(err)
	fmt.Println(l)
}
