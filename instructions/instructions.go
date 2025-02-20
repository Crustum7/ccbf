package instructions

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
	char := byte(program.state.Value())
	program.write(char)
}

func (program *Program) CharIn() {
	val := program.read()
	program.state.data[program.state.pos] = val
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
