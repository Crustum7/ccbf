package interpreter

import "fmt"

func incPos(state *ProgramState) {
	if state.ignore > 0 {
		return
	}

	state.pos++
	state.AdjustCapacity()
}

func decPos(state *ProgramState) {
	if state.ignore > 0 {
		return
	}

	state.pos--
	if state.pos < 0 {
		panic("Negative data pointer error")
	}
}

func incVal(state *ProgramState) {
	if state.ignore > 0 {
		return
	}

	state.data[state.pos]++
}

func decVal(state *ProgramState) {
	if state.ignore > 0 {
		return
	}

	state.data[state.pos]--
}

func charOut(state *ProgramState) {
	if state.ignore > 0 {
		return
	}

	fmt.Print(string(state.Value()))
}

func charIn(state *ProgramState) {
	if state.ignore > 0 {
		return
	}

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		panic("Expected integer input")
	}

	state.data[state.pos] = i
}

/*
If-statements are stack based.

An if-statement is entered if the current pointer value is 0.
If entered, the current programCounter - 1 is pushed to the stack. This allows the end if to
jump back to before the if-statement so the if-statement is executed again.

If not entered, ignore counter is increased and all commands are ignored until the next ] is found.
If there are more if-statements before, they are also ignored using the ignore counter.
*/
func initIf(state *ProgramState) {
	if state.ignore > 0 {
		state.ignore++
		return
	}

	if state.Value() == 0 {
		state.ignore++
		return
	}

	state.NewIf(state.programCounter - 1)
}

func endIf(state *ProgramState) {
	if state.ignore > 0 {
		state.ignore--
		return
	}

	if state.Value() == 0 {
		state.PopIf()
		return
	}

	state.programCounter = state.PopIf()
}
