package bytecode

type ProgramParser struct {
	program       string
	index         int
	commandParser CommandParser
}

func InitProgramParser(program string, commandParser CommandParser) ProgramParser {
	parser := ProgramParser{}
	parser.program = program
	parser.commandParser = commandParser
	parser.index = 0

	return parser
}

func (parser *ProgramParser) hasNext() bool {
	return parser.index < len(parser.program)
}

func (parser *ProgramParser) next() (string, int) {
	program := parser.program[parser.index:]
	parser.index++
	return parser.commandParser.FindPatternReapetions(program)
}

func (parser *ProgramParser) skipRepetitions(repetitions int) {
	parser.index += repetitions - 1
}
