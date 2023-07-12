package generic

// Zero returns a zero value of type T.
func Zero[T any]() T {
	var (
		t T
	)

	return t
}

// IsZero returns true if the value is zero.
func IsZero[T comparable](v T) bool {
	return v == Zero[T]()
}
