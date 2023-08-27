package record

func FindKeysByValue[K, V comparable](m map[K]V, values ...V) []K {
	result := make([]K, 0)

	for _, v := range values {
		for k, _ := range m {
			if v == m[k] {
				result = append(result, k)
			}
		}
	}

	return result
}
