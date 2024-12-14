package gptr

import (
	"reflect"
	"testing"
)

func TestOfFunctionality(t *testing.T) {
	// 基本功能测试
	var intValue int = 42
	intPtr := Of(intValue)
	if intPtr == nil || *intPtr != intValue {
		t.Errorf("Of should return a pointer to the value, got %v", intPtr)
	}

	// 类型保持测试
	var strValue string = "hello"
	strPtr := Of(strValue)
	if strPtr == nil || *strPtr != strValue {
		t.Errorf("Of should return a pointer to the value, got %v", strPtr)
	}

	// 并发安全测试
	done := make(chan bool)
	go func() {
		ptr := Of("concurrent test")
		if ptr == nil || *ptr != "concurrent test" {
			t.Errorf("Of should be concurrent safe, got %v", ptr)
		}
		done <- true
	}()
	<-done
}

func TestOfWithErrorUsage(t *testing.T) {
	// 错误使用测试
	var temp int
	tempPtr := Of(temp)
	temp = 0 // 改变temp的值，模拟变量被回收的情况

	// 这里我们不能直接测试panic，因为panic是运行时错误，不是返回值
	// 但我们可以通过检查tempPtr指向的值是否被改变来间接测试
	if *tempPtr != 0 {
		t.Errorf("The value pointed to by tempPtr should be changed to 0, got %v", *tempPtr)
	}
}

func TestOfTypePreservation(t *testing.T) {
	// 类型保持测试
	var intValue int = 42
	intPtr := Of(intValue)
	if reflect.TypeOf(intPtr) != reflect.PtrTo(reflect.TypeOf(intValue)) {
		t.Errorf("The type of the pointer returned by Of should be *int, got %v", reflect.TypeOf(intPtr))
	}
}

func BenchmarkOf(b *testing.B) {
	// 性能基准测试
	for i := 0; i < b.N; i++ {
		Of(i)
	}
}

func TestIsNil(t *testing.T) {
	// 基本功能测试
	var nilPtr *int
	if !IsNil(nilPtr) {
		t.Errorf("IsNil should return true for nil pointers, got false")
	}

	// 类型保持测试
	notNilPtr := new(int)
	if IsNil(notNilPtr) {
		t.Errorf("IsNil should return false for pointers to non-nil pointers, got true")
	}
}

func TestIndirectOf(t *testing.T) {
	var value int = 42
	ptr := &value
	result := IndirectOf(ptr)
	if result != value {
		t.Errorf("IndirectOf should return the value pointed to by the pointer, got %v", result)
	}

	var nilPtr *int
	result = IndirectOf(nilPtr)
	if result != 0 {
		t.Errorf("IndirectOf should return the zero value for nil pointers, got %v", result)
	}
}
