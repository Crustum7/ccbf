package operations

type Operation struct {
	pattern                string
	opCode                 byte
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
	{pattern: `>`, opCode: byte(OneRightStep), numberOfParameterBytes: 0},
	{pattern: `<`, opCode: byte(OneLeftStep), numberOfParameterBytes: 0},
	{pattern: `\+`, opCode: byte(IncrementOne), numberOfParameterBytes: 0},
	{pattern: `-`, opCode: byte(DecrementOne), numberOfParameterBytes: 0},
	{pattern: `\.`, opCode: byte(PrintChar), numberOfParameterBytes: 0},
	{pattern: `,`, opCode: byte(InputChar), numberOfParameterBytes: 0},
	{pattern: `\[`, opCode: byte(StartLoop), numberOfParameterBytes: 4},
	{pattern: `\]`, opCode: byte(EndLoop), numberOfParameterBytes: 4},
	{pattern: `(\++)`, opCode: byte(IncrementMultiple), numberOfParameterBytes: 1},
	{pattern: `(-+)`, opCode: byte(DecrementMultiple), numberOfParameterBytes: 1},
	{pattern: `(>+)`, opCode: byte(MultipleRightStep), numberOfParameterBytes: 1},
	{pattern: `(<+)`, opCode: byte(MultipleLeftStep), numberOfParameterBytes: 1},
	{pattern: `\[-\]>`, opCode: byte(ResetAndStep), numberOfParameterBytes: 0},
	{pattern: `\[-\]`, opCode: byte(Reset), numberOfParameterBytes: 0},
	{pattern: `\[-(>+)\+<+\]`, opCode: byte(MoveValueRight), numberOfParameterBytes: 1},
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
