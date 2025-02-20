package operations_test

import (
	"bytes"
	"fmt"
	"testing"

	"martinjonson.com/ccbf/ccbf/operations"
)

func TestObjectMethods(t *testing.T) {
	testcases := []struct {
		opCode        byte
		byteCount     int
		repetitions   int
		standardBytes []byte
		parsedSymbols int
	}{
		{opCode: 1, byteCount: 0, standardBytes: []byte{}, parsedSymbols: 1},
		{opCode: 9, byteCount: 1, repetitions: 6, standardBytes: []byte{6}, parsedSymbols: 6},
		{opCode: 13, byteCount: 0, standardBytes: []byte{}, parsedSymbols: 4},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Operation methods for op code %d", tc.opCode), func(t *testing.T) {
			operation := operations.OperationForOpCode(tc.opCode)

			byteCount := operation.GetParameterByteCount()
			if byteCount != tc.byteCount {
				t.Fatalf("Number of bytes should be %d but was %d", tc.byteCount, byteCount)
			}

			standardBytes := operation.StandardParameterBytes(tc.repetitions)
			if !bytes.Equal(standardBytes, tc.standardBytes) {
				t.Fatalf("Standard bytes should be %v but was %v", tc.standardBytes, standardBytes)
			}

			parsedSymbols := operation.ParsedSymbols(tc.repetitions)
			if parsedSymbols != tc.parsedSymbols {
				t.Fatalf("Number of parsed symbols should be %d but was %d", tc.parsedSymbols, parsedSymbols)
			}
		})
	}
}

func TestOperationPatterns(t *testing.T) {
	patterns := operations.OperationPatterns()

	if len(patterns) < 1 {
		t.Fatal("OperationPatterns should return several patterns")
	}
}

func TestOperationForPattern(t *testing.T) {
	testcases := []struct {
		pattern        string
		repeated       bool
		expectedOpCode byte
	}{
		{pattern: ">", repeated: false, expectedOpCode: 1},
		{pattern: ">", repeated: true, expectedOpCode: 11},
		{pattern: ",", repeated: false, expectedOpCode: 6},
		{pattern: ",", repeated: true, expectedOpCode: 6},
	}

	for _, tc := range testcases {
		operation := operations.OperationForPattern(tc.pattern, tc.repeated)

		if operation.GetOpCode() != tc.expectedOpCode {
			t.Fatalf("OperationForPattern should have returned operation with opcode %d for pattern %s and repeated %t", tc.expectedOpCode, tc.pattern, tc.repeated)
		}
	}
}

func TestOperationForPatternIncorrect(t *testing.T) {
	testcases := []struct {
		pattern  string
		repeated bool
	}{
		{pattern: "a", repeated: false},
		{pattern: "a", repeated: true},
		{pattern: "b", repeated: false},
		{pattern: "b", repeated: true},
	}

	for _, tc := range testcases {
		operation := operations.OperationForPattern(tc.pattern, tc.repeated)

		if operation != nil {
			t.Fatalf("OperationForPattern should have returned nil operation for pattern %s and repeated %t", tc.pattern, tc.repeated)
		}
	}
}

func TestOperationForOpCode(t *testing.T) {
	operation := operations.OperationForOpCode(1)

	if operation.GetPattern() != ">" {
		t.Fatal("OperationForOpCode should have returned operation >")
	}
}

func TestOperationForOpCodeIncorrect(t *testing.T) {
	testcases := []byte{0, 15}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Incorrect %d", tc), func(t *testing.T) {
			operation := operations.OperationForOpCode(tc)

			if operation != nil {
				t.Fatalf("%d should not be a valid op code", tc)
			}
		})
	}
}
