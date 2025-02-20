package instructions

import (
	"fmt"
	"io"
)

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

func (program *Program) write(char byte) {
	fmt.Fprintf(program.writer, "%c", char)
}

func (program *Program) read() int {
	val, err := fmt.Fscanf(program.reader, "%d")
	if err != nil {
		panic("Expected integer input")
	}
	return val
}
