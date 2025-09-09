package pointer

// Returns a pointer that points to the value provided
func New[T any](value T) *T {
	return &value
}

// Returns the value that the pointer points to
// If the pointer is nil instead returns the zero value of the type the pointer points to
func ZeroIfNil[T any](ptr *T) T {
	var zero T
	return ValIfNil(ptr, zero)
}

// Returns the value that the pointer points to
// If the pointer is nil instead returns the value provided
func ValIfNil[T any](ptr *T, value T) T {
	if ptr == nil {
		return value
	}
	return *ptr
}

// Returns true if the values the pointers a and b point to are equal
// If a and b are both nil returns true
// If either a or b is nii and the other is not returns false
func Equals[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
