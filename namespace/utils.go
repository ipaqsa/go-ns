package namespace

import "golang.org/x/sys/unix"

func newNamespace() *Namespace {
	ns := Namespace{}
	ns.ns = map[string]_ns{}
	for _, n := range NSs {
		ns.ns[n] = _ns{-1, "", false}
	}
	ns.name = ""
	return &ns
}

func (ns *Namespace) SetType(nstypes []string) {
	for _, nstype := range nstypes {
		if entry, ok := ns.ns[nstype]; ok {
			entry.share = true
		}
	}
}

func (ns *Namespace) Equal(other *Namespace) bool {
	if ns == other {
		return true
	}
	for _, n := range NSs {
		if !ns.EqualNS(other, n) {
			return false
		}
	}
	return true
}
func (ns *Namespace) EqualNS(other *Namespace, nstype string) bool {
	return equalFd(ns.ns[nstype].fd, other.ns[nstype].fd)
}
func equalFd(fd1, fd2 int) bool {
	if fd1 == -1 || fd2 == -1 {
		return false
	}
	var s1, s2 unix.Stat_t
	if err := unix.Fstat(fd1, &s1); err != nil {
		return false
	}
	if err := unix.Fstat(fd2, &s2); err != nil {
		return false
	}
	return (s1.Dev == s2.Dev) && (s1.Ino == s2.Ino)
}

func (ns *Namespace) IsOpen() bool {
	for _, n := range NSs {
		if !ns.IsOpenNS(n) {
			return false
		}
	}
	return true
}
func (ns *Namespace) IsOpenNS(nstype string) bool {
	if entry, ok := ns.ns[nstype]; ok {
		if entry.share && entry.fd == -1 {
			return false
		}
	}
	return true
}

func (ns *Namespace) Close() error {
	for _, n := range NSs {
		err := ns.CloseNS(n)
		if err != nil {
			return err
		}
	}
	return nil
}
func (ns *Namespace) CloseNS(nstype string) error {
	if entry, ok := ns.ns[nstype]; ok {
		return unix.Close(entry.fd)
	}
	return nil
}

func OpenNS(path string) (int, error) {
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_CLOEXEC, 0)
	if err != nil {
		return -1, err
	}
	return fd, nil
}
