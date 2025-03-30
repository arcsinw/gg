package gmap

// Keys return keys of map as a slice (in random sort)
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values return values of map as a slice
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Map apply function f to each element of a map and return a new map
func Map[K comparable, V any](m map[K]V, f func(K, V) (K, V)) map[K]V {
	result := make(map[K]V, len(m))
	for k1, v1 := range m {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}
	return result
}

// Merge merges multiple maps into one.
//
// if the same key exists in multiple maps, the value in the last map will be used.
//
// example:
//
//		m1 := map[string]int{"a": 1, "b": 2}
//		m2 := map[string]int{"b": 3, "c": 4}
//		m3 := map[string]int{"c": 5, "d": 6}
//
//		result := Merge(m1, m2, m3)
//		// result: map[string]int{"a": 1, "b": 3, "c": 5, "d": 6}
//
//	 // or use slice
//
//		maps := []map[string]int{m1, m2, m3}
//		result := Merge(maps...)
//		// result: map[string]int{"a": 1, "b": 3, "c": 5, "d": 6}
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Clear clears the map.
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

// Clone creates a shallow copy of the map.
func Clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// GetOrDefault retrieves the value corresponding to the key from the map.
// If the key does not exist, it returns the default value.
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}
