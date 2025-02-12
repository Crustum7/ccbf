package main

import (
	"os"

	"martinjonson.com/ccbf/compiler"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Not enought arguments")
	}

	inFileName := args[0]
	dat, err := os.ReadFile(inFileName)
	if err != nil {
		panic(err)
	}
	compiler.RunBytecode(dat)
}
