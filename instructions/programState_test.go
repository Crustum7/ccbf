package instructions

import "testing"

func TestInitProgramStateCapacity(t *testing.T) {
	expectedCap := 32

	state := InitProgramState()

	cap := state.Capacity()
	if cap != expectedCap {
		t.Fatalf("Program state should start with capacity %d, but started with %d", expectedCap, cap)
	}
}

func TestInitProgramStatePos(t *testing.T) {
	expectedPos := 0

	state := InitProgramState()

	pos := state.pos
	if pos != expectedPos {
		t.Fatalf("Program state should start with pos %d, but started with %d", expectedPos, pos)
	}
}

func TestInitProgramStateProgramCounter(t *testing.T) {
	expectedPc := 0

	state := InitProgramState()

	pc := state.pos
	if pc != expectedPc {
		t.Fatalf("Program state should start with program counter %d, but started with %d", expectedPc, pc)
	}
}

func TestInitProgramStateDataZero(t *testing.T) {
	state := InitProgramState()

	for _, val := range state.data {
		if val != 0 {
			t.Fatalf("Data should be initialized to 0 but the number %d was found", val)
		}
	}
}

func TestIncrementProgramCounter(t *testing.T) {
	state := InitProgramState()
	initialPc := 5
	state.programCounter = initialPc

	state.IncrementProgramCounter()

	if state.GetProgramCounter() != initialPc+1 {
		t.Fatalf("Program counter should be %d after IncrementProgramCounter but was %d", initialPc+1, state.GetProgramCounter())
	}
}

func TestIncrementProgramCounterWith(t *testing.T) {
	state := InitProgramState()
	initialPc := 5
	step := 42
	state.programCounter = initialPc

	state.IncrementProgramCounterWith(step)

	if state.GetProgramCounter() != initialPc+step {
		t.Fatalf("Program counter should be %d after IncrementProgramCounter but was %d", initialPc+step, state.GetProgramCounter())
	}
}
