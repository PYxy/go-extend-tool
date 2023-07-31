package slice

// Unique 去重 只保留第一个出现的(如果是可比较模型 保留那一个都可以)
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

	return unExpansion(src)
}

// UniqueFuncFirst 不可比较类型中的(只保留第一个出现的)
func UniqueFuncFirst[T any, key comparable](src []T, getKey GetKey[T, key]) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	if len(src) == 0 {
		return src, nil
	}
	tmpMap := make(map[key]struct{}, len(src))
	var loop int
	for i := range src {
		i -= loop
		mapKey := getKey(src[i])
		if _, ok := tmpMap[mapKey]; ok {
			src, _ = DeleteByIndex[T](src, i)
			loop += 1
		} else {
			tmpMap[mapKey] = struct{}{}
		}
	}

	return unExpansion(src)
}

// UniqueFuncLast 不可比较类型中的(只保留最后一个出现的)
func UniqueFuncLast[T any, key comparable](src []T, getKey GetKey[T, key]) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	if len(src) == 0 {
		return src, nil
	}
	tmpMap := make(map[key]int, len(src))
	var loop int
	for i := range src {
		i -= loop
		mapKey := getKey(src[i])
		if index, ok := tmpMap[mapKey]; ok {
			tmpMap[mapKey] = i
			//删除的index  要根据实际情况进行移动
			src, _ = DeleteByIndex[T](src, index-loop)

			loop += 1
		} else {
			tmpMap[mapKey] = i
		}
	}

	return unExpansion(src)
}
