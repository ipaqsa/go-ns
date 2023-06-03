clean:
	rm -rf build

build: clean
	go build -o build/ns cmd/main.go

build-test: clean
	go build -o build/test github.com/ipaqsa/go-ns/tests

run-test: build-test
	sudo strace -v -o trace.log build/test