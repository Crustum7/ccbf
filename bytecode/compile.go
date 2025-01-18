package bytecode

import (
	"bytes"
	"encoding/binary"
	"os"
	"slices"
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

	for i := 0; i < len(program); i++ {
		command := string(program[i])
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
