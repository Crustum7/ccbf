package instructions

type ProgramState struct {
	pos  int
	data []int
}

func (state *ProgramState) getValue() int {
	return state.data[state.pos]
}

func (state *ProgramState) setValue(val int) {
	state.data[state.pos] = val
}

func (state *ProgramState) capacity() int {
	return len(state.data)
}

func (state *ProgramState) adjustCapacity() {
	for state.pos >= state.capacity() {
		newData := make([]int, state.capacity())
		for i := range newData {
			newData[i] = 0
		}
		state.data = append(state.data, newData...)
	}
}

/*
The initial program state contains the following:

instruction pointer index starting at 0
32 data cells set to 0
*/
func initProgramState(initialCapacity int) ProgramState {
	var state ProgramState

	state.pos = 0
	state.data = make([]int, initialCapacity)
	for i := range state.data {
		state.data[i] = 0
	}

	return state
}
