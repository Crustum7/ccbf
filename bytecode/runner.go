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
	switch OpCode(opCode) {
	case OneRightStep:
		instructions.IncPos(state)
	case OneLeftStep:
		instructions.DecPos(state)
	case IncrementOne:
		instructions.IncVal(state)
	case DecrementOne:
		instructions.DecVal(state)
	case PrintChar:
		instructions.CharOut(state)
	case InputChar:
		instructions.CharIn(state)
	case StartLoop:
		jumpLoc := btoi(parameterBytes)

		instructions.InitIf(state, int(jumpLoc))
	case EndLoop:
		jumpLoc := btoi(parameterBytes)

		instructions.EndIf(state, int(jumpLoc))
	case IncrementMultiple:
		repetitions := btoi(parameterBytes)

		instructions.IncValWith(state, int(repetitions))
	case DecrementMultiple:
		repetitions := btoi(parameterBytes)

		instructions.DecValWith(state, int(repetitions))
	case MultipleRightStep:
		repetitions := btoi(parameterBytes)

		instructions.IncPosWith(state, int(repetitions))
	case MultipleLeftStep:
		repetitions := btoi(parameterBytes)

		instructions.DecPosWith(state, int(repetitions))
	case ResetAndStep:
		instructions.ResetAndStep(state)
	case Reset:
		instructions.Reset(state)
	}
}
