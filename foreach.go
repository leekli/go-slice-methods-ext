package methods

func ForEach[T any](slice []T, callback func(elem T, index int)) {
	for i, v := range slice {
		callback(v, i)
	}
}
