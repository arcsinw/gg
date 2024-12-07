package gslice

import "sort"

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128
}

//TIP 处理nil的slice输入

// Map 对切片中的每个元素应用一个函数，并返回一个新的切片
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// First 返回满足条件的第一个元素
func First[T any](slice []T, condition func(T) bool) (T, bool) {
	for _, v := range slice {
		if condition(v) {
			return v, true
		}
	}

	var zeroValue T
	return zeroValue, false
}

// Last return the last element which match condition
func Last[T any](slice []T, condition func(T) bool) (T, bool) {
	var zeroValue T
	if len(slice) == 0 {
		return zeroValue, false
	}

	for i := len(slice) - 1; i >= 0; i-- {
		if condition(slice[i]) {
			return slice[i], true
		}
	}

	return zeroValue, false
}

func Count[T any](slice []T, condition func(T) bool) int64 {
	count := int64(0)
	for _, v := range slice {
		if condition(v) {
			count++
		}
	}
	return count
}

// AllMatch return true if all elements match the condition
func AllMatch[T any](slice []T, condition func(T) bool) bool {
	if len(slice) == 0 {
		return true
	}

	for _, v := range slice {
		if !condition(v) {
			return false
		}
	}

	return true
}

// AnyMatch return if any element match the condition
func AnyMatch[T any](slice []T, condition func(T) bool) bool {
	for _, v := range slice {
		if condition(v) {
			return true
		}
	}

	return false
}

func OrderBy[T any](slice []T, less func(T, T) bool) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.SliceStable(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

func Min[T any](slice []T, less func(T, T) bool) T {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue
	}

	min := slice[0]
	for _, v := range slice {
		if less(v, min) {
			min = v
		}
	}

	return min
}

func Max[T any](slice []T, less func(T, T) bool) T {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue
	}

	max := slice[0]
	for _, v := range slice {
		if less(max, v) {
			max = v
		}
	}

	return max
}

func Sum[T any, E Number](slice []T, f func(T) E) E {
	var sum E
	for _, elem := range slice {
		sum += f(elem)
	}

	return sum
}

func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

// Flatten 将二维切片扁平化为一维切片
func Flatten[T any](slices [][]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func Uniq[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}

	var result []T

	m := map[T]interface{}{}
	for _, elem := range slice {
		k := elem
		if _, ok := m[k]; ok {
			continue
		} else {
			m[k] = elem
		}

		result = append(result, elem)
	}

	return result
}

func UniqBy[T any, K comparable](slice []T, keyFunc func(T) K) []T {
	if len(slice) == 0 {
		return slice
	}

	var result []T

	m := map[K]interface{}{}
	for _, elem := range slice {
		k := keyFunc(elem)
		if _, ok := m[k]; ok {
			continue
		} else {
			m[k] = elem
		}

		result = append(result, elem)
	}

	return result
}

func Concat[S ~[]E, E any](slices ...S) S {
	var result []E
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	for _, elem := range s {
		if elem == v {
			return true
		}
	}
	return false
}

// Filter 根据条件过滤切片中的元素，返回一个新的切片
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce 对切片中的元素进行累积操作，返回一个单一的值
func Reduce[T any](slice []T, f func(T, T) T) T {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue
	}
	result := slice[0]
	for i := 1; i < len(slice); i++ {
		result = f(result, slice[i])
	}
	return result
}

func ToMap[T, V any, K comparable](slice []T, f func(T) (K, V)) map[K]V {
	result := make(map[K]V)
	for _, item := range slice {
		key, value := f(item)
		result[key] = value
	}
	return result
}
