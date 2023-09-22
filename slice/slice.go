package slice

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cirius-go/generic/common"
	"github.com/cirius-go/generic/types"
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

func ConcatUnique[T comparable](arrs ...[]T) []T {
	joined := Concat(arrs...)

	return RemoveDuplicates(joined...)
}

func RemoveDuplicates[T comparable](arr ...T) []T {
	// Create a map to store unique elements
	uniqueMap := make(map[T]bool)

	// Create a new slice to store unique elements
	uniqueArr := []T{}

	// Iterate through the array and add elements to the map
	// Only if they are not already present in the map
	for _, num := range arr {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueArr = append(uniqueArr, num)
		}
	}

	return uniqueArr
}

// Unshift adds a value to the beginning of an array
func Unshift[T any](value T, items ...T) []T {
	return append([]T{value}, items...)
}

func IUnshift[T any](items []T, value T) []T {
	return Unshift(value, items...)
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
func FindIndex[T any](predicate func(T) bool, items ...T) (index int) {
	if predicate == nil {
		return -1
	}

	for i := 0; i < len(items); i++ {
		if predicate(items[i]) {
			return i
		}
	}

	return -1
}

// FindIndex returns the index of the first element in the arr.
func IFindIndex[T any](items []T, predicate func(T) bool) (index int) {
	return FindIndex(predicate, items...)
}

// Find returns the first element in the array that satisfies the provided
func Find[T any](predicate func(T) bool, items ...T) (value T, found bool) {
	index := FindIndex(predicate, items...)
	if index == -1 {
		return value, false
	}

	return items[index], true
}

// FindOrDefault find the first element in the array that satisfies the provided.
// If not found, return the default value
func FindOrDefault[T any](predicate func(T) bool, items ...T) (value T) {
	index := FindIndex(predicate, items...)
	if index == -1 {
		return value
	}

	return items[index]
}

// IFind returns the first element in the array that satisfies the provided
func IFind[T any](items []T, predicate func(T) bool) (value T, found bool) {
	return Find(predicate, items...)
}

// IFindOrDefault find the first element in the array that satisfies the
// provided. If not found, return the default value.
func IFindOrDefault[T any](items []T, predicate func(T) bool) (value T) {
	return FindOrDefault(predicate, items...)
}

// Filter elements of an array
func Filter[T any](predicate func(T) bool, items ...T) []T {
	result := make([]T, 0)

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

// IFilter returns a new array with the results of calling the provided
func IFilter[T any](items []T, predicate func(T) bool) []T {
	return Filter(predicate, items...)
}

// FilterAndSeparate returns filtered items and which are not satisfied
// predicate
func FilterAndSeparate[T any](predicate func(T) bool, items ...T) ([]T, []T) {
	result := make([]T, 0)
	separated := make([]T, 0)

	for i := 0; i < len(items); i++ {
		if predicate(items[i]) {
			result = append(result, items[i])
		} else {
			separated = append(separated, items[i])
		}
	}

	return result, separated
}

// IFilterAndSeparate alias FilterAndSeparate.
func IFilterAndSeparate[T any](items []T, predicate func(T) bool) ([]T, []T) {
	return FilterAndSeparate(predicate, items...)
}

// Every returns true if every element in the array satisfies the provided
func Every[T any](predicate func(T) bool, items ...T) bool {
	if predicate == nil {
		return true
	}

	for i := 0; i < len(items); i++ {
		if !predicate(items[i]) {
			return false
		}
	}

	return true
}

// IEvery alias Every.
func IEvery[T any](items []T, predicate func(T) bool) bool {
	return Every(predicate, items...)
}

// Some returns true if any element in the array satisfies the provided.
func Some[T any](predicate func(T) bool, items ...T) bool {
	if predicate == nil {
		return false
	}
	for i := 0; i < len(items); i++ {
		if predicate(items[i]) {
			return true
		}
	}

	return false
}

// ISome alias Some.
func ISome[T any](items []T, predicate func(T) bool) bool {
	return Some(predicate, items...)
}

// Map returns a new array with the results of calling the provided function
func Map[T, R any](callback func(T) R, items ...T) []R {
	var result []R

	for i := 0; i < len(items); i++ {
		result = append(result, callback(items[i]))
	}

	return result
}

// IMap alias Map.
func IMap[T, R any](items []T, callback func(T) R) []R {
	return Map(callback, items...)
}

// At returns the value at the given index
func At[T any](index int, items ...T) (T, bool) {
	if index < 0 || index >= len(items) {
		return common.Zero[T](), false
	}

	return items[index], true
}

// IAt alias At.
func IAt[T any](items []T, index int) (T, bool) {
	return At(index, items...)
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

// IIncludes alias Includes.
func IIncludes[T comparable](items []T, value T) bool {
	return Includes(value, items...)
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

// IReduce alias Reduce.
func IReduce[T, R any](items []T, initialValue R, callback func(R, T) R) R {
	return Reduce(initialValue, callback, items...)
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

// IReduceWithError alias ReduceWithError.
func IReduceWithError[T, R any](items []T, initialValue R, callback func(R, T) (R, error)) (R, error) {
	return ReduceWithError(initialValue, callback, items...)
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

func IMapTilError[T, R any](items []T, callback func(T) (R, error)) ([]R, error) {
	return MapTilError(callback, items...)
}

// MapSkip applies a function to each element of an array
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

// IMapSkip alias MapSkip.
func IMapSkip[T, R any](items []T, callback func(T) (R, bool)) []R {
	return MapSkip(callback, items...)
}

// ToAnys converts array of type T to array of type any.
func ToAnys[T any](items ...T) []any {
	return Map(func(item T) any {
		return item
	}, items...)
}

func ExcludeIfIn[T comparable](sliceA []T, sliceB ...T) []T {
	return Filter(func(itemA T) bool {
		return !Includes(itemA, sliceB...)
	}, sliceA...)
}

func ExcludeIfNotIn[T comparable](sliceA []T, sliceB ...T) []T {
	return Filter(func(itemA T) bool {
		return Includes(itemA, sliceB...)
	}, sliceA...)
}

func Sort[T comparable](swapFn func(i, j int) bool, slice ...T) []T {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if swapFn(i, j) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

// Check all elements in b are in a
func ArrContains[T comparable](a []T, b []T) bool {
	for _, v := range b {
		if !Includes(v, a...) {
			return false
		}
	}

	return true
}

// Check all elements in b are in a
func ContainsAll[T comparable](sliceA []T, sliceB ...T) bool {
	return ArrContains(sliceA, sliceB)
}

// MergeFn merge many elements as one.
func MergeFn[T types.MergingHandler[T]](cur T, next T) T {
	return cur.Merge(next)
}

// ReduceMergeFn merge many item as one.
func ReduceMergeFn[T types.MergingHandler[T]](def T, slice ...T) T {
	return Reduce(def, MergeFn[T], slice...)
}

// GetRandomArray returns random array from origin.
func GetRandomArray[T any](items []T, size int) []T {
	if size >= len(items) {
		return items
	}

	// Shuffle the array using Fisher-Yates algorithm with crypto/rand
	shuffledItems := Shuffle(items)

	// Select the first 'size' elements as the random array
	return shuffledItems[:size]
}

// Shuffle items in array.
func Shuffle[T any](arr []T) []T {
	shuffled := make([]T, len(arr))
	copy(shuffled, arr)

	for i := len(shuffled) - 1; i > 0; i-- {
		j, err := CryptoRandInt(i + 1)
		if err != nil {
			panic(err)
		}
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}

func CryptoRandInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("max should be greater than 0")
	}

	// Generate a random number within the specified range
	randomValue, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return int(randomValue.Int64()), nil
}

// Loop applies a function to each element of an array
func Loop[T any](callback func(index int, item T), items ...T) {
	for i := 0; i < len(items); i++ {
		callback(i, items[i])
	}
}

// ILoop alias Loop.
func ILoop[T any](items []T, callback func(index int, item T)) {
	for i := 0; i < len(items); i++ {
		callback(i, items[i])
	}
}
