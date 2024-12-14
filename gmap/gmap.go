package gmap

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Merge 合并多个map
//
// 注意：如果有相同的key，后面的map会覆盖前面的map
//
// 示例：
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"b": 3, "c": 4}
//	m3 := map[string]int{"c": 5, "d": 6}
//
//	result := Merge(m1, m2, m3)
//	// result: map[string]int{"a": 1, "b": 3, "c": 4, "d": 6}
//
//	// 也可以使用可变参数
//
//	result := Merge(m1, m2, m3)
//	// result: map[string]int{"a": 1, "b": 3, "c": 4, "d": 6}
//
//	// 也可以使用切片
//
//	maps := []map[string]int{m1, m2, m3}
//	result := Merge(maps...)
//	// result: map[string]int{"a": 1, "b": 3, "c": 4, "d": 6}
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Clear 清空map
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

func Clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}
