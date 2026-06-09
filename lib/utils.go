package lib

import "encoding/binary"

func ByteArrToPtr(arr []byte) uintptr {
	for i := 0; i < 8-len(arr); i++ {
		arr = append([]byte{0x0, 0x0}, arr...)
	}

	ptr := binary.BigEndian.Uint64(arr)
	return uintptr(ptr)
}
