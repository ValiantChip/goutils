package pointer

func New[T any](value T) *T {
	return &value
}

func ZeroIfNil[T any](value *T) T {
	if value == nil {
		var zero T
		return zero
	}
	return *value
}

func ValIfNil[T any](value *T, nValue T) T {
	if value == nil {
		return nValue
	}
	return *value
}

func Equals[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
