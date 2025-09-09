# gg

[![Go Reference](https://pkg.go.dev/badge/github.com/arcsinw/gg.svg)](https://pkg.go.dev/github.com/arcsinw/gg)

一个简单的Go语言工具库，提供了对切片(slice)、映射(map)和指针(pointer)的常用操作封装。

## 功能特性

- 支持Go 1.18及以上版本（基于泛型实现）
- 三个主要模块：
    - **gslice**: 提供丰富的切片操作函数
    - **gmap**: 提供实用的映射操作函数
    - **gptr**: 提供便捷的指针操作函数

## 安装

```bash
go get github.com/arcsinw/gg
```

## 使用示例

### gslice模块

```go
import "github.com/arcsinw/gg/gslice"

// 示例：使用Map函数转换切片元素
source := []int{1, 2, 3, 4, 5}
result := gslice.Map(source, func(v int) string {
    return fmt.Sprintf("item-%d", v)
})
// result: ["item-1", "item-2", "item-3", "item-4", "item-5"]

// 示例：使用Filter函数过滤切片元素
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evens := gslice.Filter(numbers, func(v int) bool {
    return v%2 == 0
})
// evens: [2, 4, 6, 8, 10]

// 示例：使用Chunk函数分割切片
slice := []int{1, 2, 3, 4, 5, 6, 7}
chunks := gslice.Chunk(slice, 3)
// chunks: [[1, 2, 3], [4, 5, 6], [7]]

// 示例：使用GroupBy函数分组
people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Carol", Age: 30},
}
groups := gslice.GroupBy(people, func(p Person) int {
    return p.Age
})
// groups: {30: [{Name: "Alice", Age: 30}, {Name: "Carol", Age: 30}], 25: [{Name: "Bob", Age: 25}]}
```

### gmap模块

```go
import "github.com/arcsinw/gg/gmap"

// 示例：获取map的所有键
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := gmap.Keys(m)
// keys: ["a", "b", "c"] (顺序可能随机)

// 示例：合并多个map
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
merged := gmap.Merge(m1, m2)
// merged: {"a": 1, "b": 3, "c": 4} (后面map的键值会覆盖前面的)

// 示例：获取值，如果键不存在则返回默认值
value := gmap.GetOrDefault(m, "d", 0)
// value: 0
```

### gptr模块

```go
import "github.com/arcsinw/gg/gptr"

// 示例：创建指针
p := gptr.Of(42)
// p: *int 类型，值为42

// 示例：安全地获取指针指向的值
var nilPtr *int
value := gptr.IndirectOf(nilPtr)
// value: 0 (int类型的零值)

// 示例：检查指针是否为nil
isNil := gptr.IsNil(nilPtr)
// isNil: true
```

## 关键API介绍

### gslice模块

#### 转换操作
- **Map[T, U any]**: 将一个切片的元素映射转换为另一个类型的切片
- **ToMap[T, V any, K comparable]**: 将切片转换为map
- **Flatten[T any]**: 将切片的切片展平为单一切片

#### 过滤和搜索
- **Filter[T any]**: 过滤出满足条件的元素
- **First[T any]**: 找出第一个满足条件的元素
- **FirstIndex[T any]**: 找出第一个满足条件的元素的索引
- **Contains[T comparable]**: 检查切片是否包含指定元素

#### 聚合操作
- **Reduce[T any]**: 将切片归约为单个值
- **Sum[T any, E Number]**: 计算切片元素的总和
- **Count[T any]**: 计算满足条件的元素数量
- **GroupBy[T any, K comparable]**: 根据指定的键函数对元素进行分组

#### 操作和修改
- **Chunk[T any]**: 将切片分割成指定大小的子切片
- **Uniq[T comparable]**: 去除切片中的重复元素
- **UniqBy[T any, K comparable]**: 根据键函数去除切片中的重复元素
- **Sort[T Ordered]**: 对切片进行排序
- **OrderBy[T any]**: 根据自定义比较函数对切片进行排序
- **Reverse[T any]**: 反转切片元素

#### 检查和判断
- **AllMatch[T any]**: 检查是否所有元素都满足条件
- **AnyMatch[T any]**: 检查是否存在满足条件的元素

### gmap模块

- **Keys[K comparable, V any]**: 获取map的所有键
- **Values[K comparable, V any]**: 获取map的所有值
- **Map[K comparable, V any]**: 对map的每个元素应用函数并返回新的map
- **Merge[K comparable, V any]**: 合并多个map
- **Clear[K comparable, V any]**: 清空map
- **Clone[K comparable, V any]**: 创建map的浅拷贝
- **GetOrDefault[K comparable, V any]**: 获取键对应的值，如果键不存在则返回默认值

### gptr模块

- **Of[T any]**: 创建指向值的指针
- **IndirectOf[T any]**: 获取指针指向的值，如果指针为nil则返回类型的零值
- **IsNil[T any]**: 检查指针是否为nil

