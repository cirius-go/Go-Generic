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

// Something represents something that not important.
type Something struct{}

// SomethingIntf represents something that used to implement purpose only.
type SomethingIntf interface{}

// Select select something based on condition.
func Select[T comparable](a, b T, selectBOpts ...bool) T {
	for _, selectB := range selectBOpts {
		if selectB {
			return b
		}
	}

	return a
}
