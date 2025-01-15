package main

import (
	"fmt"
	"os"
)

type ProgramState struct {
	pos            int
	data           []int
	ignore         int
	stack          []int
	programCounter int
}

func (state *ProgramState) SetProgramCounter(value int) {
	state.programCounter = value
}

func (state *ProgramState) Value() int {
	return state.data[state.pos]
}

func (state *ProgramState) Capacity() int {
	return len(state.data)
}

func (state *ProgramState) NewIf(jumpLocation int) {
	state.stack = append(state.stack, jumpLocation)
}

func (state *ProgramState) PopIf() int {
	val := state.stack[len(state.stack)-1]
	state.stack = state.stack[:len(state.stack)-1]
	return val
}

/*
*

	The initial program state contains the following:

	instruction pointer index starting at 0
	32 data spaces set to 0
	ignore counter set to 0
	stack of string if-statements
*/
func initProgramState() ProgramState {
	var state ProgramState

	state.pos = 0
	state.data = make([]int, 32)
	for i := range state.data {
		state.data[i] = 0
	}
	state.ignore = 0
	state.stack = make([]int, 0)
	state.programCounter = 0

	return state
}

func incPos(state *ProgramState) {
	if state.ignore > 0 {
		return
	}

	state.pos++
	if state.pos >= state.Capacity() {
		newData := make([]int, state.Capacity())
		for i := range newData {
			newData[i] = 0
		}
		state.data = append(state.data, newData...)
	}
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
*

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

	// Evaluate stack
	newCounter := state.PopIf()
	state.SetProgramCounter(newCounter)
}

func runAll(programState *ProgramState, statements string) {
	for i := 0; i < len(statements); i = programState.programCounter + 1 {
		programState.SetProgramCounter(i)
		command := string(statements[i])
		// fmt.Println("Current instruction", command, ", program counter", i)

		switch command {
		case ">":
			incPos(programState)
		case "<":
			decPos(programState)
		case "+":
			incVal(programState)
		case "-":
			decVal(programState)
		case ".":
			charOut(programState)
		case ",":
			charIn(programState)
		case "[":
			initIf(programState)
		case "]":
			endIf(programState)
		}
	}
}

func runProgram(program string) {
	programState := initProgramState()

	runAll(&programState, program)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No file given")
	}

	for _, fileName := range args {
		// fmt.Println(fileName)
		dat, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		runProgram(string(dat))
	}
}
