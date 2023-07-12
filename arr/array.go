package arr

import (
	"github.com/cirius-go/generic"
)

// Concat joins multiple slices of type T into a single slice.
func Concat[T any](arrs ...[]T) []T {
	result := make([]T, 0)

	for _, arr := range arrs {
		if arr == nil {
			continue
		}

		result = append(result, arr...)
	}

	return result
}

// Unshift adds a value to the beginning of an array
func Unshift[T any](value T, arr ...T) []T {
	return append([]T{value}, arr...)
}

// Reverse reverses an array
func Reverse[T any](arr ...T) []T {
	result := make([]T, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}

	return result
}

// FindIndex returns the index of the first element in the arr.
func FindIndex[T any](predicate func(T) bool, arr ...T) (index int, found bool) {
	if predicate == nil {
		return -1, false
	}

	for i, v := range arr {
		if predicate(v) {
			return i, true
		}
	}

	return -1, false
}

// Find returns the first element in the array that satisfies the provided
func Find[T any](predicate func(T) bool, arr ...T) (value T, found bool) {
	index, ok := FindIndex(predicate, arr...)
	if !ok || index == -1 {
		return value, false
	}

	return arr[index], true
}

// Filter elements of an array
func Filter[T any](predicate func(T) bool, arr ...T) []T {
	var result []T

	if predicate == nil {
		return arr
	}

	for _, v := range arr {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

// Every returns true if every element in the array satisfies the provided
func Every[T any](predicate func(T) bool, arr ...T) bool {
	if predicate == nil {
		return true
	}

	for _, v := range arr {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Some returns true if any element in the array satisfies the provided
func Some[T any](predicate func(T) bool, arr ...T) bool {
	if predicate == nil {
		return false
	}
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}

	return false
}

// Map returns a new array with the results of calling the provided function
func Map[T, R any](callback func(T) R, arr ...T) []R {
	var result []R

	for _, v := range arr {
		result = append(result, callback(v))
	}

	return result
}

// At returns the value at the given index
func At[T comparable](index int, arr ...T) (T, bool) {
	if index < 0 || index >= len(arr) {
		return generic.Zero[T](), false
	}

	return arr[index], true
}

// Includes returns true if the array includes the value
func Includes[T comparable](value T, arr ...T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

// Pop removes the last element of an array and returns it.
func Pop[T any](arr ...T) []T {
	if len(arr) == 0 {
		return arr
	}

	return arr[:len(arr)-1]
}

// Reduce reduces an array to a single value
func Reduce[T, R any](initialValue R, callback func(R, T) R, arr ...T) R {
	for _, v := range arr {
		initialValue = callback(initialValue, v)
	}

	return initialValue
}

// Shift removes the first element of an array and returns it.
func Shift[T any](arr ...T) []T {
	if len(arr) == 0 {
		return arr
	}

	return arr[1:]
}

// Clone clones an array
func Clone[T any](arr []T) []T {
	var result []T

	for _, v := range arr {
		result = append(result, v)
	}

	return result
}

// NonZero returns the non-zero elements of an array
func NonZero[T comparable](arr ...T) []T {
	res := make([]T, 0)
	if len(arr) == 0 {
		return res
	}

	for _, v := range arr {
		if generic.IsZero(v) {
			continue
		}

		res = append(res, v)
	}

	return res
}

// FisrtNonZero returns the first non-zero element of an array.
func FisrtNonZero[T comparable](arr ...T) (T, bool) {
	if len(arr) == 0 {
		return generic.Zero[T](), false
	}

	for _, v := range arr {
		if generic.IsZero(v) {
			continue
		}

		return v, true
	}

	return generic.Zero[T](), false
}

// FirstOrDefault returns the first non-zero element of an array.
func FirstOrDefault[T comparable](def T, arr ...T) T {
	v, found := FisrtNonZero(arr...)
	if found {
		return v
	}

	return def
}

// FirstOrDefaultArr returns the first non-zero element of an array.
func FirstOrDefaultArr[T any](def []T, arr ...[]T) []T {
	for _, v := range arr {
		if len(v) > 0 {
			return v
		}
	}

	return def
}
