package array

func Map[T, R any](array []T, mapFunction func(T) R) []R {
	result := make([]R, len(array))
	for i := range array {
		result[i] = mapFunction(array[i])
	}
	return result
}
