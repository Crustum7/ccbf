package operations_test

import (
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
		})
	}
}

func TestOperationPatterns(t *testing.T) {
	patterns := operations.OperationPatterns()

	if len(patterns) < 1 {
		t.Fatal("OperationPatterns should return several patterns")
	}
}

func TestOperationForOpCode(t *testing.T) {
	operation := operations.OperationForOpCode(1)

	if operation.GetPattern() != ">" {
		t.Fatal("OperationForOpCode should have returned operation >")
	}
}

func TestOperationForOpCodeIncorrect(t *testing.T) {
	testcases := []byte{0, 16}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Incorrect %d", tc), func(t *testing.T) {
			operation := operations.OperationForOpCode(tc)

			if operation != nil {
				t.Fatalf("%d should not be a valid op code", tc)
			}
		})
	}
}
