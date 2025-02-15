package main

import (
	"errors"
	"fmt"
)

// A stack is a data structure that allows data to be inserted (pushed) and removed (popped)
// in a last-in-first-out (LIFO) order.
type Stack struct {
	data []int
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// The last value pushed to the stack is the first value popped.
	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println(value)
	}

	// Pop from an empty stack.
	value, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
