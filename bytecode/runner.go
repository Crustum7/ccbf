package bytecode

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"martinjonson.com/ccbf/instructions"
)

func readInt8(buf *bytes.Reader) (int8, error) {
	var num int8
	err := binary.Read(buf, binary.BigEndian, &num)
	return num, err
}

func readInt32(buf *bytes.Reader) (int32, error) {
	var num int32
	err := binary.Read(buf, binary.BigEndian, &num)
	return num, err
}

func btoi(data []byte, numberOfBytes int) (int, error) {
	buf := bytes.NewReader(data)
	var num int
	var err error

	switch numberOfBytes {
	case 1:
		num8, err8 := readInt8(buf)
		num = int(num8)
		err = err8
	case 4:
		num32, err32 := readInt32(buf)
		num = int(num32)
		err = err32
	default:
		panic(fmt.Sprintf("Number of bytes of %d not supported", numberOfBytes))
	}

	return num, err
}

func parameter(data []byte, opLoc int, size int) []byte {
	offset := opLoc + 1
	return data[offset : offset+size]
}

func runAll(state *instructions.ProgramState, bytes []byte) {
	for i := 0; i < len(bytes); i = state.GetProgramCounter() {
		opCode := bytes[i]
		operation := OperationForOpCode(opCode)
		if operation == nil {
			panic(fmt.Sprintf("Incorrect bytefile parse for op code %b", opCode))
		}

		parameterBytes := parameter(bytes, i, operation.numberOfParameterBytes)
		state.IncreaseProgramCounter(operation.numberOfParameterBytes)

		switch opCode {
		case 1:
			instructions.IncPos(state)
		case 2:
			instructions.DecPos(state)
		case 3:
			instructions.IncVal(state)
		case 4:
			instructions.DecVal(state)
		case 5:
			instructions.CharOut(state)
		case 6:
			instructions.CharIn(state)
		case 7:
			jumpLoc, err := btoi(parameterBytes, operation.numberOfParameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.InitIf(state, int(jumpLoc))
		case 8:
			jumpLoc, err := btoi(parameterBytes, operation.numberOfParameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.EndIf(state, int(jumpLoc))
		case 9:
			repetitions, err := btoi(parameterBytes, operation.numberOfParameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.IncValWith(state, int(repetitions))
		case 10:
			repetitions, err := btoi(parameterBytes, operation.numberOfParameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.DecValWith(state, int(repetitions))
		case 11:
			repetitions, err := btoi(parameterBytes, operation.numberOfParameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.IncPosWith(state, int(repetitions))
		case 12:
			repetitions, err := btoi(parameterBytes, operation.numberOfParameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.DecPosWith(state, int(repetitions))
		}

		state.IncrementProgramCounter()
	}
}

func RunBytecode(bytes []byte) {
	state := instructions.InitProgramState()

	runAll(&state, bytes)
}
