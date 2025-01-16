package bytecode

import (
	"bytes"
	"encoding/binary"

	"martinjonson.com/ccbf/instructions"
)

const addressSize = 4

func btoi(data []byte) (int32, error) {
	buf := bytes.NewReader(data)

	var num int32
	err := binary.Read(buf, binary.BigEndian, &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func addressParameter(data []byte, opLoc int) []byte {
	offset := opLoc + 1
	return data[offset : offset+addressSize]
}

func runAll(state *instructions.ProgramState, bytes []byte) {
	for i := 0; i < len(bytes); i = state.GetProgramCounter() {
		command := bytes[i]
		// op := []string{">", "<", "+", "-", ".", ",", "[", "]"}
		// fmt.Print(op[command])

		switch command {
		case 0:
			instructions.IncPos(state)
		case 1:
			instructions.DecPos(state)
		case 2:
			instructions.IncVal(state)
		case 3:
			instructions.DecVal(state)
		case 4:
			instructions.CharOut(state)
		case 5:
			instructions.CharIn(state)
		case 6:
			state.IncreaseProgramCounter(addressSize)
			jumpLoc, err := btoi(addressParameter(bytes, i))
			if err != nil {
				panic(err)
			}
			instructions.InitIf(state, int(jumpLoc))
		case 7:
			state.IncreaseProgramCounter(addressSize)
			jumpLoc, err := btoi(addressParameter(bytes, i))
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
