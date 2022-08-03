package util

type Stack[T int | string | float64 | byte] struct {
	element []T
	len     int
}

func (stack *Stack[T]) Push(val T) {
	stack.element = append(stack.element, val)
	stack.len++
}

func (stack *Stack[T]) Pop() T {
	res := stack.element[stack.len-1]
	stack.len--
	stack.element = stack.element[0:stack.len]
	return res
}

func (stack *Stack[T]) Top() T {
	return stack.element[stack.len-1]
}

func (stack *Stack[T]) Size() int {
	return stack.len
}

func (stack *Stack[T]) isEmpty() bool {
	return stack.len == 0
}
