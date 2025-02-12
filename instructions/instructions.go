package instructions

import "fmt"

func IncPos(state *ProgramState) {
	IncPosWith(state, 1)
}

func IncPosWith(state *ProgramState, change int) {
	state.pos += change
	state.AdjustCapacity()
}

func DecPos(state *ProgramState) {
	DecPosWith(state, 1)
}

func DecPosWith(state *ProgramState, change int) {
	state.pos -= change
	if state.pos < 0 {
		panic("Negative data pointer error")
	}
}

func IncVal(state *ProgramState) {
	state.data[state.pos]++
}

func IncValWith(state *ProgramState, change int) {
	state.data[state.pos] += change
}

func DecVal(state *ProgramState) {
	state.data[state.pos]--
}

func DecValWith(state *ProgramState, change int) {
	state.data[state.pos] -= change
}

func CharOut(state *ProgramState) {
	fmt.Printf("%c", state.Value())
}

func CharIn(state *ProgramState) {
	_, err := fmt.Scanf("%d", &state.data[state.pos])
	if err != nil {
		panic("Expected integer input")
	}
}

func InitIf(state *ProgramState, jumpLoc int) {
	if state.Value() == 0 {
		state.programCounter = jumpLoc
	}
}

func EndIf(state *ProgramState, jumpLoc int) {
	if state.Value() != 0 {
		state.programCounter = jumpLoc
	}
}

func ResetAndStep(state *ProgramState) {
	Reset(state)
	IncPos(state)
}

func Reset(state *ProgramState) {
	state.data[state.pos] = 0
}
