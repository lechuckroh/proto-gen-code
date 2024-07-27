package fp

func Filter[T any](values []T, predicate func(T) bool) (result []T) {
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}
