package namespace

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"path"
	"path/filepath"
)

func NewWrapper(NSs []string, f func()) error {
	current, err := Current()
	if err != nil {
		return err
	}
	_, err = New(NSs)
	if err != nil {
		return err
	}
	f()
	err = current.Set(NSs)
	if err != nil {
		return err
	}
	return nil
}

func (ns *Namespace) Wrapper(NSs []string, f func()) error {
	current, err := Current()
	if err != nil {
		return err
	}
	err = ns.Set(NSs)
	if err != nil {
		return err
	}
	f()
	err = current.Set(NSs)
	if err != nil {
		return err
	}
	return nil
}

func Current() (*Namespace, error) {
	return FromThread(os.Getpid(), unix.Gettid())
}

func FromThread(pid, tid int) (*Namespace, error) {
	ns := newNamespace()
	for _, n := range NSs {
		if entry, ok := ns.ns[n]; ok {
			entry.path = fmt.Sprintf("/proc/%d/task/%d/ns/%s", pid, tid, n)
			fd, err := OpenNS(entry.path)
			if err != nil {
				return nil, err
			}
			entry.fd = fd
			ns.ns[n] = entry
		}
	}
	return ns, nil
}

func FromPid(pid int) (*Namespace, error) {
	ns := newNamespace()
	for _, n := range NSs {
		if entry, ok := ns.ns[n]; ok {
			entry.path = fmt.Sprintf("/proc/%d/ns/%s", pid, n)
			fd, err := OpenNS(entry.path)
			if err != nil {
				return nil, err
			}
			entry.fd = fd
			ns.ns[n] = entry
		}
	}
	return ns, nil
}

func New(NSs []string) (*Namespace, error) {
	var flag int = 0
	for _, ns := range NSs {
		if val, ok := CloneFlags[ns]; ok == true {
			flag = flag | val
			continue
		}
		return nil, errors.New("unsupported ns " + ns)
	}
	if err := unix.Unshare(flag); err != nil {
		return nil, err
	}
	return Current()
}

func (ns *Namespace) Set(NSs []string) error {
	for _, n := range NSs {
		if entry, ok := ns.ns[n]; ok {
			entry.share = true
			ns.ns[n] = entry
		}
	}
	for key, val := range ns.ns {
		if !val.share {
			continue
		}
		fmt.Printf("Preparing %s...\n", key)
		err := unix.Setns(val.fd, CloneFlags[key])
		if err != nil {
			return err
		}
		fmt.Printf("Set %s\n", key)
	}
	return nil
}

func NewNetNamed(name string, NSs []string) (*Namespace, error) {
	if _, err := os.Stat(bindMountPath); os.IsNotExist(err) {
		err = os.MkdirAll(bindMountPath, 0o755)
		if err != nil {
			return nil, err
		}
	}

	ns, err := New(NSs)
	if err != nil {
		return nil, err
	}

	namedPath := path.Join(bindMountPath, name)

	f, err := os.OpenFile(namedPath, os.O_CREATE|os.O_EXCL, 0o444)
	if err != nil {
		ns.Close()
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}

	nsPath := fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid())
	err = unix.Mount(nsPath, namedPath, "bind", unix.MS_BIND, "")
	if err != nil {
		ns.Close()
		return nil, err
	}

	return ns, nil
}

func RemoveNetNamed(name string) error {
	namedPath := path.Join(bindMountPath, name)

	err := unix.Unmount(namedPath, unix.MNT_DETACH)
	if err != nil {
		return err
	}
	return os.Remove(namedPath)
}

func NetFromName(name string) (*Namespace, error) {
	ns := newNamespace()
	if entry, ok := ns.ns["net"]; ok {
		entry.path = fmt.Sprintf(filepath.Join(bindMountPath, name))
		fd, err := OpenNS(entry.path)
		if err != nil {
			return nil, err
		}
		entry.fd = fd
		entry.share = true
		ns.ns["net"] = entry
	}
	return ns, nil
}
