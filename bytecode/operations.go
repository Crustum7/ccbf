package bytecode

import (
	"slices"
)

type Operation struct {
	pattern                string
	opCode                 byte
	numberOfParameterBytes int
}

var operations = []Operation{
	{pattern: ">", opCode: 1, numberOfParameterBytes: 0},
	{pattern: "<", opCode: 2, numberOfParameterBytes: 0},
	{pattern: "+", opCode: 3, numberOfParameterBytes: 0},
	{pattern: "-", opCode: 4, numberOfParameterBytes: 0},
	{pattern: ".", opCode: 5, numberOfParameterBytes: 0},
	{pattern: ",", opCode: 6, numberOfParameterBytes: 0},
	{pattern: "[", opCode: 7, numberOfParameterBytes: 4},
	{pattern: "]", opCode: 8, numberOfParameterBytes: 4},
}

func OperationForPattern(pattern string) *Operation {
	i := slices.IndexFunc(operations, func(operation Operation) bool {
		return pattern == operation.pattern
	})
	if i == -1 {
		return nil
	}
	return &operations[i]
}
