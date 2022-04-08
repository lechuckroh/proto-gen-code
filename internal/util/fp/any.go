package fp

func Any[T any](values []T, predicate func(T) bool) bool {
	for _, value := range values {
		if predicate(value) {
			return true
		}
	}
	return false
}
