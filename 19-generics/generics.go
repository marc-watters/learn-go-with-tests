package generics

type (
	StackOfInts    = Stack
	StackOfStrings = Stack
)

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value any) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (any, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
