package slice

// Unique 去重 只保留第一个出现的
func Unique[T comparable](src []T) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	if len(src) == 0 {
		return src, nil
	}
	tmpMap := make(map[T]struct{}, len(src))
	var loop int
	for i := range src {
		i -= loop
		if _, ok := tmpMap[src[i]]; ok {
			src, _ = DeleteByIndex[T](src, i)
			loop += 1
		} else {
			tmpMap[src[i]] = struct{}{}
		}
	}

	return src, nil
}

// UniqueFunc 不可比较模型中的
func UniqueFunc[T any](src []T, equalFunc EqualFunc[T]) ([]T, error) {

	return src, nil
}