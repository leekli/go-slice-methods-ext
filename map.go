package methods

func Map[T any](slice []T, callback func(elem T, index int) T) []T {
	mappedSlice := []T{}

	for i, v := range slice {
		mappedSlice = append(mappedSlice, callback(v, i))
	}

	return mappedSlice
}
