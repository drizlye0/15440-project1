package main

import (
	"os"
	"syscall"
)

func open(name string) *RPCResponse {
	fd, err := os.Open(name)
	if err != nil {
		return &RPCResponse{error: err}
	}

	defer fd.Close()

	dupFd, err := syscall.Dup(int(fd.Fd()))
	if err != nil {
		return &RPCResponse{error: err}
	}

	return &RPCResponse{sysFd: uintptr(dupFd)}
}

func close(sysFd uintptr, name string) *RPCResponse {
	fd := os.NewFile(sysFd, name)
	err := fd.Close()
	if err != nil {
		return &RPCResponse{error: err}
	}

	return &RPCResponse{}
}

func read(name string) *RPCResponse {
	r, err := os.ReadFile(name)
	if err != nil {
		return &RPCResponse{error: err}
	}

	return &RPCResponse{read: r}
}

func write(sysFd uintptr, name string, data []byte) *RPCResponse {
	fd := os.NewFile(sysFd, name)
	defer fd.Close()
	w, err := fd.Write(data)

	if err != nil {
		return &RPCResponse{error: err}
	}

	return &RPCResponse{written: w}
}
