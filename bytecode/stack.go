package bytecode

type Stack[T interface{}] struct {
	data []T
}

func (stack *Stack[T]) Size() int {
	return len(stack.data)
}

func (stack *Stack[T]) Push(val ...T) {
	stack.data = append(stack.data, val...)
}

func (stack *Stack[T]) Pop() T {
	val := stack.data[stack.Size()-1]
	stack.data = stack.data[:stack.Size()-1]
	return val
}

func InitStack[T interface{}]() Stack[T] {
	var stack Stack[T]
	stack.data = make([]T, 0)
	return stack
}
