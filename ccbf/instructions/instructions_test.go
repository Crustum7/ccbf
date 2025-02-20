package instructions

import (
	"bytes"
	"os"
	"testing"

	"martinjonson.com/ccbf/test"
)

func StdProg() Program {
	return InitProgram(os.Stdin, os.Stdout)
}

func TestNegativePositionShouldPanic(t *testing.T) {
	program := StdProg()
	test.ShouldPanic(t, func() { program.DecPosWith(1) })
}

func TestIncreaseValueWith(t *testing.T) {
	program := StdProg()
	initialValue := program.state.getValue()
	change := 42

	program.IncValWith(change)

	val := program.state.getValue()
	if val != initialValue+change {
		t.Fatalf("Value should be %d but was %d", initialValue+change, val)
	}
}

func TestDecreaseValueWith(t *testing.T) {
	program := StdProg()
	initialValue := program.state.getValue()
	change := 42

	program.DecValWith(change)

	val := program.state.getValue()
	if val != initialValue-change {
		t.Fatalf("Value should be %d but was %d", initialValue-change, val)
	}
}

func TestInitIfJump(t *testing.T) {
	program := StdProg()
	pc := program.GetProgramCounter()
	pc.Set(0)
	program.state.setValue(0)
	jumpLocation := 50

	program.InitIf(jumpLocation)

	pcVal := pc.Get()
	if pcVal != jumpLocation {
		t.Fatalf("InitIf should have jumped to %d but program counter is %d", jumpLocation, pcVal)
	}
}

func TestInitIfNoJump(t *testing.T) {
	program := StdProg()
	initialPc := 0
	pc := program.GetProgramCounter()
	pc.Set(initialPc)
	program.state.setValue(123)

	program.InitIf(50)

	pcVal := pc.Get()
	if pcVal != initialPc {
		t.Fatalf("InitIf should have jumped to %d but program counter is %d", initialPc, pcVal)
	}
}

func TestEndIfJump(t *testing.T) {
	program := StdProg()
	pc := program.GetProgramCounter()
	pc.Set(0)
	jumpLocation := 50
	program.state.setValue(123)

	program.EndIf(jumpLocation)

	pcVal := pc.Get()
	if pcVal != jumpLocation {
		t.Fatalf("EndIf should have jumped to %d but program counter is %d", jumpLocation, pcVal)
	}
}

func TestEndIfNoJump(t *testing.T) {
	program := StdProg()
	initialPc := 0
	pc := program.GetProgramCounter()
	pc.Set(initialPc)
	program.state.setValue(0)

	program.EndIf(50)

	pcVal := pc.Get()
	if pcVal != initialPc {
		t.Fatalf("EndIf should have jumped to %d but program counter is %d", initialPc, pcVal)
	}
}

func TestReset(t *testing.T) {
	program := StdProg()
	program.state.setValue(50)

	program.Reset()

	val := program.state.getValue()
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

func TestCharIn(t *testing.T) {
	expectedVal := 63
	input := "63"
	buffer := bytes.NewBufferString(input)
	program := InitProgram(buffer, os.Stdout)

	program.CharIn()

	val := program.state.getValue()
	if val != expectedVal {
		t.Fatalf("Expected CharIn to take %d but took %d", expectedVal, val)
	}
}

func TestCharInPanic(t *testing.T) {
	input := "A"
	buffer := bytes.NewBufferString(input)
	program := InitProgram(buffer, os.Stdout)

	test.ShouldPanic(t, func() { program.CharIn() })
}

func TestCharOut(t *testing.T) {
	var writer bytes.Buffer
	program := InitProgram(os.Stdin, &writer)
	program.state.setValue(65)
	expectedChar := "A"

	program.CharOut()

	val := writer.String()
	if val != expectedChar {
		t.Fatalf("Expected CharOut to output %s but got %s", expectedChar, val)
	}
}
