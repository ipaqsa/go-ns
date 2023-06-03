package namespace

type Namespaces struct {
	cgroup _ns
	ipc    _ns
	mnt    _ns
	net    _ns
	pid    _ns
	time   _ns
	user   _ns
	uts    _ns
}

type NamespacesFd struct {
	cgroup int
	ipc    int
	mnt    int
	net    int
	pid    int
	time   int
	user   int
	uts    int
}

type Namespace struct {
	ns   map[string]_ns
	name string
}

type _ns struct {
	fd    int
	path  string
	share bool
}
