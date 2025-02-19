package instructions

import "io"

type Program struct {
	state  ProgramState
	reader io.Reader
	writer io.Writer
}

const INITIALCAPACITY = 32

func InitProgram(reader io.Reader, writer io.Writer) Program {
	state := InitProgramState(INITIALCAPACITY)
	program := Program{state: state, reader: reader, writer: writer}

	return program
}

func (program *Program) GetProgramCounter() int {
	return program.state.GetProgramCounter()
}

func (program *Program) IncrementProgramCounter() {
	program.state.IncrementProgramCounter()
}

func (program *Program) IncrementProgramCounterWith(steps int) {
	program.state.IncrementProgramCounterWith(steps)
}
