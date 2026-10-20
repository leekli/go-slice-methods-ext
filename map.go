package methods

type AllowedSliceTypes interface {
	~string |
		~int | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~bool
}

type Slice[T AllowedSliceTypes] []T

func (slice Slice[T]) forEach(callback func(elem T)) {
	for i := 0; i < len(slice); i++ {
		callback(slice[i])
	}
}
