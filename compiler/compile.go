package compiler

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

func CompileProgram(program string, patterns []string) []byte {
	compiler := initCompiler(program, patterns)
	compiler.compile()

	return compiler.data
}

func initCompiler(program string, patterns []string) Compiler {
	compiler := Compiler{}
	compiler.data = make([]byte, 0)
	compiler.jumpStack = utils.InitStack[int]()
	commandParser := InitCommandParser(patterns)
	compiler.parser = InitProgramParser(program, commandParser)

	return compiler
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
