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
	runCompiler(inFileName, outFileName)
}

func runCompiler(inFileName string, outFileName string) {
	code := read(inFileName)
	bytecodeProgram := bytecode.CompileProgram(code, outFileName)
	dump(bytecodeProgram, outFileName)
}

func read(fileName string) string {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func dump(bytes []byte, fileName string) {
	os.WriteFile(fileName, bytes, 0777)
}
