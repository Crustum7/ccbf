package utils

import (
	"testing"

	"martinjonson.com/ccbf/internal/test"
)

func TestEmptyStack(t *testing.T) {
	stack := InitStack[string]()
	if stack.Size() != 0 {
		t.Fatal("Stack should initialize without elements")
	}
}

func TestPushPop(t *testing.T) {
	stack := InitStack[string]()
	value := "A"

	stack.Push(value)
	if stack.Size() != 1 {
		t.Fatal("Stack should have one element")
	}
	top := stack.Pop()
	if stack.Size() != 0 {
		t.Fatal("Stack should have 0 elements after Pop")
	}

	if top != value {
		t.Fatalf("Stack should have popped an \"%s\" but popped \"%s\"", value, top)
	}
}

func TestMoreElements(t *testing.T) {
	stack := InitStack[string]()

	stack.Push("A")
	stack.Push("B")
	stack.Push("C")
	if stack.Size() != 3 {
		t.Fatal("Stack should have three elements")
	}

	var top string
	top = stack.Pop()
	if top != "C" {
		t.Fatalf("Stack should have popped an \"C\" but popped \"%s\"", top)
	}

	top = stack.Pop()
	if top != "B" {
		t.Fatalf("Stack should have popped an \"B\" but popped \"%s\"", top)
	}

	top = stack.Pop()
	if top != "A" {
		t.Fatalf("Stack should have popped an \"A\" but popped \"%s\"", top)
	}
}

func TestCrashIfEmptyPop(t *testing.T) {
	stack := InitStack[string]()
	test.ShouldPanic(t, func() { stack.Pop() })
}
