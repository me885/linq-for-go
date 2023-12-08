package linq

// Returns a new map with only the elements for which f() returns true
func MapWhere[K comparable, V any](m map[K]V, fn func(key K, value V) bool) map[K]V {
	newMap := map[K]V{}

	for k, v := range m {
		if fn(k, v) {
			newMap[k] = v
		}
	}

	return newMap
}

// Returns a new map with f() applied to every value element in m
func MapSelect[K comparable, In any, Out any](m map[K]In, f func(In) Out) map[K]Out {
	newMap := make(map[K]Out, len(m))

	for i, v := range m {
		newMap[i] = f(v)
	}

	return newMap
}

// Returns the only element of m and an error if it does not have exactly one element
func MapSingle[K comparable, V any](m map[K]V) (V, error) {
	var singleV V

	if len(m) != 1 {
		return singleV, ErrorNotMapOfOne
	}

	for _, v := range m {
		singleV = v
	}

	return singleV, nil
}
