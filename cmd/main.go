package main

import (
	"github.com/ipaqsa/go-ns/namespace"
	"log"
	"runtime"
)

var nss = []string{"net"}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	err := namespace.NewWrapper(nss, func() {
		//Do something
	})
	if err != nil {
		log.Fatal(err)
	}
}
