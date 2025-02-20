package instructions

import (
	"fmt"
	"testing"
)

func TestInitProgramStateCapacity(t *testing.T) {
	expectedCap := 32

	state := InitProgramState(expectedCap)

	cap := state.Capacity()
	if cap != expectedCap {
		t.Fatalf("Program state should start with capacity %d, but started with %d", expectedCap, cap)
	}
}

func TestAdjustedCapacity(t *testing.T) {
	testCases := []struct {
		initialCapacity  int
		expectedCapacity int
		position         int
	}{
		{initialCapacity: 4, expectedCapacity: 4, position: 3},
		{initialCapacity: 4, expectedCapacity: 8, position: 4},
		{initialCapacity: 4, expectedCapacity: 8, position: 5},
		{initialCapacity: 4, expectedCapacity: 128, position: 100},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("initial capacity %d, position %d, expected capacity %d", tc.initialCapacity, tc.position, tc.expectedCapacity), func(t *testing.T) {
			state := InitProgramState(tc.initialCapacity)
			state.pos = tc.position

			state.AdjustCapacity()

			cap := state.Capacity()
			if cap != tc.expectedCapacity {
				t.Fatalf("Program state should have capacity %d, but has %d", tc.expectedCapacity, cap)
			}
		})
	}
}

func TestInitProgramStatePos(t *testing.T) {
	expectedPos := 0

	state := InitProgramState(32)

	pos := state.pos
	if pos != expectedPos {
		t.Fatalf("Program state should start with pos %d, but started with %d", expectedPos, pos)
	}
}

func TestInitProgramStateProgramCounter(t *testing.T) {
	expectedPc := 0

	state := InitProgramState(32)

	pc := state.pos
	if pc != expectedPc {
		t.Fatalf("Program state should start with program counter %d, but started with %d", expectedPc, pc)
	}
}

func TestInitProgramStateDataZero(t *testing.T) {
	state := InitProgramState(32)

	for _, val := range state.data {
		if val != 0 {
			t.Fatalf("Data should be initialized to 0 but the number %d was found", val)
		}
	}
}
