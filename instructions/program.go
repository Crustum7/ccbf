package instructions

import "io"

type Program struct {
	state  ProgramState
	reader io.Reader
	writer io.Writer
	pc     ProgramCounter
}

const INITIALCAPACITY = 32

func InitProgram(reader io.Reader, writer io.Writer) Program {
	state := InitProgramState(INITIALCAPACITY)
	pc := InitProgramCounter()
	program := Program{state: state, reader: reader, writer: writer, pc: pc}

	return program
}

func (program *Program) GetProgramCounter() *ProgramCounter {
	return &program.pc
}
