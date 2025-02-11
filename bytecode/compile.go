package bytecode

import (
	"slices"
)

type Compiler struct {
	data      []byte
	jumpStack Stack[int]
	parser    ProgramParser
}

func CompileProgram(program string, outFileName string) []byte {
	compiler := Compiler{}
	compiler.data = make([]byte, 0)
	compiler.jumpStack = InitStack[int]()
	// TODO: Get from list of operations
	commandParser := InitCommandParser([]string{">", "<", "+", "-", ",", ".", "[", "]", "[-]>", "[-]"})
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
	operation := OperationForPattern(command, repetitions > 1)
	if operation == nil {
		return
	}
	compiler.handleOperation(*operation, repetitions)
}

type Command struct {
	operation   Operation
	repetitions int
	opPos       int
}

func (compiler *Compiler) handleOperation(operation Operation, repetitions int) {
	opPos := len(compiler.data)
	compiler.allocateOperation(operation)
	command := Command{operation: operation, repetitions: repetitions, opPos: opPos}
	compiler.matchPattern(operation.pattern, command)
}

func (compiler *Compiler) allocateOperation(operation Operation) {
	compiler.data = append(compiler.data, operation.opCode)
	compiler.data = append(compiler.data, slices.Repeat([]byte{0}, operation.numberOfParameterBytes)...)
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
	parameterBytes := command.operation.numberOfParameterBytes

	toAddress := itob(int32(startOpPos + parameterBytes))
	compiler.assignBytes(command.opPos, command.operation, toAddress)

	backAddress := itob(int32(command.opPos + parameterBytes))
	compiler.assignBytes(startOpPos, command.operation, backAddress)
}

func (compiler *Compiler) generalOperation(command Command) {
	addedBytes, jumpLen := getBytesAndJump(command.operation, command.repetitions)

	compiler.assignBytes(command.opPos, command.operation, addedBytes)
	compiler.parser.skipRepetitions(jumpLen)
}

func (compiler *Compiler) assignBytes(opPos int, operation Operation, bytes []byte) {
	assignBytes(parameterBytesForOperation(compiler.data, opPos, operation), bytes)
}

func getBytesAndJump(operation Operation, repetitions int) ([]byte, int) {
	if operation.repeated {
		return []byte{byte(repetitions)}, len(operation.pattern) * repetitions
	}

	return []byte{}, len(operation.pattern)
}

func itob(value int32) []byte {
	return []byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)}
}

func assignBytes(to []byte, from []byte) {
	if len(from) != len(to) {
		panic("Number of bytes do not match")
	}

	for i := range len(to) {
		to[i] = from[i]
	}
}

func parameterBytesForOperation(data []byte, opPos int, operation Operation) []byte {
	offset := opPos + 1
	return data[offset : offset+operation.numberOfParameterBytes]
}
