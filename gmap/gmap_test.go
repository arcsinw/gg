package gmap

import (
	"golang.org/x/exp/slices"
	"reflect"
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

	slices.Sort(keys)
	slices.Sort(expectedKeys)

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Keys failed, expected %v, got %v", expectedKeys, keys)
	}
}

func TestValues(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	expectedValues := []string{"a", "b", "c"}
	values := Values(m)
	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Values failed, expected %v, got %v", expectedValues, values)
	}
}
