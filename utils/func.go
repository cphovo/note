package utils

// This function takes in a slice of any type and a condition function as parameters and returns a slice of the same type
func Filter[T any](slice []T, condition func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}
