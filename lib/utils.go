package lib

import (
	"encoding/binary"
	"fmt"
)

func ByteArrToPtr(arr []byte) uintptr {
	if len(arr) <= 1 {
		return 0
	}

	fmt.Println(len(arr))

	for i := 0; i < 8-len(arr); i++ {
		arr = append([]byte{0x0, 0x0}, arr...)
	}

	ptr := binary.BigEndian.Uint64(arr)
	return uintptr(ptr)
}
