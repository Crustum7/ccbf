package interpreter

import (
	"os"

	"martinjonson.com/ccbf/instructions"
)

func RunProgram(commands string) {
	program := instructions.InitProgram(os.Stdin, os.Stdout)

	runAll(&program, commands)
}

func runAll(program *instructions.Program, statements string) {
	pc := program.GetProgramCounter()
	for i := 0; i < len(statements); i = pc.Get() {
		runCommand(program, statements, i)
		pc.Increment()
	}
}

func runCommand(program *instructions.Program, statements string, statementIndex int) {
	command := string(statements[statementIndex])

	switch command {
	case ">":
		program.IncPosWith(1)
	case "<":
		program.DecPosWith(1)
	case "+":
		program.IncValWith(1)
	case "-":
		program.DecValWith(1)
	case ".":
		program.CharOut()
	case ",":
		program.CharIn()
	case "[":
		jumpLoc := findClosingBracket(statements, statementIndex)
		program.InitIf(jumpLoc)
	case "]":
		jumpLoc := findOpeningBracket(statements, statementIndex) - 1
		program.EndIf(jumpLoc)
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
