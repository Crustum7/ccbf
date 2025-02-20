package interpreter

import (
	"testing"

	"martinjonson.com/ccbf/test"
)

func TestFindOpeningBracket(t *testing.T) {
	statements := "abcga[v9jvavv,k]ocvahckdhrea"
	openingPos := 5
	closingPos := 15

	pos := findOpeningBracket(statements, closingPos)

	if pos != openingPos {
		t.Fatalf("Opening position should be %d, not %d", openingPos, pos)
	}
}

func TestFindOpeningBracketEdge(t *testing.T) {
	statements := "[]"

	pos := findOpeningBracket(statements, 1)

	if pos != 0 {
		t.Fatalf("Opening position should be 0, not %d", pos)
	}
}

func TestFindOpeningBracketPanic(t *testing.T) {
	statements := "abcga]v9jvavv,k]ocvahckdhrea"
	closingPos := 15

	test.ShouldPanic(t, func() { findOpeningBracket(statements, closingPos) })
}

func TestFindClosingBracket(t *testing.T) {
	statements := "abcga[v9jvavv,k]ocvahckdhrea"
	openingPos := 5
	closingPos := 15

	pos := findClosingBracket(statements, openingPos)

	if pos != closingPos {
		t.Fatalf("Closing position should be %d, not %d", closingPos, pos)
	}
}

func TestFindClosingBracketEdge(t *testing.T) {
	statements := "[]"

	pos := findClosingBracket(statements, 0)

	if pos != 1 {
		t.Fatalf("Opening position should be 1, not %d", pos)
	}
}

func TestFindClosingBracketPanic(t *testing.T) {
	statements := "abcga[v9jvavv,k[ocvahckdhrea"
	openingPos := 5

	test.ShouldPanic(t, func() { findClosingBracket(statements, openingPos) })
}

func ExampleRunProgram() {
	program := "This is a test Brainf*ck script written for Coding Challenges!++++++++++[>+>+++>+++++++>++++++++++<<<<-]>>>++.>+.+++++++..+++.<<++++++++++++++.------------.>-----.>.-----------.+++++.+++++.-------.<<.>.>+.-------.+++++++++++..-------.+++++++++.-------.--.++++++++++++++. What does it do?"

	RunProgram(program)
	// Output: Hello, Coding Challenges
}
