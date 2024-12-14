package gptr

// Of 返回一个指向v的指针
func Of[T any](v T) *T {
	return &v
}

// IndirectOf 对指针进行解引用，p为nil时返回类型的零值
func IndirectOf[T any](p *T) T {
	if p == nil {
		var zeroValue T
		return zeroValue
	}

	return *p
}

// IsNil 判断指针是否为空
func IsNil[T any](p *T) bool {
	return p == nil
}
