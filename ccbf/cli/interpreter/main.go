package main

import (
	"os"

	"martinjonson.com/ccbf/ccbf/compiler"
	"martinjonson.com/ccbf/ccbf/interpreter"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No file given")
	}

	processFile(args[0])
}

func processFile(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	code := string(data)
	runProgram(code)
}

func runProgram(code string) {
	compiledCode := compiler.CompileProgram(code)
	interpreter.RunBytecode(compiledCode, os.Stdin, os.Stdout)
}
