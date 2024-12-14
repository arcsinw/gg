package gslice

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128
}

//TIP 处理nil的slice输入

//TIP 切片的默认值使用空切片而不是nil

// Shift: 移除并返回切片的第一个元素。
// Index: 返回切片中第一个匹配元素的索引。

// Reduce: 对切片中的元素进行累积操作，并返回一个单一的值

// Sort 对切片中的元素进行原地排序
func Sort[T constraints.Ordered](slice []T) []T {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

// Chunk 将切片分割成指定大小的子切片
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	result := make([][]T, 0)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

// Pop 移除并返回切片的最后一个元素
func Pop[T any](slice []T) (T, []T) {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue, slice
	}
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// Append 将一个或多个元素添加到切片的末尾
func Append[T any](slice []T, elems ...T) []T {
	return append(slice, elems...)
}

// Prepend 将一个或多个元素添加到切片的开头
func Prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}

// Insert 在指定位置插入一个或多个元素
func Insert[T any](slice []T, index int, elems ...T) []T {
	if index < 0 || index > len(slice) {
		return slice
	}
	return append(slice[:index], append(elems, slice[index:]...)...)
}

// Remove 从切片中移除指定位置的元素
func Remove[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// Map 对切片中的每个元素应用一个函数，并返回一个新的切片
func Map[T, U any](slice []T, f func(T) U) []U {
	if f == nil {
		return []U{}
	}

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

func FirstIndex[T any](slice []T, condition func(T) bool) (int, bool) {
	if len(slice) == 0 || condition == nil {
		return -1, false
	}

	for i, v := range slice {
		if condition(v) {
			return i, true
		}
	}
	return -1, false
}

// Reverse 反转切片的元素顺序，返回一个新的切片
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-i-1] = v
	}
	return result
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
	result := make([]T, 0)
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

func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	result := make(map[K][]T)
	if keyFunc == nil {
		return result
	}

	for _, item := range slice {
		key := keyFunc(item)
		result[key] = append(result[key], item)
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
