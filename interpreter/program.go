package interpreter

func runAll(programState *ProgramState, statements string) {
	for i := 0; i < len(statements); i = programState.programCounter + 1 {
		programState.programCounter = i
		command := string(statements[i])

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
			jumpLoc := findClosingBracket(statements, i)
			initIf(programState, jumpLoc)
		case "]":
			jumpLoc := findOpeningBracket(statements, i) - 1
			endIf(programState, jumpLoc)
		}
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

func RunProgram(program string) {
	programState := initProgramState()

	runAll(&programState, program)
}
