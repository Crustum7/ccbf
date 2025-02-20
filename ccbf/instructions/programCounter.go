package instructions

type ProgramCounter struct {
	counter int
}

func InitProgramCounter() ProgramCounter {
	return ProgramCounter{counter: 0}
}

func (pc ProgramCounter) Get() int {
	return pc.counter
}

func (pc *ProgramCounter) Set(val int) {
	pc.counter = val
}

func (pc *ProgramCounter) Increment() {
	pc.IncrementWith(1)
}

func (pc *ProgramCounter) IncrementWith(diff int) {
	pc.counter += diff
}
