package compiler

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

func TestCompileSimpleProgram(t *testing.T) {
	program := "+++[-]"
	bytes := CompileProgram(program)
	expectedBytes := []byte{9, 3, 14}

	if !isEqual(bytes, expectedBytes) {
		t.Fatalf("Compiled byte code should be %d but was %d", expectedBytes, bytes)
	}
}

func TestCompileEveryInstructionButJump(t *testing.T) {
	program := "><+-.,+++--->>><<<[-]>[-]abcde"
	bytes := CompileProgram(program)
	expectedBytes := []byte{1, 2, 3, 4, 5, 6, 9, 3, 10, 3, 11, 3, 12, 3, 13, 14}

	if !isEqual(bytes, expectedBytes) {
		t.Fatalf("Compiled byte code should be %d but was %d", expectedBytes, bytes)
	}
}

func TestCompileJumps(t *testing.T) {
	program := "[+]"
	bytes := CompileProgram(program)
	expectedBytes := []byte{7, 0, 0, 0, 10, 3, 8, 0, 0, 0, 4}

	if !isEqual(bytes, expectedBytes) {
		t.Fatalf("Compiled byte code should be %d but was %d", expectedBytes, bytes)
	}
}
