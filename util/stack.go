package util

type Stack[T any] struct {
	contents []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{contents: make([]T, 0)}
}

func (stack *Stack[T]) Push(value T) {
	stack.contents = append(stack.contents, value)
}

func (stack *Stack[T]) Pop() T {
	if len(stack.contents) == 0 {
		panic("pop on empty stack")
	}
	topIndex := len(stack.contents) - 1
	value := stack.contents[topIndex]
	stack.contents = stack.contents[:topIndex]
	return value
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.contents) == 0
}
