package gslice

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestMap 测试 Map 函数
func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []string{"1", "2", "3", "4"}
	strFunc := func(i int) string { return string(rune(i + '0')) }

	result := Map(input, strFunc)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map was incorrect, got %v, expected %v", result, expected)
	}
}

// TestMapEmptySlice 测试 Map 函数处理空切片
func TestMapEmptySlice(t *testing.T) {
	var input []int
	expected := make([]string, 0)

	result := Map(input, func(i int) string { return string(rune(i + '0')) })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map was incorrect, got %v, expected %v", result, expected)
	}
}

func TestMapNonEmptySlice(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 4, 9}
	result := Map(input, func(v int) int { return v * v })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMapNilSlice(t *testing.T) {
	var input []int
	result := Map(input, func(v int) int { return v })
	if len(result) != 0 {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestMapNilFunction(t *testing.T) {
	input := []int{1, 2, 3}
	var f func(int) int = nil
	result := Map(input, f)
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}
}

func TestMapFunctionReturnsDifferentType(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []string{"1", "2", "3"}
	result := Map(input, func(v int) string { return string(rune(v + '0')) })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMapLargeSlice(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	expected := make([]int, 1000)
	for i := range expected {
		expected[i] = i * i
	}
	result := Map(input, func(v int) int { return v * v })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestOrderBy 测试 OrderBy 函数
func TestOrderBy(t *testing.T) {
	input := []int{3, 1, 4, 2}
	expected := []int{1, 2, 3, 4}
	lessFunc := func(a, b int) bool { return a < b }

	result := OrderBy(input, lessFunc)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("OrderBy was incorrect, got %v, expected %v", result, expected)
	}
}

// TestSum 测试 Sum 函数
func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := 10
	sumFunc := func(i int) int { return i }

	result := Sum(input, sumFunc)
	if result != expected {
		t.Errorf("Sum was incorrect, got %v, expected %v", result, expected)
	}
}

// TestFlatten 测试 Flatten 函数
func TestFlatten(t *testing.T) {
	input := [][]int{{1, 2}, {3, 4}}
	expected := []int{1, 2, 3, 4}

	result := Flatten(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Flatten was incorrect, got %v, expected %v", result, expected)
	}
}

// TestUniqBy 测试 UniqBy 函数
func TestUniqBy(t *testing.T) {
	input := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	result := UniqBy(input, func(i int) int { return i })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("UniqBy was incorrect, got %v, expected %v", result, expected)
	}
}

// TestConcat 测试 Concat 函数
func TestConcat(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}
	expected := []int{1, 2, 3, 4}

	result := Concat(input1, input2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Concat was incorrect, got %v, expected %v", result, expected)
	}
}

// TestContains 测试 Contains 函数
func TestContains(t *testing.T) {
	input := []int{1, 2, 3}
	value := 2

	result := Contains(input, value)
	if !result {
		t.Errorf("Contains was incorrect, got %v, expected true", result)
	}
}

// TestFilter 测试 Filter 函数
func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{2, 4}
	isEven := func(i int) bool { return i%2 == 0 }

	result := Filter(input, isEven)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter was incorrect, got %v, expected %v", result, expected)
	}
}

// TestReduce 测试 Reduce 函数
func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := 10
	addFunc := func(a, b int) int { return a + b }

	result := Reduce(input, addFunc)
	if result != expected {
		t.Errorf("Reduce was incorrect, got %v, expected %v", result, expected)
	}
}

// Person 结构体用于测试
type Person struct {
	Name string
	Age  int
}

// TestUniqByObjectSlice 测试 UniqBy 函数处理对象切片
func TestUniqByObjectSlice(t *testing.T) {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Alice", 30}, // 重复的Alice
		{"Carol", 20},
	}
	expected := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 20},
	}

	result := UniqBy(people, func(p Person) string { return p.Name })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("UniqBy was incorrect, got %v, expected %v", result, expected)
	}
}

// TestFilterObjectSlice 测试 Filter 函数处理对象切片
func TestFilterObjectSlice(t *testing.T) {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 20},
		{"Dave", 22},
	}
	expected := []Person{
		{"Alice", 30},
	}
	isAdult := func(p Person) bool { return p.Age >= 30 }

	result := Filter(people, isAdult)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter was incorrect, got %v, expected %v", result, expected)
	}
}

// TestMapObjectSlice 测试 Map 函数处理对象切片
func TestMapObjectSlice(t *testing.T) {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 20},
	}
	expected := []string{"Alice", "Bob", "Carol"}

	result := Map(people, func(p Person) string { return p.Name })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map was incorrect, got %v, expected %v", result, expected)
	}
}

// TestOrderByEmptySlice 测试 OrderBy 函数处理空切片
func TestOrderByEmptySlice(t *testing.T) {
	var input []int
	expected := make([]int, 0)

	result := OrderBy(input, func(a, b int) bool { return a < b })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("OrderBy was incorrect, got %v, expected %v", result, expected)
	}
}

// TestSumEmptySlice 测试 Sum 函数处理空切片
func TestSumEmptySlice(t *testing.T) {
	var input []int
	var expected int = 0

	result := Sum(input, func(i int) int { return i })
	if result != expected {
		t.Errorf("Sum was incorrect, got %v, expected %v", result, expected)
	}
}

// TestFlattenEmptySliceOfSlices 测试 Flatten 函数处理空切片
func TestFlattenEmptySliceOfSlices(t *testing.T) {
	var slices [][]int
	var expected []int

	result := Flatten(slices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Flatten was incorrect, got %v, expected %v", result, expected)
	}
}

// TestUniqByEmptySlice 测试 UniqBy 函数处理空切片
func TestUniqByEmptySlice(t *testing.T) {
	var input []int
	var expected []int

	result := UniqBy(input, func(i int) int { return i })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("UniqBy was incorrect, got %v, expected %v", result, expected)
	}
}

// TestConcatEmptySlices 测试 Concat 函数处理空切片
func TestConcatEmptySlices(t *testing.T) {
	var input1 []int
	var input2 []int
	var expected []int

	result := Concat(input1, input2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Concat was incorrect, got %v, expected %v", result, expected)
	}
}

// TestContainsEmptySlice 测试 Contains 函数处理空切片
func TestContainsEmptySlice(t *testing.T) {
	var input []int
	value := 2

	result := Contains(input, value)
	if result {
		t.Errorf("Contains was incorrect, got %v, expected false", result)
	}
}

// TestFilterEmptySlice 测试 Filter 函数处理空切片
func TestFilterEmptySlice(t *testing.T) {
	var input []int
	var expected []int

	result := Filter(input, func(i int) bool { return i > 0 })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter was incorrect, got %v, expected %v", result, expected)
	}
}

// TestReduceEmptySlice 测试 Reduce 函数处理空切片
func TestReduceEmptySlice(t *testing.T) {
	var input []int
	expected := 0

	result := Reduce(input, func(a, b int) int { return a + b })
	if result != expected {
		t.Errorf("Reduce was incorrect, got %v, expected %v", result, expected)
	}
}

// TestToMapNormal 测试 ToMap 函数处理普通情况
func TestToMapNormal(t *testing.T) {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 22},
	}
	expectedMap := map[string]Person{
		"Alice": {"Alice", 30},
		"Bob":   {"Bob", 25},
		"Carol": {"Carol", 22},
	}
	actualMap := ToMap(people, func(p Person) (string, Person) {
		return p.Name, p
	})
	if !reflect.DeepEqual(actualMap, expectedMap) {
		t.Errorf("ToMap was incorrect, got %v, expected %v", actualMap, expectedMap)
	}
}

// TestToMapEmptySlice 测试 ToMap 函数处理空切片
func TestToMapEmptySlice(t *testing.T) {
	var emptyPeople []Person
	expectedEmptyMap := map[string]Person{}
	actualEmptyMap := ToMap(emptyPeople, func(p Person) (string, Person) {
		return p.Name, p
	})
	if !reflect.DeepEqual(actualEmptyMap, expectedEmptyMap) {
		t.Errorf("ToMap was incorrect, got %v, expected %v", actualEmptyMap, expectedEmptyMap)
	}
}

// TestToMapNilSlice 测试 ToMap 函数处理 nil 切片
func TestToMapNilSlice(t *testing.T) {
	var nilSlice []Person = nil
	result := ToMap(nilSlice, func(item Person) (string, Person) {
		return item.Name, item
	})
	if len(result) != 0 {
		t.Errorf("Expected nil result for nil slice, got %v", result)
	}
}

// TestToMapDuplicateKeys 测试 ToMap 函数处理重复键
func TestToMapDuplicateKeys(t *testing.T) {
	duplicatePeople := []Person{
		{"Alice", 30},
		{"Alice", 25}, // 重复键
		{"Carol", 22},
	}
	actualDuplicateMap := ToMap(duplicatePeople, func(p Person) (string, Person) {
		return p.Name, p
	})
	if len(actualDuplicateMap) != 2 {
		t.Errorf("ToMap with duplicate keys was incorrect, got %v", actualDuplicateMap)
	}
	if _, ok := actualDuplicateMap["Alice"]; !ok {
		t.Errorf("ToMap with duplicate keys did not contain the last key 'Alice'")
	}
}

func TestFirst(t *testing.T) {
	// 测试找到匹配元素的情况
	slice := []int{1, 2, 3, 4, 5}
	condition := func(v int) bool { return v == 3 }
	result, found := First(slice, condition)
	if !found {
		t.Errorf("Expected to find a match, but found none")
	}
	if result != 3 {
		t.Errorf("Expected to find 3, but found %v", result)
	}

	// 测试未找到匹配元素的情况
	slice = []int{1, 2, 4, 5}
	condition = func(v int) bool { return v == 3 }
	result, found = First(slice, condition)
	if found {
		t.Errorf("Expected to not find a match, but found one")
	}
	if result != 0 {
		t.Errorf("Expected to find 0, but found %v", result)
	}

	// 测试空切片的情况
	slice = []int{}
	condition = func(v int) bool { return v == 3 }
	if found {
		t.Errorf("Expected not to find a match, but found one")
	}
	if result != 0 {
		t.Errorf("Expected to find 0, but found %v", result)
	}
	assert.Equal(t, 0, result) // 这里根据实际情况设置默认值
}

func TestLast(t *testing.T) {
	// 测试空切片
	result, found := Last([]int{}, func(i int) bool { return i > 0 })
	assert.Equal(t, 0, result)
	assert.False(t, found)

	// 测试切片中没有满足条件的元素
	result, found = Last([]int{1, 2, 3}, func(i int) bool { return i > 5 })
	assert.Equal(t, 0, result)
	assert.False(t, found)

	// 测试切片中有满足条件的元素
	result, found = Last([]int{1, 2, 3, 4, 5}, func(i int) bool { return i > 3 })
	assert.Equal(t, 5, result)
	assert.True(t, found)
}

func TestCount(t *testing.T) {
	// 测试用例 1：空切片
	slice1 := []int{}
	expectedCount1 := int64(0)
	actualCount1 := Count(slice1, func(v int) bool { return v > 10 })
	assert.Equal(t, expectedCount1, actualCount1, "Test case 1 failed")

	// 测试用例 2：满足条件的元素
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	expectedCount2 := int64(2)
	actualCount2 := Count(slice2, func(v int) bool { return v > 10 })
	assert.Equal(t, expectedCount2, actualCount2, "Test case 2 failed")

	// 测试用例 3：不满足条件的元素
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedCount3 := int64(0)
	actualCount3 := Count(slice3, func(v int) bool { return v > 10 })
	assert.Equal(t, expectedCount3, actualCount3, "Test case 3 failed")
}

func TestAllMatch(t *testing.T) {
	// 测试空切片
	assert.True(t, AllMatch([]int{}, func(v int) bool { return true }))

	// 测试所有元素都满足条件
	assert.True(t, AllMatch([]int{1, 2, 3}, func(v int) bool { return v > 0 }))

	// 测试有元素不满足条件
	assert.False(t, AllMatch([]int{1, 2, -1}, func(v int) bool { return v > 0 }))
}

func TestAnyMatch(t *testing.T) {
	// 测试空切片
	slice1 := []int{}
	condition1 := func(v int) bool { return v > 10 }
	assert.False(t, AnyMatch(slice1, condition1))

	// 测试切片中没有满足条件的元素
	slice2 := []int{1, 2, 3, 4, 5}
	condition2 := func(v int) bool { return v > 10 }
	assert.False(t, AnyMatch(slice2, condition2))

	// 测试切片中有满足条件的元素
	slice3 := []int{1, 2, 3, 4, 5, 11}
	condition3 := func(v int) bool { return v > 10 }
	assert.True(t, AnyMatch(slice3, condition3))
}

func TestMin(t *testing.T) {
	// 测试空切片
	emptySlice := []int{}
	minEmpty := Min(emptySlice, func(a, b int) bool { return a < b })
	assert.Equal(t, 0, minEmpty)

	// 测试非空切片
	slice := []int{5, 3, 9, 1, 7}
	minValue := Min(slice, func(a, b int) bool { return a < b })
	assert.Equal(t, 1, minValue)
}

func TestMax(t *testing.T) {
	// 测试空切片
	var emptySlice []int
	result := Max(emptySlice, func(a, b int) bool { return a < b })
	assert.Equal(t, 0, result)

	// 测试非空切片
	slice := []int{1, 3, 2}
	result = Max(slice, func(a, b int) bool { return a < b })
	assert.Equal(t, 3, result)
}

func TestForEach(t *testing.T) {
	// 测试空切片
	var emptySlice []int
	ForEach(emptySlice, func(v int) {
		t.Errorf("不应该执行回调函数")
	})

	// 测试非空切片
	slice := []int{1, 2, 3}
	var result []int
	ForEach(slice, func(v int) {
		result = append(result, v)
	})
	assert.Equal(t, slice, result)
}

func TestUniq(t *testing.T) {
	// 测试空切片
	emptySlice := []int{}
	expectedEmptySlice := []int{}
	assert.Equal(t, expectedEmptySlice, Uniq(emptySlice))

	// 测试包含重复元素的切片
	sliceWithDuplicates := []int{1, 2, 2, 3, 3, 3}
	expectedUniqueSlice := []int{1, 2, 3}
	assert.Equal(t, expectedUniqueSlice, Uniq(sliceWithDuplicates))

	// 测试不包含重复元素的切片
	uniqueSlice := []int{4, 5, 6}
	expectedSameSlice := []int{4, 5, 6}
	assert.Equal(t, expectedSameSlice, Uniq(uniqueSlice))
}
