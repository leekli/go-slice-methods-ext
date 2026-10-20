package methods

func (slice Slice[T]) forEach(callback func(elem T)) {
	for i := 0; i < len(slice); i++ {
		callback(slice[i])
	}
}
