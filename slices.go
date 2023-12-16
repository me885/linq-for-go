package linq

import "errors"

var (
	ErrorEmptySlice    = errors.New("slice must not be empty or nil")
	ErrorNotSliceOfOne = errors.New("slice must contain only a single element")
	ErrorEmptyMap      = errors.New("map must not be empty or nil")
	ErrorNotMapOfOne   = errors.New("map must contain only a single element")
)

// Returns a new slice with only the elements for which f() returns true
func SliceWhere[T any](s []T, f func(T) bool) []T {
	newSlice := make([]T, 0)

	for _, v := range s {
		if f(v) {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

// Returns a new slice with f() applied to every element in s
func SliceSelect[In any, Out any](s []In, f func(In) Out) []Out {
	newSlice := make([]Out, len(s))

	for i, v := range s {
		newSlice[i] = f(v)
	}

	return newSlice
}

// Returns the first element of s and an error if it is empty
func SliceFirst[T any](s []T) (T, error) {
	if len(s) == 0 {
		return *new(T), ErrorEmptySlice
	}

	return s[0], nil
}

// Returns the only element of s and an error if it does not have exactly one element
func SliceSingle[T any](s []T) (T, error) {
	if len(s) != 1 {
		return *new(T), ErrorNotSliceOfOne
	}

	return s[0], nil
}

// Returns true if f() return true for at least one element of s
func SliceAny[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}

	return false
}

// Returns true only if f() return true for all elements of s
func SliceAll[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

func SliceAggregate[T any](slice []T, seed T, f func(a T, b T) T) T {
	agg := seed

	for _, v := range slice {
		agg = f(agg, v)
	}

	return agg
}
