package interpreter_test

import (
	"os"

	"martinjonson.com/ccbf/ccbf/interpreter"
)

func ExampleRunBytecode() {
	program := []byte{9, 10, 7, 0, 0, 0, 26, 1, 3, 1, 9, 3, 1, 9, 7, 1, 9, 10, 12, 3, 2, 4, 8, 0, 0, 0, 6, 11, 3, 9, 2, 5, 1, 3, 5, 9, 7, 5, 5, 9, 3, 5, 12, 2, 9, 12, 9, 2, 5, 10, 12, 5, 1, 10, 5, 5, 1, 5, 10, 11, 5, 9, 2, 9, 3, 5, 9, 5, 5, 10, 7, 5, 12, 2, 5, 1, 5, 1, 3, 5, 10, 7, 5, 9, 5, 9, 6, 5, 5, 10, 7, 5, 9, 9, 5, 10, 7, 5, 10, 2, 5, 9, 2, 9, 12, 5}

	interpreter.RunBytecode(program, os.Stdin, os.Stdout)
	// Output: Hello, Coding Challenges
}
