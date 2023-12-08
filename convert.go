package linq

// Creates a map from the elements of a Slice.
// The key() function determines the key for each element
func SliceToMap[T any, K comparable](s []T, key func(index int, element T) K) map[K]T {
	newMap := make(map[K]T, len(s))

	for i, v := range s {
		newMap[key(i, v)] = v
	}

	return newMap
}

// Creates a slice from the values in a map
func MapToSlice[T any, K comparable](m map[K]T) []T {
	slice := make([]T, 0)

	for _, v := range m {
		slice = append(slice, v)
	}

	return slice
}

// Creates a slice from each key value pair in a map
func MapSelectToSlice[T any, K comparable, V any](m map[K]V, f func(K, V) T) []T {
	slice := make([]T, 0)

	for k, v := range m {
		slice = append(slice, f(k, v))
	}

	return slice
}

// Creates a slice of all the keys in the map
func MapKeysToSlice[T any, K comparable](m map[K]T) []K {
	slice := make([]K, 0)

	for k := range m {
		slice = append(slice, k)
	}

	return slice
}
