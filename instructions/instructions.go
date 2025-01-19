package instructions

import "fmt"

func IncPos(state *ProgramState) {
	state.pos++
	state.AdjustCapacity()
}

func DecPos(state *ProgramState) {
	state.pos--
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
