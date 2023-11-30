package generic

// Select select something based on condition.
func Select[T any](a, b T, selectBOpts ...bool) T {
	if len(selectBOpts) == 0 {
		return a
	}

	if selectBOpts[0] {
		return b
	}

	return a
}

// SelectA select something based on condition.
func SelectA[T any](a, b T, selectAOpts ...bool) T {
	if len(selectAOpts) == 0 {
		return a
	}

	if selectAOpts[0] {

		return b
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

// Must is a helper function that panics if an error is not nil.
// It returns the value v if the error is nil.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
