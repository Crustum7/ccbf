package main

import (
	"os"

	"martinjonson.com/ccbf/bytecode"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Not enought arguments")
	}

	inFileName := args[0]
	outFileName := args[1]
	dat, err := os.ReadFile(inFileName)
	if err != nil {
		panic(err)
	}
	bytecode.CompileProgram(string(dat), outFileName)
}
