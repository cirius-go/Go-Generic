package generic

import "github.com/cirius-go/generic/slice"

// SliceA represents the slice of any type item.
type SliceA[T any] []T

func NewSliceA[T any](initialVars ...T) SliceA[T] {
	s := make([]T, len(initialVars))

	return append(s, initialVars...)
}

func (s SliceA[T]) JoinSlices(tails ...[]T) SliceA[T] {
	slices := slice.Unshift([]T(s), tails...)

	return slice.Concat(slices...)
}

func (s SliceA[T]) Concat(tails ...[]T) SliceA[T] {
	return s.JoinSlices(tails...)
}

func (s SliceA[T]) Unshift(v T) SliceA[T] {
	return slice.Unshift(v, s...)
}

func (s SliceA[T]) Reverse() SliceA[T] {
	return slice.Reverse(s...)
}

func (s SliceA[T]) FindIndex(fn func(T) bool) (int, bool) {
	return slice.FindIndex(fn, s...)
}

func (s SliceA[T]) Find(fn func(T) bool) (T, bool) {
	return slice.Find(fn, s...)
}

func (s SliceA[T]) Filter(fn func(T) bool) SliceA[T] {
	return slice.Filter(fn, s...)
}

func (s SliceA[T]) Every(fn func(T) bool) bool {
	return slice.Every(fn, s...)
}

func (s SliceA[T]) Some(fn func(T) bool) bool {
	return slice.Some(fn, s...)
}

func (s SliceA[T]) MapAny(fn func(T) any) SliceA[any] {
	return slice.Map(fn, s...)
}

func (s SliceA[T]) MapOrigin(fn func(T) T) SliceA[T] {
	return slice.Map(fn, s...)
}

func (s SliceA[T]) Shift() SliceA[T] {
	return slice.Shift(s...)
}

func (s SliceA[T]) Pop() SliceA[T] {
	return slice.Pop(s...)
}

func (s SliceA[T]) At(index int) (T, bool) {
	return slice.At(index, s...)
}

func (s SliceA[T]) Clone() SliceA[T] {
	return slice.Clone(s)
}

func AnyTo[T any](tails ...any) (SliceA[T], bool) {
	result := NewSliceA[T]()
	for _, tail := range tails {
		t, ok := tail.(T)
		if !ok {
			return nil, false
		}

		result = append(result, t)
	}

	return result, true
}
