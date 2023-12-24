package main

import "errors"

func remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func contains[T comparable](slice []T, elem T) bool {
	for _, s := range slice {
		if elem == s {
			return true
		}
	}
	return false
}

type Stack[T any] []T

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value to stack
func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

// Pop removes and returns the top element of the stack. Returns false if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item, nil
}
