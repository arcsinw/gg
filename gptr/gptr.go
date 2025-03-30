package gptr

// Of return a pointer to v
func Of[T any](v T) *T {
	return &v
}

// IndirectOf returns the value it points to
// or the zero value of the type if the pointer is nil.
func IndirectOf[T any](p *T) T {
	if p == nil {
		var zeroValue T
		return zeroValue
	}

	return *p
}

// IsNil return ture if the pointer is nil
func IsNil[T any](p *T) bool {
	return p == nil
}
