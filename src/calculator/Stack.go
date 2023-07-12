package calculator

type Stack struct {
	input []Number
	size  int
}

func (stack *Stack) Push(input Number) {
	stack.input = append(stack.input, input)
	stack.size++
}

func (stack *Stack) Pop() Number {
	item := stack.input[stack.size-1]
	stack.input = stack.input[:stack.size-1]
	stack.size--
	return item
}

func (stack *Stack) Peek() Number {
	return stack.input[stack.size-1]
}
