package utils

import (
	"fmt"
)

func Itob(value int32) []byte {
	return []byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)}
}

func Btoi(data []byte) int {
	var num int

	switch len(data) {
	case 1:
		num = int(data[0])
	case 4:
		num = int(data[0])<<24 + int(data[1])<<16 + int(data[2])<<8 + int(data[3])
	default:
		panic(fmt.Sprintf("Number of bytes of %d not supported", len(data)))
	}

	return num
}
