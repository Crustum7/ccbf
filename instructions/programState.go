package instructions

type ProgramState struct {
	pos  int
	data []int
}

func (state *ProgramState) Value() int {
	return state.data[state.pos]
}

func (state *ProgramState) Capacity() int {
	return len(state.data)
}

func (state *ProgramState) AdjustCapacity() {
	for state.pos >= state.Capacity() {
		newData := make([]int, state.Capacity())
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
func InitProgramState(initialCapacity int) ProgramState {
	var state ProgramState

	state.pos = 0
	state.data = make([]int, initialCapacity)
	for i := range state.data {
		state.data[i] = 0
	}

	return state
}
