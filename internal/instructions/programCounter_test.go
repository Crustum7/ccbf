package instructions_test

import (
	"testing"

	"martinjonson.com/ccbf/internal/instructions"
)

func TestProgramCounterInitialValue(t *testing.T) {
	pc := instructions.InitProgramCounter()

	if pc.Get() != 0 {
		t.Fatalf("Program counter should start at 0 and not %d", pc.Get())
	}
}

func TestProgramCounterIncrement(t *testing.T) {
	pc := instructions.InitProgramCounter()

	pc.Increment()

	if pc.Get() != 1 {
		t.Fatalf("Program counter should be 1 and not %d", pc.Get())
	}
}

func TestProgramCounterIncrementWith(t *testing.T) {
	pc := instructions.InitProgramCounter()

	pc.IncrementWith(5)

	if pc.Get() != 5 {
		t.Fatalf("Program counter should be 5 and not %d", pc.Get())
	}
}

func TestProgramCounterSet(t *testing.T) {
	pc := instructions.InitProgramCounter()

	pc.Set(42)

	if pc.Get() != 42 {
		t.Fatalf("Program counter should be 42 and not %d", pc.Get())
	}
}
