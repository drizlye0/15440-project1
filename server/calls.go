package main

import "os"

func open(name string) *RPCResponse {
	fd, err := os.Open(name)
	defer fd.Close()

	if err != nil {
		return &RPCResponse{Value: nil, Error: err}
	}

	return &RPCResponse{Value: fd, Error: nil}
}

func read(name string) *RPCResponse {
	content, err := os.ReadFile(name)
	if err != nil {
		return &RPCResponse{Value: nil, Error: err}
	}

	return &RPCResponse{Value: content, Error: err}
}

func write(name string, data []byte) *RPCResponse {
	fd, err := os.Open(name)
	if err != nil {
		return &RPCResponse{Value: nil, Error: err}
	}

	written, err := fd.Write(data)
	if err != nil {
		return &RPCResponse{Value: nil, Error: err}
	}

	return &RPCResponse{Value: written, Error: nil}
}
