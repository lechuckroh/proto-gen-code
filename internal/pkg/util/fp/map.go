package fp

func Map[S, T any](values []S, transform func(S) T) []T {
	result := make([]T, 0, len(values))
	for _, value := range values {
		result = append(result, transform(value))
	}
	return result
}
