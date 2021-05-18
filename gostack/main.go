package gostack

type Stack struct {
	ptr   int
	stack []int
}

func NewStack() *Stack {
	return &Stack{ptr: 0, stack: make([]int, 0)}
}

func (s *Stack) Push(i int) *Stack {
	s.stack = append(s.stack, i)
	s.ptr += 1
	return s
}

func (s *Stack) Pop() *Stack {
	if !s.isEmpty() {
		s.stack = s.stack[:s.ptr-1]
		s.ptr -= 1
	}

	return s
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) <= 0
}
