package slice

import (
	"github.com/cirius-go/generic/common"
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
func Unshift[T any](value T, items ...T) []T {
	return append([]T{value}, items...)
}

// Reverse reverses an array
func Reverse[T any](items ...T) []T {
	result := make([]T, 0)

	for i := len(items) - 1; i >= 0; i-- {
		result = append(result, items[i])
	}

	return result
}

// FindIndex returns the index of the first element in the arr.
func FindIndex[T any](predicate func(T) bool, items ...T) (index int, found bool) {
	if predicate == nil {
		return -1, false
	}

	for i, v := range items {
		if predicate(v) {
			return i, true
		}
	}

	return -1, false
}

// Find returns the first element in the array that satisfies the provided
func Find[T any](predicate func(T) bool, items ...T) (value T, found bool) {
	index, ok := FindIndex(predicate, items...)
	if !ok || index == -1 {
		return value, false
	}

	return items[index], true
}

// Filter elements of an array
func Filter[T any](predicate func(T) bool, items ...T) []T {
	var result = make([]T, 0)

	if predicate == nil {
		return items
	}

	for _, v := range items {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

// Every returns true if every element in the array satisfies the provided
func Every[T any](predicate func(T) bool, items ...T) bool {
	if predicate == nil {
		return true
	}

	for _, v := range items {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Some returns true if any element in the array satisfies the provided
func Some[T any](predicate func(T) bool, items ...T) bool {
	if predicate == nil {
		return false
	}
	for _, v := range items {
		if predicate(v) {
			return true
		}
	}

	return false
}

// Map returns a new array with the results of calling the provided function
func Map[T, R any](callback func(T) R, items ...T) []R {
	var result []R

	for _, v := range items {
		result = append(result, callback(v))
	}

	return result
}

// At returns the value at the given index
func At[T any](index int, items ...T) (T, bool) {
	if index < 0 || index >= len(items) {
		return common.Zero[T](), false
	}

	return items[index], true
}

// Includes returns true if the array includes the value
func Includes[T comparable](value T, items ...T) bool {
	for _, v := range items {
		if v == value {
			return true
		}
	}

	return false
}

// Pop removes the last element of an array and returns it.
func Pop[T any](items ...T) []T {
	if len(items) == 0 {
		return make([]T, 0)
	}

	return items[:len(items)-1]
}

// Reduce reduces an array to a single value
func Reduce[T, R any](initialValue R, callback func(R, T) R, items ...T) R {
	for _, v := range items {
		initialValue = callback(initialValue, v)
	}

	return initialValue
}

// ReduceWithError reduces an array to a single value with an error
func ReduceWithError[T, R any](initialValue R, callback func(R, T) (R, error), items ...T) (R, error) {
	for _, v := range items {
		r, err := callback(initialValue, v)
		if err != nil {
			return r, err
		}

		initialValue = r
	}

	return initialValue, nil
}

// Shift removes the first element of an array and returns it.
func Shift[T any](items ...T) []T {
	if len(items) == 0 {
		return make([]T, 0)
	}

	return items[1:]
}

// Clone clones an array
func Clone[T any](arr []T) []T {
	var result []T

	return append(result, arr...)
}

// NonZero returns the non-zero elements of an array
func NonZero[T comparable](items ...T) []T {
	res := make([]T, 0)
	if len(items) == 0 {
		return res
	}

	for _, v := range items {
		if common.IsZero(v) {
			continue
		}

		res = append(res, v)
	}

	return res
}

// FisrtNonZero returns the first non-zero element of an array.
func FisrtNonZero[T comparable](items ...T) (T, bool) {
	if len(items) == 0 {
		return common.Zero[T](), false
	}

	for _, v := range items {
		if common.IsZero(v) {
			continue
		}

		return v, true
	}

	return common.Zero[T](), false
}

// FirstOrDefault returns the first non-zero element of an array.
func FirstOrDefault[T comparable](def T, items ...T) T {
	v, found := FisrtNonZero(items...)
	if found {
		return v
	}

	return def
}

// FirstOrDefaultArr returns the first non-zero element of an array.
func FirstOrDefaultArr[T any](def []T, items ...[]T) []T {
	for _, v := range items {
		if len(v) > 0 {
			return v
		}
	}

	return def
}

// MapTilError applies a function to each element of an array.
func MapTilError[T, R any](callback func(T) (R, error), items ...T) ([]R, error) {
	result := make([]R, 0)

	for _, v := range items {
		r, err := callback(v)
		if err != nil {
			return result, err
		}

		result = append(result, r)
	}

	return result, nil
}

func MapSkip[T, R any](callback func(T) (R, bool), items ...T) []R {
	result := make([]R, 0)

	for _, v := range items {
		r, skip := callback(v)
		if skip {
			continue
		}

		result = append(result, r)
	}

	return result
}

// ToAnys converts array of type T to array of type any.
func ToAnys[T any](items ...T) []any {
	return Map(func(item T) any {
		return item
	}, items...)
}