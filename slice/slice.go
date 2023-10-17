package slice

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cirius-go/generic/common"
	"github.com/cirius-go/generic/types"
)

// Concat concatenates multiple arrays of any type into a single array.
//
// arrs: variadic parameter of arrays to be concatenated.
// Returns: a new array containing all the elements from the input arrays.
func Concat[T any](arrs ...[]T) []T {
	result := make([]T, 0)

	for i := range arrs {
		arr := arrs[i]
		if arr == nil {
			continue
		}

		result = append(result, arr...)
	}

	return result
}

// ConcatUnique concatenates multiple slices and removes duplicates.
//
// The function takes multiple slices of the same type as input and
// concatenates them using the Concat function. It then removes
// duplicate elements from the resulting slice using the
// RemoveDuplicates function.
//
// Parameters:
// - arrs: Variadic parameter that accepts multiple slices of the same type.
//
// Returns:
// - []T: The resulting slice after concatenation and removal of duplicates.
func ConcatUnique[T comparable](arrs ...[]T) []T {
	joined := Concat(arrs...)

	return RemoveDuplicates(joined...)
}

// RemoveDuplicates removes duplicates from the given array.
//
// The function takes a variadic parameter `arr` of type `T`, which is a slice of elements of any type that is comparable. It iterates through the array and adds elements to a map only if they are not already present in the map. It then returns a new slice with the unique elements.
//
// The return type is `[]T`, which is a slice of elements of type `T`.
func RemoveDuplicates[T comparable](arr ...T) []T {
	// Create a map to store unique elements
	uniqueMap := make(map[T]bool)

	// Create a new slice to store unique elements
	uniqueArr := []T{}

	// Iterate through the array and add elements to the map
	// Only if they are not already present in the map
	for i := 0; i < len(arr); i++ {
		item := arr[i]

		if !uniqueMap[item] {
			uniqueMap[item] = true
			uniqueArr = append(uniqueArr, item)
		}
	}

	return uniqueArr
}

// Unshift adds an element to the beginning of a slice.
//
// The first parameter, value, is the element to be added to the slice.
// The second parameter, items, is a variadic parameter representing the
// existing elements of the slice.
// The return type is a slice of type T, which is the same as the type
// of the elements in the slice.
func Unshift[T any](value T, items ...T) []T {
	return append([]T{value}, items...)
}

// IUnshift inserts a value at the beginning of a slice and returns the modified slice.
//
// The items parameter is the slice to which the value will be inserted.
// The value parameter is the value to be inserted at the beginning of the slice.
// The function returns the modified slice.
func IUnshift[T any](items []T, value T) []T {
	return Unshift(value, items...)
}

// Reverse reverses the order of the given items.
//
// The function takes a variadic parameter `items` of type `T` and returns a slice of type `[]T`.
func Reverse[T any](items ...T) []T {
	result := make([]T, 0)

	for i := len(items) - 1; i >= 0; i-- {
		result = append(result, items[i])
	}

	return result
}

// FindIndex finds the index of the first element in the given slice that satisfies the provided predicate function.
//
// The predicate function takes an element of type T and returns a boolean value indicating whether the element satisfies the condition.
// The items parameter is a variadic parameter of type T that represents the slice of elements to search in.
// The function returns the index of the first element that satisfies the predicate, or -1 if no element satisfies the predicate.
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

// IFindIndex finds the index of the first element in the given slice that satisfies the provided predicate function.
//
// It takes a slice of any type 'T' and a predicate function that accepts an element of type 'T' and returns a boolean value.
//
// The function returns the index of the first element that matches the predicate function, or -1 if no element is found.
func IFindIndex[T any](items []T, predicate func(T) bool) (index int) {
	return FindIndex(predicate, items...)
}

// Find is a generic function that takes a predicate function and a variadic number of items of any type.
// It returns the first item in the list that satisfies the predicate function and a boolean value indicating
// whether the item was found or not.
//
// The predicate function takes an item of type T and returns a boolean value indicating whether the item
// satisfies the condition.
//
// The function returns the found item of type T and a boolean value indicating whether the item was found or not.
func Find[T any](predicate func(T) bool, items ...T) (value T, found bool) {
	index := FindIndex(predicate, items...)
	if index == -1 {
		return value, false
	}

	return items[index], true
}

// FindOrDefault returns the first item in the given list that satisfies the given predicate.
//
// The predicate function takes an item of type T and returns a boolean indicating whether the item satisfies the condition.
// The items parameter is a variadic argument of type T, representing the list of items to search in.
// The function returns the first item in the list that satisfies the predicate function, or the zero value of type T if no item satisfies the condition.
func FindOrDefault[T any](predicate func(T) bool, items ...T) (value T) {
	index := FindIndex(predicate, items...)
	if index == -1 {
		return value
	}

	return items[index]
}

// IFind finds the first element in the given slice that satisfies the provided predicate function.
//
// It takes a slice of any type and a predicate function that takes an element of that type and returns a boolean value indicating whether the element satisfies the condition.
// It returns the first element that satisfies the condition and a boolean value indicating whether such an element was found.
func IFind[T any](items []T, predicate func(T) bool) (value T, found bool) {
	return Find(predicate, items...)
}

// IFindOrDefault returns the first element in the given slice that satisfies the given predicate function.
//
// The function takes two parameters:
// - items: a slice of type T, representing the collection of elements to search through.
// - predicate: a function that takes a parameter of type T and returns a boolean value indicating whether the element satisfies the condition.
//
// The function returns a value of type T, which is the first element in the slice that satisfies the predicate function.
func IFindOrDefault[T any](items []T, predicate func(T) bool) (value T) {
	return FindOrDefault(predicate, items...)
}

// Filter applies a given predicate function to each element in a slice and returns a new slice
// containing only the elements that satisfy the predicate.
//
// The predicate function takes an element of type T as input and returns a boolean value indicating
// whether the element satisfies the predicate or not.
//
// The parameter 'predicate' is the function to apply to each element in the slice.
// The parameter 'items' is the slice of elements to filter.
//
// The return type is a new slice of elements of type T that satisfy the predicate.
func Filter[T any](predicate func(T) bool, items ...T) []T {
	result := make([]T, 0)

	if predicate == nil {
		return items
	}

	for i := range items {
		v := items[i]

		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

// IFilter applies a predicate function to each item in a slice and returns a new slice containing
// only the items for which the predicate function returns true.
//
// items: The slice of items to filter.
// predicate: The function that determines whether an item should be included in the filtered slice.
//
// []T: The new slice that contains the filtered items.
func IFilter[T any](items []T, predicate func(T) bool) []T {
	return Filter(predicate, items...)
}

// FilterAndSeparate filters and separates items based on the provided predicate function.
//
// The function takes a predicate function as its first parameter, which is used to determine whether an item should be included in the result or not.
// The remaining parameters are the items to be filtered and separated.
//
// The function returns two slices: the first slice contains the items that satisfy the predicate, while the second slice contains the items that do not satisfy the predicate.
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

// IFilterAndSeparate is a Go function that filters and separates items based on a given predicate.
//
// It takes in a slice of items of any type and a predicate function as parameters.
// The predicate function determines whether an item should be included in the filtered slice.
//
// The function returns two slices of the same type: the filtered items and the separated items.
func IFilterAndSeparate[T any](items []T, predicate func(T) bool) ([]T, []T) {
	return FilterAndSeparate(predicate, items...)
}

// Every checks if all elements in items satisfy the given predicate function.
//
// The predicate function takes an element of type T as a parameter and returns a boolean value.
// The items parameter is a variadic parameter of type T, representing the elements to be checked.
// The function returns a boolean value indicating whether all elements satisfy the predicate function.
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

// IEvery applies the predicate function to every element in the items slice.
//
// It returns true if the predicate function returns true for all elements,
// and false otherwise.
//
// Parameters:
// - items: A slice of any type that contains the elements to be tested.
// - predicate: A function that takes an element of type T and returns a boolean value.
//
// Return type: bool
func IEvery[T any](items []T, predicate func(T) bool) bool {
	return Every(predicate, items...)
}

// Some is a function that takes a predicate function and a variadic list of items of type T, and returns a boolean value.
//
// The predicate function is used to determine whether an item satisfies a certain condition. It should take an item of type T as its argument and return a boolean value.
//
// The function returns true if at least one item in the list satisfies the predicate condition, and false otherwise.
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

// ISome checks if any element in the given slice satisfies the provided predicate function.
//
// items: The slice of elements to be checked.
// predicate: The function that determines if an element satisfies a condition.
// Returns: True if any element satisfies the condition, false otherwise.
func ISome[T any](items []T, predicate func(T) bool) bool {
	return Some(predicate, items...)
}

// Map applies a callback function to each element in the items slice and returns a new slice containing the results.
//
// The callback function takes an element of type T as input and returns an element of type R.
// The items slice is a variadic parameter that can accept zero or more elements of type T.
// The return type is a slice of elements of type R.
func Map[T, R any](callback func(T) R, items ...T) []R {
	var result []R

	for i := 0; i < len(items); i++ {
		result = append(result, callback(items[i]))
	}

	return result
}

// IMap applies a function to each item in a slice and returns a new slice
// containing the results.
//
// Parameters:
//   - items: A slice of type T, representing the items to be processed.
//   - callback: A function that takes an item of type T as input and returns a
//     result of type R.
//
// Return:
//   - A new slice of type R, containing the results of applying the callback
//     function to each item in the input slice.
func IMap[T, R any](items []T, callback func(T) R) []R {
	return Map(callback, items...)
}

// At returns the element at the specified index in the given slice.
//
// Parameters:
// - index: the index of the element to retrieve.
// - items: the slice of elements.
//
// Returns:
// - T: the element at the specified index.
// - bool: true if the element was found, false otherwise.
func At[T any](index int, items ...T) (T, bool) {
	if index < 0 || index >= len(items) {
		return common.Zero[T](), false
	}

	return items[index], true
}

// IAt returns the element at the specified index in the given slice.
//
// Parameters:
// - items: A slice of any type.
// - index: The index of the element to retrieve.
//
// Returns:
// - T: The element at the specified index.
// - bool: A boolean value indicating whether the element was found or not.
func IAt[T any](items []T, index int) (T, bool) {
	return At(index, items...)
}

// Includes checks if a value is present in a variadic list of items.
//
// The function takes a value of type T and a variadic list of items of type T. It
// iterates through the list and checks if any item in the list is equal to the
// given value. If a match is found, it returns true. Otherwise, it returns false.
//
// Parameters:
// - value: The value to search for in the list of items.
// - items: The variadic list of items to search through.
//
// Returns:
// - bool: True if the value is found in the list, false otherwise.
func Includes[T comparable](value T, items ...T) bool {
	for i := range items {
		v := items[i]

		if v == value {
			return true
		}
	}

	return false
}

// IIncludes checks if a value is included in a slice of comparable items.
//
// Parameters:
// - items: a slice of comparable items.
// - value: the value to check for inclusion.
//
// Return type:
// - bool
func IIncludes[T comparable](items []T, value T) bool {
	return Includes(value, items...)
}

// Pop removes and returns the last element of a slice.
//
// items: The slice to pop from.
// Returns: A new slice with the last element removed.
func Pop[T any](items ...T) []T {
	if len(items) == 0 {
		return make([]T, 0)
	}

	return items[:len(items)-1]
}

// Reduce applies a function against an accumulator and each element in the items slice,
// from left to right, to reduce it to a single value.
//
//   - initialValue: The initial value of the accumulator.
//   - callback: A function that takes the accumulator and the current element as arguments
//     and returns a new accumulator value.
//   - items: The slice of elements to iterate over.
//
// Returns the final value of the accumulator.
func Reduce[T, R any](initialValue R, callback func(R, T) R, items ...T) R {
	for i := range items {
		v := items[i]

		initialValue = callback(initialValue, v)
	}

	return initialValue
}

// IReduce reduces an array of items to a single value using a callback function.
//
// The items parameter is an array of any type.
// The initialValue parameter is the initial value for reduction.
// The callback parameter is a function that takes the accumulated value and the current item as arguments and returns the updated accumulated value.
// The return type is the same as the initialValue parameter.
func IReduce[T, R any](items []T, initialValue R, callback func(R, T) R) R {
	return Reduce(initialValue, callback, items...)
}

// ReduceWithError applies a callback function to each item in a collection and reduces them to a single value, while handling any errors that occur.
//
// Parameters:
// - initialValue: The initial value for the reduction.
// - callback: The callback function that is applied to each item in the collection.
// - items: The items to be reduced.
//
// Returns:
// - R: The reduced value.
// - error: An error that occurred during the reduction, if any.
func ReduceWithError[T, R any](initialValue R, callback func(R, T) (R, error), items ...T) (R, error) {
	for i := range items {
		v := items[i]

		r, err := callback(initialValue, v)
		if err != nil {
			return r, err
		}

		initialValue = r
	}

	return initialValue, nil
}

// IReduceWithError applies a callback function to each element of a slice, accumulating the results into a single value of type R.
//
// The function takes in the following parameters:
// - items: a slice of type T, representing the elements to be iterated over.
// - initialValue: a value of type R, representing the initial value for the accumulation.
// - callback: a function that takes in two parameters - a value of type R and a value of type T, and returns a value of type R and an error.
//
// The function returns two values:
// - The accumulated value of type R.
// - An error, if any occurred during the iteration process.
func IReduceWithError[T, R any](items []T, initialValue R, callback func(R, T) (R, error)) (R, error) {
	return ReduceWithError(initialValue, callback, items...)
}

// Shift returns a new slice containing all elements of the input slice except for the first element.
//
// items: The input slice of any type.
// []T: The resulting slice with the same type as the input slice.
func Shift[T any](items ...T) []T {
	if len(items) == 0 {
		return make([]T, 0)
	}

	return items[1:]
}

// Clone creates a shallow copy of the provided slice.
//
// The `arr` parameter is the slice to be cloned.
// The function returns a new slice that is a shallow copy of `arr`.
func Clone[T any](arr []T) []T {
	var result []T

	return append(result, arr...)
}

// NonZero filters out zero values from a list of items.
//
// The function takes a variadic parameter `items` of type `T`, which is a comparable type.
// It returns a slice of type `[]T` that contains only the non-zero values from the input slice.
func NonZero[T comparable](items ...T) []T {
	res := make([]T, 0)
	if len(items) == 0 {
		return res
	}

	for i := range items {
		v := items[i]
		if common.IsZero(v) {
			continue
		}

		res = append(res, v)
	}

	return res
}

// FisrtNonZero returns the first non-zero value from the given items, along with a boolean indicating if a non-zero value was found.
//
// It takes a variadic parameter 'items' of type 'T' which should be a comparable type.
// The return type is a tuple of 'T' which is the first non-zero value found, and a boolean indicating if a non-zero value was found.
func FisrtNonZero[T comparable](items ...T) (T, bool) {
	if len(items) == 0 {
		return common.Zero[T](), false
	}

	for i := range items {
		v := items[i]

		if common.IsZero(v) {
			continue
		}

		return v, true
	}

	return common.Zero[T](), false
}

// FirstOrDefault finds and returns the first non-zero value from the given items. If no non-zero value is found, it returns the default value.
//
// def: the default value to be returned if no non-zero value is found.
// items: variadic parameter representing the list of items to search for a non-zero value.
// T: the type of the items and the default value.
// Return: the first non-zero value found or the default value if no non-zero value is found.
func FirstOrDefault[T comparable](def T, items ...T) T {
	v, found := FisrtNonZero(items...)
	if found {
		return v
	}

	return def
}

// FirstOrDefaultArr returns the first non-empty array from a list of arrays.
//
// def: the default array to return if all other arrays are empty.
// items: variadic parameter representing the list of arrays.
// []T: the type of the arrays.
// []T: the type of the returned array.
func FirstOrDefaultArr[T any](def []T, items ...[]T) []T {
	for i := range items {
		v := items[i]
		if len(v) > 0 {
			return v
		}
	}

	return def
}

// MapTilError applies the callback function to each item in the items slice and returns a new slice with the results. If the callback function returns an error for any item, the mapping process is stopped and the error is returned.
//
// The callback function takes an item of type T and returns a result of type R and an error. The items parameter is a variadic parameter that accepts multiple items of type T. The function returns a new slice with the results of applying the callback function to each item, and an error if any occurred during the mapping process.
func MapTilError[T, R any](callback func(T) (R, error), items ...T) ([]R, error) {
	result := make([]R, 0)

	for i := range items {
		v := items[i]

		r, err := callback(v)
		if err != nil {
			return result, err
		}

		result = append(result, r)
	}

	return result, nil
}

// IMapTilError applies a callback function to each item in the given slice until an error occurs,
// and returns a new slice with the results of the callback function.
//
// The items parameter is a slice of type T, representing the input items to be processed.
// The callback parameter is a function that takes a single item of type T as input and returns
// a value of type R and an error. It is applied to each item in the slice until an error occurs.
// The function returns a new slice of type R, containing the results of the callback function,
// and an error if one occurred during the process.
func IMapTilError[T, R any](items []T, callback func(T) (R, error)) ([]R, error) {
	return MapTilError(callback, items...)
}

// MapSkip applies the callback function to each item in the input slice and returns a new slice containing the mapped results. It also skips any items for which the callback function returns true as the second value.
//
// The callback function takes an item of type T as its parameter and returns two values: the mapped result of type R and a boolean value indicating whether to skip the item.
// The input slice, items, is of type []T.
// The return type is []R.
func MapSkip[T, R any](callback func(T) (R, bool), items ...T) []R {
	result := make([]R, 0)

	for i := range items {
		v := items[i]

		r, skip := callback(v)
		if skip {
			continue
		}

		result = append(result, r)
	}

	return result
}

// IMapSkip applies the provided callback function to each element in the given slice `items` and returns a new slice containing the results. The callback function takes an element of type `T` as input and returns a result of type `R` and a boolean value indicating whether the element should be included in the output slice or skipped.
//
// Parameters:
// - items: The slice of elements of type `T` to be processed.
// - callback: The callback function that will be applied to each element in the input slice.
//
// Returns:
// - A new slice containing the results of applying the callback function to each element in the input slice.
func IMapSkip[T, R any](items []T, callback func(T) (R, bool)) []R {
	return MapSkip(callback, items...)
}

// ToAnys converts the given items to a slice of type []any.
//
// The function takes a variadic parameter of type `T`, which can be any type. It
// uses the `Map` function to iterate over each item in the `items` slice and
// returns a new slice of type `[]any` containing the same items.
//
// The return type of the function is `[]any`, which is the converted slice.
func ToAnys[T any](items ...T) []any {
	return Map(func(item T) any {
		return item
	}, items...)
}

// ExcludeIfIn filters out elements from sliceA that are present in sliceB.
//
// sliceA - the slice from which elements will be excluded.
// sliceB - the elements to be excluded from sliceA.
// Returns a new slice with the excluded elements.
func ExcludeIfIn[T comparable](sliceA []T, sliceB ...T) []T {
	return Filter(func(itemA T) bool {
		return !Includes(itemA, sliceB...)
	}, sliceA...)
}

// ExcludeIfNotIn filters out elements from sliceA that are not present in sliceB.
//
// Parameters:
// - sliceA: The input slice to be filtered.
// - sliceB: The values to be checked for inclusion in sliceA.
//
// Returns:
// - []T: The filtered slice with elements that are present in sliceB.
func ExcludeIfNotIn[T comparable](sliceA []T, sliceB ...T) []T {
	return Filter(func(itemA T) bool {
		return Includes(itemA, sliceB...)
	}, sliceA...)
}

// Sort sorts the given slice using the provided swapFn.
//
// The swapFn is a function that takes two indices i and j, and returns true if the elements at index i and j should be swapped, false otherwise.
// The slice is the input slice that needs to be sorted.
// The return value is the sorted slice.
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

// ArrContains checks if all elements in slice b are included in slice a.
//
// Parameters:
// - a: The first slice to check.
// - b: The second slice to check.
// Return:
// - bool: True if all elements in slice b are included in slice a, false otherwise.
func ArrContains[T comparable](a []T, b []T) bool {
	if len(b) == 0 {
		return true
	}

	for i := range b {
		v := b[i]
		if !Includes(v, a...) {
			return false
		}
	}

	return true
}

// ContainsAll checks if sliceA contains all elements of sliceB.
//
// sliceA: The first slice to check.
// sliceB: The elements to check for in sliceA.
// bool: Returns true if sliceA contains all elements of sliceB, false otherwise.
func ContainsAll[T comparable](sliceA []T, sliceB ...T) bool {
	return ArrContains(sliceA, sliceB)
}

// MergeFn is a function that takes two parameters of type T and returns a value of type T.
// It merges the current value with the next value using the Merge method of the type T.
//
// Parameters:
// - cur: the current value of type T.
// - next: the next value of type T.
//
// Return type: T
func MergeFn[T types.MergingHandler[T]](cur T, next T) T {
	return cur.Merge(next)
}

// ReduceMergeFn is a Go function that takes a default value and a slice of values of type T,
// and returns a value of type T. It uses the Reduce function with the MergeFn merging handler
// to reduce the slice of values to a single value.
//
// Parameters:
// - def: the default value of type T.
// - slice: the slice of values of type T.
//
// Return type:
// - T: the reduced value of type T.
func ReduceMergeFn[T types.MergingHandler[T]](def T, slice ...T) T {
	return Reduce(def, MergeFn[T], slice...)
}

// GetRandomArray returns a random array of size 'size' from the given 'items' array.
//
// Parameters:
// - items: The array from which the random elements are selected.
// - size: The number of elements to be selected for the random array.
//
// Return type:
// - []T: The random array of size 'size'.
func GetRandomArray[T any](items []T, size int) []T {
	if size >= len(items) {
		return items
	}

	// Shuffle the array using Fisher-Yates algorithm with crypto/rand
	shuffledItems := Shuffle(items)

	// Select the first 'size' elements as the random array
	return shuffledItems[:size]
}

// Shuffle shuffles the elements in the given slice in a random order.
//
// Parameters:
// - arr: the slice to be shuffled.
//
// Return:
// - []T: the shuffled slice.
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

// CryptoRandInt generates a random integer within the specified range using cryptographic random number generator.
//
// It takes an integer parameter 'max' which represents the upper bound of the range.
// It returns an integer value and an error. The integer value is a random number within the range [0, max).
// If 'max' is less than or equal to 0, it returns an error with the message "max should be greater than 0".
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

// Loop iterates over a slice of any type and invokes the callback function for each item.
//
// The callback function takes two parameters: the index of the item in the slice and the item itself.
// The items parameter is a variadic parameter of type T, which represents the slice of any type.
func Loop[T any](callback func(index int, item T), items ...T) {
	for i := 0; i < len(items); i++ {
		callback(i, items[i])
	}
}

// ILoop iterates over a slice of any type and calls the provided callback function for each item.
//
// items: the slice to iterate over.
// callback: the function to call for each item, takes the index and the item as parameters.
func ILoop[T any](items []T, callback func(index int, item T)) {
	for i := 0; i < len(items); i++ {
		callback(i, items[i])
	}
}

type PipeFn[T any] func(T) T

// Pipe applies a series of functions to an item in a pipeline fashion.
//
// item: The item to be processed by the pipeline.
// callbacks: The functions to be applied to the item in the order they are provided.
// T: The type of the item and the return type of the pipeline.
func Pipe[T any](item T, callbacks ...PipeFn[T]) T {
	v := item

	for i := 0; i < len(callbacks); i++ {
		v = callbacks[i](v)
	}

	return v
}

// SPipe is a function that takes a slice of any type and a variable number of PipeFn[T] callbacks.
//
// It iterates over the items in the slice and applies the Pipe function with the callbacks to each item.
// The results are collected into a new slice, which is then returned.
//
// Parameters:
// - items: The slice of any type.
// - callbacks: The variable number of PipeFn[T] callbacks.
//
// Returns:
// - []T: The resulting slice after applying the Pipe function to each item.
func SPipe[T any](items []T, callbacks ...PipeFn[T]) []T {
	var result []T

	for i := 0; i < len(items); i++ {
		result = append(result, Pipe(items[i], callbacks...))
	}

	return result
}

// Divide slices the given items into smaller slices of size to.
// It returns a slice of slices where each inner slice has a maximum size of to.
func Divide[T any](to int, items ...T) [][]T {
	// Calculate the number of smaller slices needed
	numSlices := (len(items) + to - 1) / to

	// Initialize the result slice of slices
	result := make([][]T, numSlices)

	// Iterate over the items and split them into smaller slices
	for i := 0; i < len(items); i += to {
		end := i + to
		if end > len(items) {
			end = len(items)
		}
		result[i/to] = items[i:end]
	}

	return result
}

func ExcludeByIndex[T any](arr []T, indices []int) []T {
	excluded := []T{}

	for index, item := range arr {
		if !IIncludes(indices, index) {
			excluded = append(excluded, item)
		}
	}

	return excluded
}

// Intersection finds the intersection of two integer slices.
func Intersection[T comparable](a, b []T) []T {
	set := make(map[T]*struct{})
	for _, v := range a {
		set[v] = new(struct{})
	}

	var result []T
	for _, v := range b {
		if _, exists := set[v]; exists {
			result = append(result, v)
		}
	}

	return result
}
