package interpreter

import (
	"martinjonson.com/ccbf/instructions"
)

func RunProgram(program string) {
	programState := instructions.InitProgramState()

	runAll(&programState, program)
}

func runAll(state *instructions.ProgramState, statements string) {
	for i := 0; i < len(statements); i = state.GetProgramCounter() {
		runCommand(state, statements, i)
		state.IncrementProgramCounter()
	}
}

func runCommand(state *instructions.ProgramState, statements string, statementIndex int) {
	command := string(statements[statementIndex])

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
		jumpLoc := findClosingBracket(statements, statementIndex)
		instructions.InitIf(state, jumpLoc)
	case "]":
		jumpLoc := findOpeningBracket(statements, statementIndex) - 1
		instructions.EndIf(state, jumpLoc)
	}
}

func findOpeningBracket(statements string, start int) int {
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

func findClosingBracket(statements string, start int) int {
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
