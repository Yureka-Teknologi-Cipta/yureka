package helpers

// InArrayString Checks if a value exists in an array
func InArrayString(v string, haystack []string) bool {
	for _, val := range haystack {
		if v == val {
			return true
		}
	}
	return false
}

// InArrayInt Checks if a value exists in an array
func InArrayInt(v int, haystack []int) bool {
	for _, val := range haystack {
		if v == val {
			return true
		}
	}
	return false
}

// InArrayInt64 Checks if a value exists in an array
func InArrayInt64(v int64, haystack []int64) bool {
	for _, val := range haystack {
		if v == val {
			return true
		}
	}
	return false
}
