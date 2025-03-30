package gslice

import (
	"sort"
)

//TIP using empty slice as default value of slice, not nil

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Ordered interface {
	Number | ~string
}

func IsNotZeroFunc[T Number](t T) bool {
	return t != 0
}

func IsPositiveFunc[T Number](t T) bool {
	return t > 0
}

func IsNegativeFunc[T Number](t T) bool {
	return t < 0
}

// Sort sort the elements in the slice in place
func Sort[T Ordered](slice []T) []T {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

// Chunk divides the slice into smaller slices, each containing at most 'size' elements
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

// Pop removes and returns the last element from the slice and new slice
func Pop[T any](slice []T) (T, []T) {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue, slice
	}
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// Append appends elements to the end of a slice and returns a new slice
func Append[T any](slice []T, elems ...T) []T {
	return append(slice, elems...)
}

// Prepend append one or more items to the beginning of slice and returns a new slice
func Prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}

// Insert inserts elements into a slice at a specified index and returns a new slice
func Insert[T any](slice []T, index int, elems ...T) []T {
	if index < 0 || index > len(slice) {
		return slice
	}
	return append(slice[:index], append(elems, slice[index:]...)...)
}

// Remove remove an element from a slice at a given index
func Remove[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// Map apply function f to each element of a slice and return a new slice
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

// First return first element which match condition
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

// Reverse reverse elements in slice and returns a new slice
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

// Count return count of elements which match the condition
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

// OrderBy return a new slice sorted by less function
func OrderBy[T any](slice []T, less func(T, T) bool) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.SliceStable(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// Min return the min value of slice
func Min[T any](slice []T, less func(T, T) bool) T {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue
	}

	minValue := slice[0]
	for _, v := range slice {
		if less(v, minValue) {
			minValue = v
		}
	}

	return minValue
}

// Max return the max value of slice
func Max[T any](slice []T, less func(T, T) bool) T {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue
	}

	maxValue := slice[0]
	for _, v := range slice {
		if less(maxValue, v) {
			maxValue = v
		}
	}

	return maxValue
}

// Sum return the sum of slice
func Sum[T any, E Number](slice []T, f func(T) E) E {
	var sum E
	for _, elem := range slice {
		sum += f(elem)
	}

	return sum
}

// ForEach apply f to each element of slice
func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

// Flatten flattens a slice of slices into a single slice
func Flatten[T any](slices [][]T) []T {
	result := make([]T, 0)
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// Uniq remove duplicate elements from slice
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

// UniqBy remove duplicate elements from slice by keyFunc
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

// Concat concat multi slice to a new slice
func Concat[S ~[]E, E any](slices ...S) S {
	var result []E
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// Contains checks if slice contains a specific value
func Contains[T comparable](s []T, v T) bool {
	for _, elem := range s {
		if elem == v {
			return true
		}
	}
	return false
}

// Filter return elements in slice that match the given condition
func Filter[T any](slice []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduce slice to a single value
func Reduce[T any](slice []T, f func(T, T) T) T {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue
	}
	var result T
	for i := 0; i < len(slice); i++ {
		result = f(result, slice[i])
	}
	return result
}

// GroupBy groups the elements in the slice according to the specified key function
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

// ToMap converts a slice into a map using a specified function to extract keys and values.
func ToMap[T, V any, K comparable](slice []T, f func(T) (K, V)) map[K]V {
	result := make(map[K]V)
	for _, item := range slice {
		key, value := f(item)
		result[key] = value
	}
	return result
}
