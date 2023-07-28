package slice

// Max 返回切片的最大值
func Max[T Number](src []T) (max T, index int, err error) {
	if src == nil {
		err = SourceSliceIsNil
		return
	}
	if len(src) == 0 {
		err = ZEROSliceLength
		return
	}
	max = src[0]
	index = 0
	for i := 1; i <= len(src)-1; i++ {
		if src[i] > max {
			max = src[i]
			index = i
		}
	}
	return

}

// Min 返回切片的最小值
func Min[T Number](src []T) (min T, index int, err error) {
	if src == nil {
		err = SourceSliceIsNil
		return
	}
	if len(src) == 0 {
		err = ZEROSliceLength
		return
	}
	min = src[0]
	index = 0
	for i := 1; i <= len(src)-1; i++ {
		if src[i] < min {
			min = src[i]
			index = i
		}
	}
	return
}

// Intersection 交集
func Intersection[T any](src []T, other []T) []T {
	return nil
}

// Union 并集
func Union[T any](src []T, other []T) []T {
	return nil
}

// DifferenceSet 差集
func DifferenceSet[T any](src []T, other []T) []T {
	return nil
}
