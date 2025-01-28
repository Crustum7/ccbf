package bytecode

import (
	"fmt"

	"martinjonson.com/ccbf/instructions"
)

func btoi(data []byte, numberOfBytes int) int {
	var num int

	switch numberOfBytes {
	case 1:
		num = int(data[0])
	case 4:
		num = int(data[0])<<24 + int(data[1])<<16 + int(data[2])<<8 + int(data[3])
	default:
		panic(fmt.Sprintf("Number of bytes of %d not supported", numberOfBytes))
	}

	return num
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
			jumpLoc := btoi(parameterBytes, operation.numberOfParameterBytes)

			instructions.InitIf(state, int(jumpLoc))
		case 8:
			jumpLoc := btoi(parameterBytes, operation.numberOfParameterBytes)

			instructions.EndIf(state, int(jumpLoc))
		case 9:
			repetitions := btoi(parameterBytes, operation.numberOfParameterBytes)

			instructions.IncValWith(state, int(repetitions))
		case 10:
			repetitions := btoi(parameterBytes, operation.numberOfParameterBytes)

			instructions.DecValWith(state, int(repetitions))
		case 11:
			repetitions := btoi(parameterBytes, operation.numberOfParameterBytes)

			instructions.IncPosWith(state, int(repetitions))
		case 12:
			repetitions := btoi(parameterBytes, operation.numberOfParameterBytes)

			instructions.DecPosWith(state, int(repetitions))
		case 13:
			instructions.ResetAndStep(state)
		case 14:
			instructions.Reset(state)
		}

		state.IncrementProgramCounter()
	}
}

func RunBytecode(bytes []byte) {
	state := instructions.InitProgramState()

	runAll(&state, bytes)
}
