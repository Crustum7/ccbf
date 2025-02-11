package bytecode

import (
	"slices"
)

type Operation struct {
	pattern                string
	opCode                 byte
	repeated               bool
	numberOfParameterBytes int
}

var operations = []Operation{
	{pattern: ">", opCode: 1, repeated: false, numberOfParameterBytes: 0},
	{pattern: "<", opCode: 2, repeated: false, numberOfParameterBytes: 0},
	{pattern: "+", opCode: 3, repeated: false, numberOfParameterBytes: 0},
	{pattern: "-", opCode: 4, repeated: false, numberOfParameterBytes: 0},
	{pattern: ".", opCode: 5, repeated: false, numberOfParameterBytes: 0},
	{pattern: ",", opCode: 6, repeated: false, numberOfParameterBytes: 0},
	{pattern: "[", opCode: 7, repeated: false, numberOfParameterBytes: 4},
	{pattern: "]", opCode: 8, repeated: false, numberOfParameterBytes: 4},
	{pattern: "+", opCode: 9, repeated: true, numberOfParameterBytes: 1},
	{pattern: "-", opCode: 10, repeated: true, numberOfParameterBytes: 1},
	{pattern: ">", opCode: 11, repeated: true, numberOfParameterBytes: 1},
	{pattern: "<", opCode: 12, repeated: true, numberOfParameterBytes: 1},
	{pattern: "[-]>", opCode: 13, repeated: false, numberOfParameterBytes: 0},
	{pattern: "[-]", opCode: 14, repeated: false, numberOfParameterBytes: 0},
}

func (operation Operation) standardParameterBytes(repetitions int) []byte {
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
