package generic

// Select select something based on condition.
func Select[T comparable](a, b T, selectBOpts ...bool) T {
	for _, selectB := range selectBOpts {
		if selectB {
			return b
		}
	}

	return a
}
