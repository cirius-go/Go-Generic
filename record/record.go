package record

// FindKeysByValue returns all keys that have a given value
func FindKeysByValue[K, V comparable](m map[K]V, values ...V) []K {
	result := make([]K, 0)

	for _, v := range values {
		for k := range m {
			if v == m[k] {
				result = append(result, k)
			}
		}
	}

	return result
}

// Keys returns all keys.
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0)

	for k := range m {
		result = append(result, k)
	}

	return result
}

// Vals returns all values.
func Vals[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0)

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func ValByKeys[K comparable, V any](m map[K]V, keys ...K) []V {
	result := make([]V, 0)

	for _, k := range keys {
		if v, ok := m[k]; ok {
			result = append(result, v)
		}
	}

	return result
}

type ValidKeyCondFn[K comparable] func(k K) bool

func ValsByKeyConds[K comparable, V any](m map[K]V, keyConds ...func(K) bool) []V {
	result := make([]V, 0)

	for k := range m {
		valid := true
		for _, cond := range keyConds {
			if !cond(k) {
				valid = false
				break
			}
		}

		if valid {
			result = append(result, m[k])
		}
	}

	return result
}

func Reduce[R any, K comparable, V any](init R, fn func(R, K, V) R, m map[K]V) R {
	for k, v := range m {
		init = fn(init, k, v)
	}

	return init
}
