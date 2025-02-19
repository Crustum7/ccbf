package instructions

import (
	"os"
	"testing"
)

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("Should have panicked")
}

func StdProg() Program {
	return InitProgram(os.Stdin, os.Stdout)
}

func TestNegativePositionShouldPanic(t *testing.T) {
	program := StdProg()
	shouldPanic(t, func() { program.DecPosWith(1) })
}

func TestIncreaseValueWith(t *testing.T) {
	program := StdProg()
	initialValue := program.state.Value()
	change := 42

	program.IncValWith(change)

	val := program.state.Value()
	if val != initialValue+change {
		t.Fatalf("Value should be %d but was %d", initialValue+change, val)
	}
}

func TestDecreaseValueWith(t *testing.T) {
	program := StdProg()
	initialValue := program.state.Value()
	change := 42

	program.DecValWith(change)

	val := program.state.Value()
	if val != initialValue-change {
		t.Fatalf("Value should be %d but was %d", initialValue-change, val)
	}
}

func TestInitIfJump(t *testing.T) {
	program := StdProg()
	program.state.programCounter = 0
	program.state.data[program.state.pos] = 0
	jumpLocation := 50

	program.InitIf(jumpLocation)

	pc := program.GetProgramCounter()
	if pc != jumpLocation {
		t.Fatalf("InitIf should have jumped to %d but program counter is %d", jumpLocation, pc)
	}
}

func TestInitIfNoJump(t *testing.T) {
	program := StdProg()
	initialPc := 0
	program.state.programCounter = initialPc
	program.state.data[program.state.pos] = 123

	program.InitIf(50)

	pc := program.GetProgramCounter()
	if pc != initialPc {
		t.Fatalf("InitIf should have jumped to %d but program counter is %d", initialPc, pc)
	}
}

func TestEndIfJump(t *testing.T) {
	program := StdProg()
	program.state.programCounter = 0
	jumpLocation := 50
	program.state.data[program.state.pos] = 123

	program.EndIf(jumpLocation)

	pc := program.GetProgramCounter()
	if pc != jumpLocation {
		t.Fatalf("EndIf should have jumped to %d but program counter is %d", jumpLocation, pc)
	}
}

func TestEndIfNoJump(t *testing.T) {
	program := StdProg()
	initialPc := 0
	program.state.programCounter = initialPc
	program.state.data[program.state.pos] = 0

	program.EndIf(50)

	pc := program.GetProgramCounter()
	if pc != initialPc {
		t.Fatalf("EndIf should have jumped to %d but program counter is %d", initialPc, pc)
	}
}

func TestReset(t *testing.T) {
	program := StdProg()
	program.state.data[program.state.pos] = 50

	program.Reset()

	val := program.state.Value()
	if val != 0 {
		t.Fatalf("Reset should have set value to 0 but set it to %d", val)
	}
}

func TestResetAndStep(t *testing.T) {
	program := StdProg()
	initialPos := program.state.pos
	program.state.data[initialPos] = 50
	program.state.data[initialPos+1] = 60

	program.ResetAndStep()

	initVal := program.state.data[initialPos]
	if initVal != 0 {
		t.Fatalf("ResetAndStep should have set value to 0 but set it to %d", initVal)
	}

	if program.state.pos != initialPos+1 {
		t.Fatal("ResetAndStep should have stepped")
	}
}
