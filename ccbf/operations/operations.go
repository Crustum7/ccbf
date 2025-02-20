package operations

import (
	"slices"
)

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
)

var operations = []Operation{
	{pattern: ">", opCode: byte(OneRightStep), repeated: false, numberOfParameterBytes: 0},
	{pattern: "<", opCode: byte(OneLeftStep), repeated: false, numberOfParameterBytes: 0},
	{pattern: "+", opCode: byte(IncrementOne), repeated: false, numberOfParameterBytes: 0},
	{pattern: "-", opCode: byte(DecrementOne), repeated: false, numberOfParameterBytes: 0},
	{pattern: ".", opCode: byte(PrintChar), repeated: false, numberOfParameterBytes: 0},
	{pattern: ",", opCode: byte(InputChar), repeated: false, numberOfParameterBytes: 0},
	{pattern: "[", opCode: byte(StartLoop), repeated: false, numberOfParameterBytes: 4},
	{pattern: "]", opCode: byte(EndLoop), repeated: false, numberOfParameterBytes: 4},
	{pattern: "+", opCode: byte(IncrementMultiple), repeated: true, numberOfParameterBytes: 1},
	{pattern: "-", opCode: byte(DecrementMultiple), repeated: true, numberOfParameterBytes: 1},
	{pattern: ">", opCode: byte(MultipleRightStep), repeated: true, numberOfParameterBytes: 1},
	{pattern: "<", opCode: byte(MultipleLeftStep), repeated: true, numberOfParameterBytes: 1},
	{pattern: "[-]>", opCode: byte(ResetAndStep), repeated: false, numberOfParameterBytes: 0},
	{pattern: "[-]", opCode: byte(Reset), repeated: false, numberOfParameterBytes: 0},
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

func (operation Operation) StandardParameterBytes(repetitions int) []byte {
	if operation.repeated {
		return []byte{byte(repetitions)}
	}

	return []byte{}
}

func (operation Operation) ParsedSymbols(repetitions int) int {
	if operation.repeated {
		return len(operation.pattern) * repetitions
	}

	return len(operation.pattern)
}

func OperationPatterns() []string {
	unique := make(map[string]bool, 0)
	for _, operation := range operations {
		unique[operation.pattern] = true
	}
	patterns := make([]string, 0)
	for pattern := range unique {
		patterns = append(patterns, pattern)
	}
	return patterns
}

func OperationForPattern(pattern string, repeated bool) *Operation {
	i := slices.IndexFunc(operations, func(operation Operation) bool {
		return pattern == operation.pattern && repeated == operation.repeated
	})

	if i == -1 {
		i = slices.IndexFunc(operations, func(operation Operation) bool {
			return pattern == operation.pattern
		})
	}

	if i == -1 {
		return nil
	}

	return &operations[i]
}

func OperationForOpCode(opCode byte) *Operation {
	index := int(opCode) - 1
	if index < 0 || index >= len(operations) {
		return nil
	}

	return &operations[index]
}
