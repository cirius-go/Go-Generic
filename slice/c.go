package slice

type C[T comparable] []T

func (c C[T]) Concat(slices ...[]T) C[T] {
	elems := Concat(slices...)
	return append(c, elems...)
}

func (c C[T]) RemoveDuplicates() C[T] {
	return RemoveDuplicates[T](c...)
}

func (c C[T]) ConcatUnique(slices ...[]T) C[T] {
	return c.Concat(slices...).RemoveDuplicates()
}

func (c C[T]) Unshift(elem T) C[T] {
	return Unshift[T](elem, c...)
}

func (c C[T]) Reverse() C[T] {
	return Reverse[T](c...)
}

func (c C[T]) FindIndex(f func(T) bool) int {
	return FindIndex[T](f, c...)
}

func (c C[T]) Find(f func(T) bool) (elem T, found bool) {
	return Find[T](f, c...)
}

func (c C[T]) Filter(predicate func(T) bool, items ...T) C[T] {
	return Filter(predicate, items...)
}

func (c C[T]) Every(predicate func(T) bool, items ...T) bool {
	return Every[T](predicate, items...)
}

func (c C[T]) Some(predicate func(T) bool, items ...T) bool {
	return Some[T](predicate, items...)
}

func (c C[T]) At(index int) (T, bool) {
	return At[T](index, c...)
}
