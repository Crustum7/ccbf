package operations

type Operation struct {
	pattern                string
	opCode                 byte
	repeated               bool
	numberOfParameterBytes int
}

type OpCode int

const (
	Undefined OpCode = iota
	OneRightStep
	OneLeftStep
	IncrementOne
	DecrementOne
	PrintChar
	InputChar
	StartLoop
	EndLoop
	IncrementMultiple
	DecrementMultiple
	MultipleRightStep
	MultipleLeftStep
	ResetAndStep
	Reset
	MoveValueRight
)

var operations = []Operation{
	{pattern: `>`, opCode: byte(OneRightStep), repeated: false, numberOfParameterBytes: 0},
	{pattern: `<`, opCode: byte(OneLeftStep), repeated: false, numberOfParameterBytes: 0},
	{pattern: `\+`, opCode: byte(IncrementOne), repeated: false, numberOfParameterBytes: 0},
	{pattern: `-`, opCode: byte(DecrementOne), repeated: false, numberOfParameterBytes: 0},
	{pattern: `\.`, opCode: byte(PrintChar), repeated: false, numberOfParameterBytes: 0},
	{pattern: `,`, opCode: byte(InputChar), repeated: false, numberOfParameterBytes: 0},
	{pattern: `\[`, opCode: byte(StartLoop), repeated: false, numberOfParameterBytes: 4},
	{pattern: `\]`, opCode: byte(EndLoop), repeated: false, numberOfParameterBytes: 4},
	{pattern: `(\++)`, opCode: byte(IncrementMultiple), repeated: true, numberOfParameterBytes: 1},
	{pattern: `(-+)`, opCode: byte(DecrementMultiple), repeated: true, numberOfParameterBytes: 1},
	{pattern: `(>+)`, opCode: byte(MultipleRightStep), repeated: true, numberOfParameterBytes: 1},
	{pattern: `(<+)`, opCode: byte(MultipleLeftStep), repeated: true, numberOfParameterBytes: 1},
	{pattern: `\[-\]>`, opCode: byte(ResetAndStep), repeated: false, numberOfParameterBytes: 0},
	{pattern: `\[-\]`, opCode: byte(Reset), repeated: false, numberOfParameterBytes: 0},
	// {pattern: `\[->\+<\]`, opCode: byte(MoveValueRight), repeated: false, numberOfParameterBytes: 0},
}

func (operation Operation) GetPattern() string {
	return operation.pattern
}

func (operation Operation) GetOpCode() byte {
	return operation.opCode
}

func (operation Operation) GetParameterByteCount() int {
	return operation.numberOfParameterBytes
}

func GetOperations() []Operation {
	return operations
}

func OperationPatterns() []string {
	patterns := make([]string, 0)
	for _, operation := range operations {
		patterns = append(patterns, operation.pattern)
	}
	return patterns
}

func OperationForOpCode(opCode byte) *Operation {
	index := int(opCode) - 1
	if index < 0 || index >= len(operations) {
		return nil
	}

	return &operations[index]
}
