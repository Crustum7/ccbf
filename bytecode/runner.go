package bytecode

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"martinjonson.com/ccbf/instructions"
)

func btoi(data []byte) (int32, error) {
	buf := bytes.NewReader(data)

	var num int32
	err := binary.Read(buf, binary.BigEndian, &num)
	if err != nil {
		return 0, err
	}
	return num, nil
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
			jumpLoc, err := btoi(parameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.InitIf(state, int(jumpLoc))
		case 8:
			jumpLoc, err := btoi(parameterBytes)
			if err != nil {
				panic(err)
			}
			instructions.EndIf(state, int(jumpLoc))
		}

		state.IncrementProgramCounter()
	}
}

func RunBytecode(bytes []byte) {
	state := instructions.InitProgramState()

	runAll(&state, bytes)
}
