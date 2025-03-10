package compiler

import (
	"slices"

	"martinjonson.com/ccbf/internal/operations"
	"martinjonson.com/ccbf/internal/utils"
)

type Compiler struct {
	data      []byte
	jumpStack utils.Stack[int]
	parser    ProgramParser
}

type Command struct {
	operation     operations.Operation
	parsedCommand ParsedCommand
	opPos         int
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
	commandParser := InitCommandParser(patterns, ops)
	compiler.parser = InitProgramParser(program, commandParser)

	return compiler
}

func (compiler *Compiler) compile() {
	for compiler.parser.hasNext() {
		parsedCommand := compiler.parser.next()
		if parsedCommand.Pattern == "" {
			continue
		}

		compiler.parser.skipRepetitions(len(parsedCommand.Match))
		compiler.handleOperation(*parsedCommand.Operation, parsedCommand)
	}
}

func (compiler *Compiler) handleOperation(operation operations.Operation, parsedCommand ParsedCommand) {
	opPos := len(compiler.data)
	compiler.allocateOperationSpace(operation)

	command := Command{operation: operation, parsedCommand: parsedCommand, opPos: opPos}
	compiler.matchPattern(operation.GetPattern(), command)
}

func (compiler *Compiler) allocateOperationSpace(operation operations.Operation) {
	compiler.data = append(compiler.data, operation.GetOpCode())
	compiler.data = append(compiler.data, slices.Repeat([]byte{0}, operation.GetParameterByteCount())...)
}

func (compiler *Compiler) matchPattern(pattern string, command Command) {
	switch pattern {
	case `\[`:
		compiler.startLoop(command)
	case `\]`:
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
	addedBytes := []byte{}
	if len(command.parsedCommand.Groups) > 0 {
		addedBytes = append(addedBytes, byte(len(command.parsedCommand.Groups[0])))
	}
	compiler.assignParameterBytes(command.opPos, addedBytes)
}

func (compiler *Compiler) assignParameterBytes(opPos int, from []byte) {
	to := compiler.data[opPos+1:]
	for i := range len(from) {
		to[i] = from[i]
	}
}
