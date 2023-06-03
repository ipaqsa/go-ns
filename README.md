## Simple namespaces wrapper

## Overview

Project provides wrapper and management functions of Linux namespace.

Simple use:

```
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
```

## Provided functions

- Current - return current namespace from current pid and tid

- Set - set chosen namespaces

- New - create new namespace via unshare(don`t need set)

- FromPid - return namespace from pid

- FromThread - return namespace from tid

- NewNetNamed - create named net namespace

- RemoveNetNamed - delete named net namespace

- NetFromName - return named net namespace

- NewWrapper - create new namespace, call function and return current ns

- Wrapper - call function and return current ns


## Notes

```
	// Work
	uts = CLONE_NEWUTS = unix.CLONE_NEWUTS

	// Maybe work
	mnt = CLONE_NEWNS = unix.CLONE_NEWNS

	// Don`t test
	cgroup = CLONE_NEWCGROUP = unix.CLONE_NEWCGROUP

	// Maybe work
	time = CLONE_NEWTIME = unix.CLONE_NEWTIME

	// Work
	ipc = CLONE_NEWIPC = unix.CLONE_NEWIPC //Work

	// Don`t work
	user = CLONE_NEWUSER = unix.CLONE_NEWUSER

	// Don`t work
	pid = CLONE_NEWPID = unix.CLONE_NEWPID

	// Work
	net = CLONE_NEWNET = unix.CLONE_NEWNET //Work

```


