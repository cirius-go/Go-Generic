package generic

// Select select something based on condition.
func Select[T any](a, b T, selectBOpts ...bool) T {
	for _, selectB := range selectBOpts {
		if selectB {
			return b
		}
	}

	return a
}

// ValueOrInitPointer return v if v != nil, otherwise return a new instance of T
func ValueOrInitPointer[T any](v *T) *T {
	if v == nil {
		return new(T)
	}

	return v
}

// Ptr return *T
func Ptr[T any](v T) *T {
	return &v
}

// FromPtr return T
func FromPtr[T any](v *T) T {
	return *v
}
