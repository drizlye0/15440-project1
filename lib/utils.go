package lib

import (
	"fmt"
	"strconv"
)

func ByteArrToPtr(arr []byte) (uintptr, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("byte array is empty")
	}

	val, err := strconv.ParseUint(string(arr), 10, 0)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse to ptr")
	}

	return uintptr(val), nil
}
