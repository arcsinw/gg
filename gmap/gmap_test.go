package gmap

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	map1 := map[int]string{1: "a", 2: "b"}
	map2 := map[int]string{3: "c", 4: "d"}
	expected := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	result := Merge(map1, map2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Merge failed, expected %v, got %v", expected, result)
	}

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m3 := map[string]int{"c": 5, "d": 6}
	exp := map[string]int{"a": 1, "b": 3, "c": 5, "d": 6}
	res := Merge(m1, m2, m3)
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Merge failed, expected %v, got %v", expected, result)
	}
}

func TestClear(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	Clear(m)
	if len(m) != 0 {
		t.Errorf("Clear failed, expected empty map, got %v", m)
	}
}

func TestClone(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	cloned := Clone(m)
	if !reflect.DeepEqual(cloned, m) {
		t.Errorf("Clone failed, expected %v, got %v", m, cloned)
	}
}

func TestGetOrDefault(t *testing.T) {
	m := map[int]string{1: "a"}
	value := GetOrDefault(m, 1, "default")
	if value != "a" {
		t.Errorf("GetOrDefault failed, expected 'a', got %v", value)
	}
	value = GetOrDefault(m, 2, "default")
	if value != "default" {
		t.Errorf("GetOrDefault failed, expected 'default', got %v", value)
	}
}

func TestKeys(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	expectedKeys := []int{1, 2, 3}
	keys := Keys(m)

	sort.Ints(keys)

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Keys failed, expected %v, got %v", expectedKeys, keys)
	}
}

func TestValues(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	expectedValues := []string{"a", "b", "c"}
	values := Values(m)
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Values failed, expected %v, got %v", expectedValues, values)
	}
}

func TestMap(t *testing.T) {
	// 测试空 map
	emptyMap := make(map[int]string)
	emptyResult := Map(emptyMap, func(k int, v string) (int, string) {
		return k, v
	})
	if len(emptyResult) != 0 {
		t.Errorf("Expected empty map, got %v", emptyResult)
	}

	// 测试包含一个键值对的 map
	singleMap := map[int]string{1: "one"}
	singleResult := Map(singleMap, func(k int, v string) (int, string) {
		return k, v
	})
	expectedSingleResult := map[int]string{1: "one"}
	if !reflect.DeepEqual(singleResult, expectedSingleResult) {
		t.Errorf("Expected %v, got %v", expectedSingleResult, singleResult)
	}

	// 测试包含多个键值对的 map
	multiMap := map[int]string{1: "one", 2: "two", 3: "three"}
	multiResult := Map(multiMap, func(k int, v string) (int, string) {
		return k, v
	})
	expectedMultiResult := map[int]string{1: "one", 2: "two", 3: "three"}
	if !reflect.DeepEqual(multiResult, expectedMultiResult) {
		t.Errorf("Expected %v, got %v", expectedMultiResult, multiResult)
	}

	// 测试使用不同的数据类型作为键和值
	stringMap := map[string]int{"one": 1, "two": 2, "three": 3}
	stringResult := Map(stringMap, func(k string, v int) (string, int) {
		return k, v
	})
	expectedStringResult := map[string]int{"one": 1, "two": 2, "three": 3}
	if !reflect.DeepEqual(stringResult, expectedStringResult) {
		t.Errorf("Expected %v, got %v", expectedStringResult, stringResult)
	}

	// 测试使用不同的函数 f 来验证映射逻辑
	transformMap := map[int]int{1: 10, 2: 20, 3: 30}
	transformResult := Map(transformMap, func(k int, v int) (int, int) {
		return k + 1, v * 2
	})
	expectedTransformResult := map[int]int{2: 20, 3: 40, 4: 60}
	if !reflect.DeepEqual(transformResult, expectedTransformResult) {
		t.Errorf("Expected %v, got %v", expectedTransformResult, transformResult)
	}
}
