package methods

func ForEach[T any](slice []T, callback func(elem T, index int)) {
	for i, v := range slice {
		callback(v, i)
	}
}

func Map[T any](slice []T, callback func(elem T, index int) T) []T {
	mappedSlice := []T{}

	for i, v := range slice {
		mappedSlice = append(mappedSlice, callback(v, i))
	}

	return mappedSlice
}

func Filter[T any](slice []T, callback func(elem T, index int) bool) []T {
	filteredSlice := []T{}

	for i, v := range slice {
		if callback(v, i) {
			filteredSlice = append(filteredSlice, v)
		}
	}

	return filteredSlice
}
