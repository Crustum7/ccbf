package compiler

import (
	"fmt"

	"martinjonson.com/ccbf/instructions"
	"martinjonson.com/ccbf/operations"
)

func RunBytecode(bytes []byte) {
	state := instructions.InitProgramState()

	runAll(&state, bytes)
}

func runAll(state *instructions.ProgramState, bytes []byte) {
	for i := 0; i < len(bytes); i = state.GetProgramCounter() {
		opCode := bytes[i]
		operation := operations.OperationForOpCode(opCode)
		if operation == nil {
			panic(fmt.Sprintf("Incorrect bytefile parse for op code %b", opCode))
		}

		byteCount := operation.GetParameterByteCount()
		parameterBytes := parameter(bytes, i, byteCount)
		state.IncreaseProgramCounter(byteCount)
		matchInstruction(state, opCode, parameterBytes)

		state.IncrementProgramCounter()
	}
}

func parameter(data []byte, opLoc int, size int) []byte {
	offset := opLoc + 1
	return data[offset : offset+size]
}

func matchInstruction(state *instructions.ProgramState, opCode byte, parameterBytes []byte) {
	switch operations.OpCode(opCode) {
	case operations.OneRightStep:
		instructions.IncPos(state)
	case operations.OneLeftStep:
		instructions.DecPos(state)
	case operations.IncrementOne:
		instructions.IncVal(state)
	case operations.DecrementOne:
		instructions.DecVal(state)
	case operations.PrintChar:
		instructions.CharOut(state)
	case operations.InputChar:
		instructions.CharIn(state)
	case operations.StartLoop:
		jumpLoc := btoi(parameterBytes)

		instructions.InitIf(state, int(jumpLoc))
	case operations.EndLoop:
		jumpLoc := btoi(parameterBytes)

		instructions.EndIf(state, int(jumpLoc))
	case operations.IncrementMultiple:
		repetitions := btoi(parameterBytes)

		instructions.IncValWith(state, int(repetitions))
	case operations.DecrementMultiple:
		repetitions := btoi(parameterBytes)

		instructions.DecValWith(state, int(repetitions))
	case operations.MultipleRightStep:
		repetitions := btoi(parameterBytes)

		instructions.IncPosWith(state, int(repetitions))
	case operations.MultipleLeftStep:
		repetitions := btoi(parameterBytes)

		instructions.DecPosWith(state, int(repetitions))
	case operations.ResetAndStep:
		instructions.ResetAndStep(state)
	case operations.Reset:
		instructions.Reset(state)
	}
}
