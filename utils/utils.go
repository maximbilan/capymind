package utils

import "sync"

// ThreadSafeArray is a generic struct that contains a slice and a mutex
type ThreadSafeArray[T comparable] struct {
	mu    sync.Mutex
	array []T
}

// NewThreadSafeArray initializes a new ThreadSafeArray
func NewThreadSafeArray[T comparable]() *ThreadSafeArray[T] {
	return &ThreadSafeArray[T]{
		array: make([]T, 0),
	}
}

// Append safely appends an element to the array
func (tsa *ThreadSafeArray[T]) Append(value T) {
	tsa.mu.Lock()
	defer tsa.mu.Unlock()
	tsa.array = append(tsa.array, value)
}

// Get safely retrieves an element from the array by index
func (tsa *ThreadSafeArray[T]) Get(index int) (T, bool) {
	tsa.mu.Lock()
	defer tsa.mu.Unlock()
	if index < 0 || index >= len(tsa.array) {
		var zero T
		return zero, false
	}
	return tsa.array[index], true
}

// Length safely returns the length of the array
func (tsa *ThreadSafeArray[T]) Length() int {
	tsa.mu.Lock()
	defer tsa.mu.Unlock()
	return len(tsa.array)
}

// Contains safely checks if the array contains the specified element
func (tsa *ThreadSafeArray[T]) Contains(value T) bool {
	tsa.mu.Lock()
	defer tsa.mu.Unlock()
	for _, v := range tsa.array {
		if v == value {
			return true
		}
	}
	return false
}

// Remove safely removes the first occurrence of the specified element from the array
func (tsa *ThreadSafeArray[T]) Remove(value T) bool {
	tsa.mu.Lock()
	defer tsa.mu.Unlock()
	for i, v := range tsa.array {
		if v == value {
			tsa.array = append(tsa.array[:i], tsa.array[i+1:]...)
			return true
		}
	}
	return false
}
