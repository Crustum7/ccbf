package bytecode

import (
	"bytes"
	"encoding/binary"
	"os"
	"slices"

	"martinjonson.com/ccbf/byteoperation"
)

func dump(bytes []byte, outFileName string) {
	os.WriteFile(outFileName, bytes, 0777)
}

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

func CompileProgram(program string, outFileName string) {
	data := make([]byte, 0)
	jumps := InitStack[int]()
	parser := InitCommandParser([]string{">", "<", "+", "-", ",", ".", "[", "]", "[-]>", "[-]"})

	for i := 0; i < len(program); {
		command, repetitions := parser.FindPatternReapetions(program[i:])
		addedBytes, jumpLen := getBytesAndJump(command, repetitions)

		if jumpLen > 0 {
			data = append(data, addedBytes...)
			i += jumpLen
			continue
		}

		i++

		operation := OperationForPattern(command)
		if operation == nil {
			continue
		}

		opPos := len(data)
		data = append(data, operation.opCode)
		data = append(data, slices.Repeat([]byte{0}, operation.numberOfParameterBytes)...)

		switch command {
		case "[":
			jumps.Push(opPos)
		case "]":
			startOpPos := jumps.Pop()

			toAddress, err := itob(int32(startOpPos + operation.numberOfParameterBytes))
			if err != nil {
				panic("Could not parse jump address to byte slice")
			}
			assignBytes(parameterBytesForOperation(data, opPos, *operation), toAddress)

			backAddress, err := itob(int32(opPos + operation.numberOfParameterBytes))
			if err != nil {
				panic("Could not parse jump address to byte slice")
			}

			assignBytes(parameterBytesForOperation(data, startOpPos, *operation), backAddress)
		}
	}

	dump(data, outFileName)
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
