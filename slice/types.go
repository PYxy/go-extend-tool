package slice

type EqualFunc[T any] func(a T, b T) bool

type Number interface {
	~int64 | ~int32 | ~int8 | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~int | ~float64 | ~float32
}
