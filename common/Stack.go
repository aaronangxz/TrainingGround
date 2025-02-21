package common

type Stack[T any] struct {
	elements []T
}

// IsEmpty Helper function to check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Push the new element onto the stack
func (s *Stack[T]) Push(e T) {
	s.elements = append(s.elements, e)
}

// Pop Removes the element on the top of the stack
func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}
	index := len(s.elements) - 1
	s.elements = s.elements[:index]
}

// Top Returns the element on top of the stack
func (s *Stack[T]) Top() T {
	var zero T
	if s.IsEmpty() {
		return zero
	}
	index := len(s.elements) - 1
	return s.elements[index]
}

func (s *Stack[T]) Elements() []T {
	return append([]T(nil), s.elements...) // Safe copy
}
