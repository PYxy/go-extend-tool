package slice

// UnExpansion 切片缩容 (必须大于256)
// hard true  表示 len = cap
// hard false 表示 len <= cap / 3 就缩一半
func UnExpansion[T any](src []T, hard bool) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	//硬缩容
	if hard {

		tmpSlice := make([]T, 0, len(src))
		for i := range src {
			tmpSlice = append(tmpSlice, src[i])
		}

		return tmpSlice, nil
	}

	//缩部分
	if cap(src) < 256 {
		return src, nil
	}
	if len(src) <= cap(src)/3 {
		tmpSlice := make([]T, 0, cap(src)/2)
		for i := range src {
			tmpSlice = append(tmpSlice, src[i])
		}
		return tmpSlice, nil
	}
	return src, nil
}
