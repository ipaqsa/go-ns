package namespace

import (
	"syscall"
)

func PivotRoot(newroot, oldroot string) error {
	err := syscall.PivotRoot(newroot, oldroot)
	if err != nil {
		return err
	}
	return nil
}

func ChangeDir(path string) error {
	err := syscall.Chdir(path)
	if err != nil {
		return err
	}
	return nil
}

func ChangeRoot(path string) error {
	err := syscall.Chroot(path)
	if err != nil {
		return err
	}
	return nil
}
