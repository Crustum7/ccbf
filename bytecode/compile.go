package bytecode

import (
	"slices"

	"martinjonson.com/ccbf/operations"
	"martinjonson.com/ccbf/utils"
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
	compiler := Compiler{}
	compiler.data = make([]byte, 0)
	compiler.jumpStack = utils.InitStack[int]()
	patterns := operations.OperationPatterns()
	commandParser := InitCommandParser(patterns)
	compiler.parser = InitProgramParser(program, commandParser)

	compiler.compile()

	return compiler.data
}

func (compiler *Compiler) compile() {
	for compiler.parser.hasNext() {
		command, repetitions := compiler.parser.next()
		compiler.handleCommand(command, repetitions)
	}
}

func (compiler *Compiler) handleCommand(command string, repetitions int) {
	operation := operations.OperationForPattern(command, repetitions > 1)
	if operation == nil {
		return
	}
	compiler.handleOperation(*operation, repetitions)
}

func (compiler *Compiler) handleOperation(operation operations.Operation, repetitions int) {
	opPos := len(compiler.data)
	compiler.allocateOperation(operation)
	command := Command{operation: operation, repetitions: repetitions, opPos: opPos}
	compiler.matchPattern(operation.GetPattern(), command)
}

func (compiler *Compiler) allocateOperation(operation operations.Operation) {
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

	toAddress := itob(int32(startOpPos + byteCount))
	compiler.assignBytes(command.opPos, command.operation, toAddress)

	backAddress := itob(int32(command.opPos + byteCount))
	compiler.assignBytes(startOpPos, command.operation, backAddress)
}

func (compiler *Compiler) generalOperation(command Command) {
	addedBytes := command.operation.StandardParameterBytes(command.repetitions)
	compiler.assignBytes(command.opPos, command.operation, addedBytes)

	jumpLen := command.operation.ParsedSymbols(command.repetitions)
	compiler.parser.skipRepetitions(jumpLen)
}

func (compiler *Compiler) assignBytes(opPos int, operation operations.Operation, bytes []byte) {
	assignBytes(parameterBytesForOperation(compiler.data, opPos, operation), bytes)
}
