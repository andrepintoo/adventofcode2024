package utils

import (
	"errors"
	"sync"
)

// Stack using slices (for linkedlist structure, use a *node instead of the slice)
// it is debatable but slices are usually more efficient due to better memory locality and lower memory overhead (no extra pointers)
// any is an alias for interface{} 

// TODO: testing
type Stack[T any] struct {
	values []T
	capacity int
	lock sync.Mutex
}

func NewStack[T any](cap int) *Stack[T]{
	return &Stack[T]{values: make([]T, 0, cap), capacity: cap}
}

func (s *Stack[T]) Push(val T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.values = append([]T{val}, s.values...)

	//if len(s.values) != 0 {
	//	s.values = append(s.values, val)
	//	copy(s.values[1:], s.values)
	//}
	//s.values[0] = val
}

func (s *Stack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.values) == 0 {
		var zeroVal T
		return zeroVal, errors.New("empty stack")
	}
	
	val := s.values[0]
	s.values = s.values[1:]
	return val, nil
}

// Return the top element without removing it
func (s *Stack[T]) Peek() (T, error) {
	if len(s.values) == 0 {
		var zeroVal T
		return zeroVal, errors.New("empty stack")
	}

	return s.values[0], nil
}

func (s *Stack[T]) Len() int{
	return len(s.values)
}

// Clears the stack
func (s *Stack[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.values = make([]T, 0, s.capacity)
}

func (s *Stack[T]) GetValues() []T {
	s.lock.Lock()
	defer s.lock.Unlock()

	return append([]T(nil), s.values...)
}
