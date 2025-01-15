package main

import (
	"os"

	"martinjonson.com/ccbf/interpreter"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No file given")
	}

	for _, fileName := range args {
		dat, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		interpreter.RunProgram(string(dat))
	}
}
