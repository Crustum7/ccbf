package bytecode

import "testing"

func isEqual(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("Should have panicked")
}

func TestIntegerToByteArraySmallestByte(t *testing.T) {
	number := 123
	byteArr := itob(int32(number))
	if !isEqual(byteArr, []byte{0, 0, 0, byte(number)}) {
		t.Fatalf("%d does not equal %d", byteArr, number)
	}
}

func TestIntegerToByteArray(t *testing.T) {
	number := 16909060
	byteArr := itob(int32(number))
	if !isEqual(byteArr, []byte{1, 2, 3, 4}) {
		t.Fatalf("%d does not equal %d", byteArr, number)
	}
}

func TestByteArrayToInt32(t *testing.T) {
	expected := 16909060
	number := btoi([]byte{1, 2, 3, 4})
	if number != expected {
		t.Fatalf("%d does not equal %d", number, expected)
	}
}

func TestByteArrayToInt8(t *testing.T) {
	expected := 123
	number := btoi([]byte{123})
	if number != expected {
		t.Fatalf("%d does not equal %d", number, expected)
	}
}

func TestByteArrayToInt16(t *testing.T) {
	shouldPanic(t, func() { btoi([]byte{1, 2}) })
}

func TestAssignBytes(t *testing.T) {
	from := []byte{1, 2, 3, 4}
	to := []byte{0, 0, 0, 0}

	assignBytes(to, from)
	if !isEqual(to, from) {
		t.Fatalf("%d does not equal %d", to, from)
	}
}

func TestAssignBytesWrongSize(t *testing.T) {
	from := []byte{1, 2, 3, 4}
	to := []byte{0, 0, 0}

	shouldPanic(t, func() { assignBytes(to, from) })
}
