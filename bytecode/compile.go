package bytecode

import (
	"bytes"
	"encoding/binary"
	"slices"

	"martinjonson.com/ccbf/byteoperation"
)

// Compiler, ProgramParser, Command
// Compiler has a program parser that delivers next command and bytes
// Compiler creates a Command that has everything needed for add bytes to data

type Compiler struct {
	data      []byte
	jumpStack Stack[int]
	parser    ProgramParser
}

func CompileProgram(program string, outFileName string) []byte {
	compiler := Compiler{}
	compiler.data = make([]byte, 0)
	compiler.jumpStack = InitStack[int]()
	// TODO: Get from list of operations
	commandParser := InitCommandParser([]string{">", "<", "+", "-", ",", ".", "[", "]", "[-]>", "[-]"})
	compiler.parser = InitProgramParser(program, commandParser)

	compiler.compile()

	return compiler.data
}

func (compiler *Compiler) compile() {
	for compiler.parser.hasNext() {
		command, repetitions := compiler.parser.next()
		operation := OperationForPattern(command, repetitions > 1)
		if operation == nil {
			continue
		}

		addedBytes, jumpLen := getBytesAndJump(command, repetitions)

		if jumpLen > 0 {
			compiler.data = append(compiler.data, addedBytes...)
			compiler.parser.skipRepetitions(jumpLen)
			continue
		}

		opPos := len(compiler.data)
		compiler.data = append(compiler.data, operation.opCode)
		compiler.data = append(compiler.data, slices.Repeat([]byte{0}, operation.numberOfParameterBytes)...)

		switch command {
		case "[":
			compiler.jumpStack.Push(opPos)
		case "]":
			startOpPos := compiler.jumpStack.Pop()

			toAddress, err := itob(int32(startOpPos + operation.numberOfParameterBytes))
			if err != nil {
				panic("Could not parse jump address to byte slice")
			}
			assignBytes(parameterBytesForOperation(compiler.data, opPos, *operation), toAddress)

			backAddress, err := itob(int32(opPos + operation.numberOfParameterBytes))
			if err != nil {
				panic("Could not parse jump address to byte slice")
			}

			assignBytes(parameterBytesForOperation(compiler.data, startOpPos, *operation), backAddress)
		}
	}
}

func getBytesAndJump(command string, repetitions int) ([]byte, int) {
	switch command {
	case ">":
		return byteoperation.RightMove(repetitions)
	case "<":
		return byteoperation.LeftMove(repetitions)
	case "+":
		return byteoperation.Add(repetitions)
	case "-":
		return byteoperation.Sub(repetitions)
	case ".":
		return byteoperation.Print(repetitions)
	case ",":
		return byteoperation.Input(repetitions)
	case "[-]>":
		return byteoperation.ResetAndStep(repetitions)
	case "[-]":
		return byteoperation.Reset(repetitions)
	default:
		return []byte{}, 0
	}
}

// TODO: Refactor
func itob(value int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		return nil, err
	}
	if len(buf.Bytes()) != 4 {
		panic("Wrong length of byte slice produced by itob()")
	}
	return buf.Bytes(), nil
}

func assignBytes(to []byte, from []byte) {
	if len(from) != len(to) {
		panic("Number of bytes do not match")
	}

	for i := range len(to) {
		to[i] = from[i]
	}
}

func parameterBytesForOperation(data []byte, opPos int, operation Operation) []byte {
	offset := opPos + 1
	return data[offset : offset+operation.numberOfParameterBytes]
}
