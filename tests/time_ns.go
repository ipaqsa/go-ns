package main

import (
	"fmt"
	"github.com/ipaqsa/go-ns/namespace"
	"os"
	"runtime"
	"strconv"
)

// Use of time namespaces requires a  kernel  that  is  configured  with  the  CONFIG_TIME_NS  option.
//
// Note  that  time namespaces do not virtualize the CLOCK_REALTIME clock.  Virtualization of
// this clock was avoided for reasons of complexity and overhead within the kernel.
//
// For compatibility with  the  initial  implementation,  when  writing  a  clock-id  to  the
// /proc/[pid]/timens_offsets file, the numerical values of the IDs can be written instead of
// the symbolic names show above; i.e., 1 instead of monotonic, and 7  instead  of  boottime.
// For redability, the use of the symbolic names over the numbers is preferred.
//
// The  motivation for adding time namespaces was to allow the monotonic and boot-time clocks
// to maintain consistent values during container migration and checkpoint/restore.

func time_test() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	err := namespace.NewWrapper([]string{"time"}, func() {
		fmt.Println("--------------------------IN-----------------------------------")
		readlink, err := os.Readlink("/proc/" + strconv.Itoa(os.Getpid()) + "/ns/time")
		check(err)
		println(readlink)
	})
	check(err)

	fmt.Println("--------------------------OUT-----------------------------------")
	readlink, err := os.Readlink("/proc/" + strconv.Itoa(os.Getpid()) + "/ns/time")
	check(err)
	println(readlink)
}
