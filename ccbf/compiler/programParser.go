package compiler

import "martinjonson.com/ccbf/ccbf/operations"

type ProgramParser struct {
	program       string
	index         int
	commandParser CommandParser2
}

func InitProgramParser(program string, commandParser CommandParser2) ProgramParser {
	parser := ProgramParser{}
	parser.program = program
	parser.commandParser = commandParser
	parser.index = 0

	return parser
}

func (parser *ProgramParser) hasNext() bool {
	return parser.index < len(parser.program)
}

func (parser *ProgramParser) next() (string, *operations.Operation) {
	program := parser.program[parser.index:]
	parser.index++
	return parser.commandParser.FindLongest(program)
}

func (parser *ProgramParser) skipRepetitions(repetitions int) {
	parser.index += repetitions - 1
}
