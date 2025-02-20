package main

import (
	"os"

	"martinjonson.com/ccbf/compiler"
	"martinjonson.com/ccbf/virtual"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No file given")
	}

	processFiles(args)
}

func processFiles(fileNames []string) {
	for _, fileName := range fileNames {
		processFile(fileName)
	}
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
	virtual.RunBytecode(compiledCode, os.Stdin, os.Stdout)
}
