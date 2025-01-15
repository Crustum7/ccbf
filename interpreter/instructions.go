package interpreter

import "fmt"

func incPos(state *ProgramState) {
	state.pos++
	state.AdjustCapacity()
}

func decPos(state *ProgramState) {
	state.pos--
	if state.pos < 0 {
		panic("Negative data pointer error")
	}
}

func incVal(state *ProgramState) {
	state.data[state.pos]++
}

func decVal(state *ProgramState) {
	state.data[state.pos]--
}

func charOut(state *ProgramState) {
	fmt.Printf("%c", state.Value())
}

func charIn(state *ProgramState) {
	_, err := fmt.Scanf("%d", &state.data[state.pos])
	if err != nil {
		panic("Expected integer input")
	}
}

func initIf(state *ProgramState, jumpLoc int) {
	if state.Value() == 0 {
		state.programCounter = jumpLoc
	}
}

func endIf(state *ProgramState, jumpLoc int) {
	if state.Value() != 0 {
		state.programCounter = jumpLoc
	}
}
