package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run(command string, args string) {
	var cmd *exec.Cmd
	if args == "" {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(command, args)
	}
	o, err := cmd.Output()
	check(err)
	fmt.Println(string(o))
}

func ls(path string) {
	dir, err := os.ReadDir(path)
	check(err)
	for _, f := range dir {
		fmt.Println(f.Name())
	}
}

func cat(path string) {
	b, err := os.ReadFile(path)
	check(err)
	fmt.Println(string(b))
}

func mountProc() error {
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		return err
	}
	return nil
}
