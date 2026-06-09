package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"
)

var addr string = "localhost:8080"

func sendRequest(req *RPCRequest) (*RPCResponse, error) {
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	buf := req.Encode()
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	response = strings.TrimSpace(response)
	buf = bytes.NewBuffer([]byte(response))
	res := DecodeRPCResponse(buf)

	if res == nil {
		return nil, fmt.Errorf("Nil response")
	}

	if res.Error != nil {
		return nil, err
	}

	return res, nil
}

func Open(fileName string) (uintptr, error) {
	req := &RPCRequest{
		CallSignature: "open",
		FileName:      fileName,
	}

	res, err := sendRequest(req)
	if err != nil {
		return 0, err
	}

	return res.SysFd, nil
}

func Close(sysFd uintptr, fileName string) error {
	req := &RPCRequest{
		CallSignature: "close",
		SysFd:         sysFd,
		FileName:      fileName,
	}

	_, err := sendRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func Read(fileName string) ([]byte, error) {
	req := &RPCRequest{
		CallSignature: "read",
		FileName:      fileName,
	}

	res, err := sendRequest(req)
	if err != nil {
		return nil, err
	}

	return res.Read, nil
}

func Write(sysFd uintptr, fileName string, data []byte) (int, error) {
	req := &RPCRequest{
		CallSignature: "write",
		SysFd:         sysFd,
		FileName:      fileName,
		Data:          data,
	}

	res, err := sendRequest(req)
	if err != nil {
		return -1, err
	}

	return res.Written, nil
}
