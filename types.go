package methods

type AllowedSliceTypes interface {
	~string |
		~int | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~bool
}

type Slice[T AllowedSliceTypes] []T
