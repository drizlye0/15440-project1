package main

import (
	"os"
	"syscall"

	"github.com/drizlye0/15440-project1/lib"
)

func open(name string) *lib.RPCResponse {
	fd, err := os.Open(name)
	if err != nil {
		return &lib.RPCResponse{Error: err}
	}

	defer fd.Close()

	dupFd, err := syscall.Dup(int(fd.Fd()))
	if err != nil {
		return &lib.RPCResponse{Error: err}
	}

	return &lib.RPCResponse{SysFd: uintptr(dupFd)}
}

func close(sysFd uintptr, name string) *lib.RPCResponse {
	fd := os.NewFile(sysFd, name)
	err := fd.Close()
	if err != nil {
		return &lib.RPCResponse{Error: err}
	}

	return &lib.RPCResponse{}
}

func read(name string) *lib.RPCResponse {
	r, err := os.ReadFile(name)
	if err != nil {
		return &lib.RPCResponse{Error: err}
	}

	return &lib.RPCResponse{Read: r}
}

func write(name string, data []byte) *lib.RPCResponse {
	fd, err := os.OpenFile(name, os.O_WRONLY, 0644)
	if err != nil {
		return &lib.RPCResponse{Error: err}
	}
	defer fd.Close()

	w, err := fd.WriteString(string(data))
	if err != nil {
		return &lib.RPCResponse{Error: err}
	}

	return &lib.RPCResponse{Written: w}
}
