package methods

func At[T any](slice []T, index int) (T, bool) {
	// Return a false existence of value, if slice is empty
	if len(slice) <= 0 {
		var value T

		return value, false
	}

	// Return a false existence of value, if index given is larger than length of slice
	if index >= len(slice) {
		var value T

		return value, false
	}

	// Deal with if the index given is a negative number (permitted and legal)
	if index < 0 {
		// Turn it to a positive number
		absIndex := -index

		// Return a false existence of value, if converted index number is larger than length of slice (Needs factor in that -0 will not be given, i.e. when couting back in a slice of len = 5, index -5 could represent index 0 in reality)
		if absIndex > len(slice) {
			var value T

			return value, false
		}

		// Return the found value of given index and existence of true
		return slice[len(slice)-absIndex], true
	}

	return slice[index], true
}

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

func Reduce[T any](slice []T, callback func(accumulator T, currentValue T, index int) T, initialValue T) T {
	var finalValue T

	finalValue = initialValue

	for i, v := range slice {
		finalValue = callback(finalValue, v, i)
	}

	return finalValue
}
