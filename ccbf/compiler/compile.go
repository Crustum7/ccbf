package compiler

import (
	"slices"

	"martinjonson.com/ccbf/ccbf/operations"
	"martinjonson.com/ccbf/ccbf/utils"
)

type Compiler struct {
	data      []byte
	jumpStack utils.Stack[int]
	parser    ProgramParser
}

type Command struct {
	operation   operations.Operation
	repetitions int
	opPos       int
}

func CompileProgram(program string) []byte {
	patterns := operations.OperationPatterns()
	ops := operations.GetOperations()
	compiler := initCompiler(program, patterns, ops)
	compiler.compile()

	return compiler.data
}

func initCompiler(program string, patterns []string, ops []operations.Operation) Compiler {
	compiler := Compiler{}
	compiler.data = make([]byte, 0)
	compiler.jumpStack = utils.InitStack[int]()
	commandParser, _ := InitCommandParser2(patterns, ops)
	compiler.parser = InitProgramParser(program, commandParser)

	return compiler
}

func (compiler *Compiler) compile() {
	for compiler.parser.hasNext() {
		str, operation := compiler.parser.next()
		if operation == nil {
			continue
		}

		compiler.handleOperation(*operation)
	}
}

func subpatternRepetitions(pattern string, str string) []int {
	// Some operations behave differently based on number of subpattern repetitions
	// \\++ for example: we need to find how many plus are found in a row
	// \\[->+\\+<+\\]: we need to know how many steps the value is moved
	// This is probably more of a regex problem and I should look up regex magic

	// pattern + program -> string representation + operation + number of repetitions per "interesting" value

	return []int{}
}

func (compiler *Compiler) handleOperation(operation operations.Operation, repetitions int) {
	opPos := len(compiler.data)
	compiler.allocateOperationSpace(operation)

	command := Command{operation: operation, repetitions: repetitions, opPos: opPos}
	compiler.matchPattern(operation.GetPattern(), command)
}

func (compiler *Compiler) allocateOperationSpace(operation operations.Operation) {
	compiler.data = append(compiler.data, operation.GetOpCode())
	compiler.data = append(compiler.data, slices.Repeat([]byte{0}, operation.GetParameterByteCount())...)
}

func (compiler *Compiler) matchPattern(pattern string, command Command) {
	switch pattern {
	case "[":
		compiler.startLoop(command)
	case "]":
		compiler.endLoop(command)
	default:
		compiler.generalOperation(command)
	}
}

func (compiler *Compiler) startLoop(command Command) {
	compiler.jumpStack.Push(command.opPos)
}

func (compiler *Compiler) endLoop(command Command) {
	startOpPos := compiler.jumpStack.Pop()
	byteCount := command.operation.GetParameterByteCount()

	toAddress := utils.Itob(int32(startOpPos + byteCount))
	compiler.assignParameterBytes(command.opPos, toAddress)

	backAddress := utils.Itob(int32(command.opPos + byteCount))
	compiler.assignParameterBytes(startOpPos, backAddress)
}

func (compiler *Compiler) generalOperation(command Command) {
	addedBytes := command.operation.StandardParameterBytes(command.repetitions)
	compiler.assignParameterBytes(command.opPos, addedBytes)

	jumpLen := command.operation.ParsedSymbols(command.repetitions)
	compiler.parser.skipRepetitions(jumpLen)
}

func (compiler *Compiler) assignParameterBytes(opPos int, from []byte) {
	to := compiler.data[opPos+1:]
	for i := range len(from) {
		to[i] = from[i]
	}
}
