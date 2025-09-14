package slices

// ContainsAll returns true if all elements in b are present in a
func ContainsAll[T comparable](a, b []T) bool {
	checkMap := make(map[T]bool)
	for _, v := range b {
		checkMap[v] = true
	}

	for _, v := range a {
		if _, ok := checkMap[v]; !ok {
			return false
		}
	}

	return true
}
