package bytecode

import (
	"fmt"

	"martinjonson.com/ccbf/instructions"
)

func RunBytecode(bytes []byte) {
	state := instructions.InitProgramState()

	runAll(&state, bytes)
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
		matchInstruction(state, opCode, parameterBytes)

		state.IncrementProgramCounter()
	}
}

func parameter(data []byte, opLoc int, size int) []byte {
	offset := opLoc + 1
	return data[offset : offset+size]
}

func matchInstruction(state *instructions.ProgramState, opCode byte, parameterBytes []byte) {
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
		jumpLoc := btoi(parameterBytes)

		instructions.InitIf(state, int(jumpLoc))
	case 8:
		jumpLoc := btoi(parameterBytes)

		instructions.EndIf(state, int(jumpLoc))
	case 9:
		repetitions := btoi(parameterBytes)

		instructions.IncValWith(state, int(repetitions))
	case 10:
		repetitions := btoi(parameterBytes)

		instructions.DecValWith(state, int(repetitions))
	case 11:
		repetitions := btoi(parameterBytes)

		instructions.IncPosWith(state, int(repetitions))
	case 12:
		repetitions := btoi(parameterBytes)

		instructions.DecPosWith(state, int(repetitions))
	case 13:
		instructions.ResetAndStep(state)
	case 14:
		instructions.Reset(state)
	}
}
