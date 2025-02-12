package main

import (
	"flag"
	"os"

	"martinjonson.com/ccbf/compiler"
	"martinjonson.com/ccbf/interpreter"
	"martinjonson.com/ccbf/virtual"
)

func main() {
	bytecodePtr := flag.Bool("bytecode", false, "Compile to bytecode before running code")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		panic("No file given")
	}

	for _, fileName := range args {
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		program := string(data)

		if *bytecodePtr {
			compiledCode := compiler.CompileProgram(program)
			virtual.RunBytecode(compiledCode)
		} else {
			interpreter.RunProgram(program)
		}
	}
}
