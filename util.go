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
