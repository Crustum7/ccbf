package interpreter

import (
	"fmt"
	"io"

	"martinjonson.com/ccbf/internal/instructions"
	"martinjonson.com/ccbf/internal/operations"
	"martinjonson.com/ccbf/internal/utils"
)

func RunBytecode(bytes []byte, reader io.Reader, writer io.Writer) {
	program := instructions.InitProgram(reader, writer)

	runAll(&program, bytes)
}

func runAll(program *instructions.Program, bytes []byte) {
	pc := program.GetProgramCounter()
	for i := 0; i < len(bytes); i = pc.Get() {
		opCode := bytes[i]
		operation := operations.OperationForOpCode(opCode)
		if operation == nil {
			panic(fmt.Sprintf("Incorrect bytefile parse for op code %b", opCode))
		}

		byteCount := operation.GetParameterByteCount()
		parameterBytes := parameter(bytes, i, byteCount)
		pc.IncrementWith(byteCount)
		matchInstruction(program, opCode, parameterBytes)

		pc.Increment()
	}
}

func parameter(data []byte, opLoc int, size int) []byte {
	offset := opLoc + 1
	return data[offset : offset+size]
}

func matchInstruction(program *instructions.Program, opCode byte, parameterBytes []byte) {
	switch operations.OpCode(opCode) {
	case operations.OneRightStep:
		program.IncPosWith(1)
	case operations.OneLeftStep:
		program.DecPosWith(1)
	case operations.IncrementOne:
		program.IncValWith(1)
	case operations.DecrementOne:
		program.DecValWith(1)
	case operations.PrintChar:
		program.CharOut()
	case operations.InputChar:
		program.CharIn()
	case operations.StartLoop:
		jumpLoc := utils.Btoi(parameterBytes)

		program.InitIf(int(jumpLoc))
	case operations.EndLoop:
		jumpLoc := utils.Btoi(parameterBytes)

		program.EndIf(int(jumpLoc))
	case operations.IncrementMultiple:
		repetitions := utils.Btoi(parameterBytes)

		program.IncValWith(int(repetitions))
	case operations.DecrementMultiple:
		repetitions := utils.Btoi(parameterBytes)

		program.DecValWith(int(repetitions))
	case operations.MultipleRightStep:
		repetitions := utils.Btoi(parameterBytes)

		program.IncPosWith(int(repetitions))
	case operations.MultipleLeftStep:
		repetitions := utils.Btoi(parameterBytes)

		program.DecPosWith(int(repetitions))
	case operations.ResetAndStep:
		program.ResetAndStep()
	case operations.Reset:
		program.Reset()
	case operations.MoveValueRight:
		steps := utils.Btoi(parameterBytes)

		program.MoveValueRight(steps)
	default:
		panic(fmt.Sprintf("Interpreter does not handle op code %b", opCode))
	}
}
