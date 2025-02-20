package utils

import (
	"bytes"
	"testing"
)

func TestIntegerToByteArraySmallestByte(t *testing.T) {
	number := 123
	byteArr := Itob(int32(number))
	if !bytes.Equal(byteArr, []byte{0, 0, 0, byte(number)}) {
		t.Fatalf("%d does not equal %d", byteArr, number)
	}
}

func TestIntegerToByteArray(t *testing.T) {
	number := 16909060
	byteArr := Itob(int32(number))
	if !bytes.Equal(byteArr, []byte{1, 2, 3, 4}) {
		t.Fatalf("%d does not equal %d", byteArr, number)
	}
}

func TestByteArrayToInt32(t *testing.T) {
	expected := 16909060
	number := Btoi([]byte{1, 2, 3, 4})
	if number != expected {
		t.Fatalf("%d does not equal %d", number, expected)
	}
}

func TestByteArrayToInt8(t *testing.T) {
	expected := 123
	number := Btoi([]byte{123})
	if number != expected {
		t.Fatalf("%d does not equal %d", number, expected)
	}
}

func TestByteArrayToInt16(t *testing.T) {
	shouldPanic(t, func() { Btoi([]byte{1, 2}) })
}

func TestAssignBytes(t *testing.T) {
	from := []byte{1, 2, 3, 4}
	to := []byte{0, 0, 0, 0}

	AssignBytes(to, from)
	if !bytes.Equal(to, from) {
		t.Fatalf("%d does not equal %d", to, from)
	}
}

func TestAssignBytesWrongSize(t *testing.T) {
	from := []byte{1, 2, 3, 4}
	to := []byte{0, 0, 0}

	shouldPanic(t, func() { AssignBytes(to, from) })
}
