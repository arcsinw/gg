package gptr

// Of 返回一个指向v的指针
func Of[T any](v T) *T {
	return &v
}

// IndirectOf 对指针进行解引用
func IndirectOf[T any](p *T) T {
	return *p
}

// IsNil 判断指针是否为空
func IsNil[T any](p *T) bool {
	return p == nil
}

// ValueOf 获取指针指向的值，如果指针为空，则返回默认值
func ValueOf[T any](p *T, defaultValue T) T {
	if p == nil {
		return defaultValue
	}
	return *p
}
