package compiler

import (
	"bytes"
	"testing"
)

func TestCompileSimpleProgram(t *testing.T) {
	program := "+++[-]"
	code := CompileProgram(program)
	expectedBytes := []byte{9, 3, 14}

	if !bytes.Equal(code, expectedBytes) {
		t.Fatalf("Compiled byte code should be %d but was %d", expectedBytes, code)
	}
}

func TestCompileEveryInstructionButJump(t *testing.T) {
	program := "><+-.,+++--->>><<<[-]>[-]abcde"
	code := CompileProgram(program)
	expectedBytes := []byte{1, 2, 3, 4, 5, 6, 9, 3, 10, 3, 11, 3, 12, 3, 13, 14}

	if !bytes.Equal(code, expectedBytes) {
		t.Fatalf("Compiled byte code should be %d but was %d", expectedBytes, code)
	}
}

func TestCompileJumps(t *testing.T) {
	program := "[+]"
	code := CompileProgram(program)
	expectedBytes := []byte{7, 0, 0, 0, 10, 3, 8, 0, 0, 0, 4}

	if !bytes.Equal(code, expectedBytes) {
		t.Fatalf("Compiled byte code should be %d but was %d", expectedBytes, code)
	}
}
