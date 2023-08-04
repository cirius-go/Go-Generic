package slice

type E[T any] []T

func (e E[T]) Concat(slices ...[]T) E[T] {
	elems := Concat(slices...)
	return append(e, elems...)
}

func (e E[T]) Unshift(elem T) E[T] {
	return Unshift[T](elem, e...)
}

func (e E[T]) Reverse() E[T] {
	return Reverse[T](e...)
}

func (e E[T]) FindIndex(f func(T) bool) int {
	return FindIndex[T](f, e...)
}

func (e E[T]) Find(f func(T) bool) (elem T, found bool) {
	return Find[T](f, e...)
}

func (e E[T]) Filter(predicate func(T) bool, items ...T) E[T] {
	return Filter(predicate, items...)
}

func (e E[T]) Every(predicate func(T) bool, items ...T) bool {
	return Every[T](predicate, items...)
}

func (e E[T]) Some(predicate func(T) bool, items ...T) bool {
	return Some[T](predicate, items...)
}

func (e E[T]) At(index int) (T, bool) {
	return At[T](index, e...)
}
