package util

// RemoveDuplicate removes duplicate elements from a slice of integers.
func RemoveDuplicate[S ~[]E, E comparable](arr S, str E) S {
	result := S{}
	for _, v := range arr {
		if v == str {
			continue
		}
		result = append(result, v)
	}
	return result
}

func ContainsDuplicate[S ~[]E, E comparable](arr S) bool {
	res := make(map[E]struct{})
	for _, v := range arr {
		if _, exists := res[v]; exists {
			return true
		}
		res[v] = struct{}{}
	}
	return false
}
