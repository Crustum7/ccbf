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
			initIf(programState)
		case "]":
			endIf(programState)
		}
	}
}

func RunProgram(program string) {
	programState := initProgramState()

	runAll(&programState, program)
}
