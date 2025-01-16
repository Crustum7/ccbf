package interpreter

import (
	"martinjonson.com/ccbf/instructions"
)

func runAll(state *instructions.ProgramState, statements string) {
	for i := 0; i < len(statements); i = state.GetProgramCounter() {
		command := string(statements[i])
		// fmt.Print(command)

		switch command {
		case ">":
			instructions.IncPos(state)
		case "<":
			instructions.DecPos(state)
		case "+":
			instructions.IncVal(state)
		case "-":
			instructions.DecVal(state)
		case ".":
			instructions.CharOut(state)
		case ",":
			instructions.CharIn(state)
		case "[":
			jumpLoc := FindClosingBracket(statements, i)
			instructions.InitIf(state, jumpLoc)
		case "]":
			jumpLoc := FindOpeningBracket(statements, i) - 1
			instructions.EndIf(state, jumpLoc)
		}

		state.IncrementProgramCounter()
	}
}

func FindOpeningBracket(statements string, start int) int {
	counter := 0

	for i := start; i >= 0; i-- {
		switch string(statements[i]) {
		case "]":
			counter++
		case "[":
			counter--
		}
		if counter == 0 {
			return i
		}
	}
	panic("Could not find opening bracket")
}

func FindClosingBracket(statements string, start int) int {
	counter := 0

	for i := start; i < len(statements); i++ {
		switch string(statements[i]) {
		case "[":
			counter++
		case "]":
			counter--
		}
		if counter == 0 {
			return i
		}
	}
	panic("Could not find closing bracket")
}

func RunProgram(program string) {
	programState := instructions.InitProgramState()

	runAll(&programState, program)
}
