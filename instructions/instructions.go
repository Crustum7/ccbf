package instructions

import (
	"fmt"
)

func (program *Program) IncPosWith(change int) {
	program.state.pos += change
	program.state.AdjustCapacity()
}

func (program *Program) DecPosWith(change int) {
	program.state.pos -= change
	if program.state.pos < 0 {
		panic("Negative data pointer error")
	}
}

func (program *Program) IncValWith(change int) {
	program.state.data[program.state.pos] += change
}

func (program *Program) DecValWith(change int) {
	program.state.data[program.state.pos] -= change
}

func (program *Program) CharOut() {
	fmt.Fprintf(program.writer, "%c", program.state.Value())
}

func (program *Program) CharIn() {
	_, err := fmt.Fscanf(program.reader, "%d", &program.state.data[program.state.pos])
	if err != nil {
		panic("Expected integer input")
	}
}

func (program *Program) InitIf(jumpLoc int) {
	if program.state.Value() == 0 {
		program.pc.Set(jumpLoc)
	}
}

func (program *Program) EndIf(jumpLoc int) {
	if program.state.Value() != 0 {
		program.pc.Set(jumpLoc)
	}
}

func (program *Program) ResetAndStep() {
	program.Reset()
	program.IncPosWith(1)
}

func (program *Program) Reset() {
	program.state.data[program.state.pos] = 0
}
