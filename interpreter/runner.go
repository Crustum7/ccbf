package interpreter

import "martinjonson.com/ccbf/instructions"

func runAll(programState *instructions.ProgramState, statements string) {
	for i := 0; i < len(statements); i = programState.GetProgramCounter() {
		command := string(statements[i])

		switch command {
		case ">":
			instructions.IncPos(programState)
		case "<":
			instructions.DecPos(programState)
		case "+":
			instructions.IncVal(programState)
		case "-":
			instructions.DecVal(programState)
		case ".":
			instructions.CharOut(programState)
		case ",":
			instructions.CharIn(programState)
		case "[":
			jumpLoc := FindClosingBracket(statements, i)
			instructions.InitIf(programState, jumpLoc)
		case "]":
			jumpLoc := FindOpeningBracket(statements, i) - 1
			instructions.EndIf(programState, jumpLoc)
		}

		programState.IncrementProgramCounter()
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
