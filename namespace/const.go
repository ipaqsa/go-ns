package namespace

import "golang.org/x/sys/unix"

const (
	// CLONE_NEWUTS Work
	CLONE_NEWUTS = unix.CLONE_NEWUTS

	// CLONE_NEWNS Maybe work
	CLONE_NEWNS = unix.CLONE_NEWNS

	// CLONE_NEWCGROUP Don`t test
	CLONE_NEWCGROUP = unix.CLONE_NEWCGROUP

	// CLONE_NEWTIME Maybe work
	CLONE_NEWTIME = unix.CLONE_NEWTIME

	// CLONE_NEWIPC Work
	CLONE_NEWIPC = unix.CLONE_NEWIPC //Work

	// CLONE_NEWUSER Don`t work
	CLONE_NEWUSER = unix.CLONE_NEWUSER

	// CLONE_NEWPID Don`t work
	CLONE_NEWPID = unix.CLONE_NEWPID

	// CLONE_NEWNET Work
	CLONE_NEWNET = unix.CLONE_NEWNET //Work

	// CLONE_IO Dont test
	CLONE_IO = unix.CLONE_IO
)

const bindMountPath = "/run/netns"

const (
	cgroupPath = "cgroup"
	ipcPath    = "ipc"
	mntPath    = "mnt"
	netPath    = "net"
	pidPath    = "pid"
	timePath   = "time"
	userPath   = "user"
	utsPath    = "uts"
)

var NSs = []string{pidPath, userPath, utsPath, mntPath, netPath, timePath, ipcPath, cgroupPath}

var CloneFlags = map[string]int{pidPath: CLONE_NEWPID, userPath: CLONE_NEWUSER, utsPath: CLONE_NEWUTS, mntPath: CLONE_NEWNS,
	netPath: CLONE_NEWNET, timePath: CLONE_NEWTIME, ipcPath: CLONE_NEWIPC, cgroupPath: CLONE_NEWCGROUP}
