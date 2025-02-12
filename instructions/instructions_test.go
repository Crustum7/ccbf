package instructions

import "testing"

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("Should have panicked")
}

func TestIncreasedCapacity(t *testing.T) {
	state := InitProgramState()

	if state.Capacity() != 32 {
		t.Fatal("Incorrect initial capacity for program state")
	}

	IncPosWith(&state, 40)
	if state.Capacity() != 32*2 {
		t.Fatal("Program state capacity did not double")
	}
}

func TestNegativePositionShouldPanic(t *testing.T) {
	state := InitProgramState()
	shouldPanic(t, func() { DecPos(&state) })
}

func TestIncreasePositionOneStep(t *testing.T) {
	state := InitProgramState()
	initialPos := state.pos

	IncPos(&state)

	if state.pos != initialPos+1 {
		t.Fatal("Position did not increase by one for IncPos")
	}
}

func TestIncreaseValueOneStep(t *testing.T) {
	state := InitProgramState()
	initialValue := state.Value()

	IncVal(&state)

	if state.Value() != initialValue+1 {
		t.Fatal("Value did not increase by one for IncVal")
	}
}

func TestIncreaseValueWith(t *testing.T) {
	state := InitProgramState()
	initialValue := state.Value()
	change := 42

	IncValWith(&state, change)

	if state.Value() != initialValue+change {
		t.Fatalf("Value should be %d but was %d", initialValue+change, state.Value())
	}
}

func TestDecreaseValueOneStep(t *testing.T) {
	state := InitProgramState()
	initialValue := state.Value()

	DecVal(&state)

	if state.Value() != initialValue-1 {
		t.Fatal("Value did not increase by one for IncVal")
	}
}

func TestDecreaseValueWith(t *testing.T) {
	state := InitProgramState()
	initialValue := state.Value()
	change := 42

	DecValWith(&state, change)

	if state.Value() != initialValue-change {
		t.Fatalf("Value should be %d but was %d", initialValue-change, state.Value())
	}
}

func TestInitIfJump(t *testing.T) {
	state := InitProgramState()
	state.programCounter = 0
	state.data[state.pos] = 0
	jumpLocation := 50

	InitIf(&state, jumpLocation)

	if state.GetProgramCounter() != jumpLocation {
		t.Fatalf("InitIf should have jumped to %d but program counter is %d", jumpLocation, state.GetProgramCounter())
	}
}

func TestInitIfNoJump(t *testing.T) {
	state := InitProgramState()
	initialPc := 0
	state.programCounter = initialPc
	state.data[state.pos] = 123

	InitIf(&state, 50)

	if state.GetProgramCounter() != initialPc {
		t.Fatalf("InitIf should have jumped to %d but program counter is %d", initialPc, state.GetProgramCounter())
	}
}

func TestEndIfJump(t *testing.T) {
	state := InitProgramState()
	state.programCounter = 0
	jumpLocation := 50
	state.data[state.pos] = 123

	EndIf(&state, jumpLocation)

	if state.GetProgramCounter() != jumpLocation {
		t.Fatalf("EndIf should have jumped to %d but program counter is %d", jumpLocation, state.GetProgramCounter())
	}
}

func TestEndIfNoJump(t *testing.T) {
	state := InitProgramState()
	initialPc := 0
	state.programCounter = initialPc
	state.data[state.pos] = 0

	EndIf(&state, 50)

	if state.GetProgramCounter() != initialPc {
		t.Fatalf("EndIf should have jumped to %d but program counter is %d", initialPc, state.GetProgramCounter())
	}
}

func TestReset(t *testing.T) {
	state := InitProgramState()
	state.data[state.pos] = 50

	Reset(&state)

	if state.Value() != 0 {
		t.Fatalf("Reset should have set value to 0 but set it to %d", state.Value())
	}
}

func TestResetAndStep(t *testing.T) {
	state := InitProgramState()
	initialPos := state.pos
	state.data[initialPos] = 50
	state.data[initialPos+1] = 60

	ResetAndStep(&state)

	if state.data[initialPos] != 0 {
		t.Fatalf("ResetAndStep should have set value to 0 but set it to %d", state.data[initialPos])
	}

	if state.pos != initialPos+1 {
		t.Fatal("ResetAndStep should have stepped")
	}
}
