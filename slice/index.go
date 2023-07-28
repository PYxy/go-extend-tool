package slice

// Index 查询第一个命中的索引值
// -1 表示没找到
func Index[T comparable](src []T, tag T) int {
	for index, val := range src {
		if val == tag {
			return index
		}
	}
	return -1
}

// IndexAll 查询命中的所有索引值
// nil 表示不存在
func IndexAll[T comparable](src []T, tag T) []int {
	var indexSlice []int
	for index, val := range src {
		if val == tag {
			indexSlice = append(indexSlice, index)
		}
	}
	return indexSlice
}

func IndexFunc[T any](src []T, tag T, equal EqualFunc[T]) int {
	for index, val := range src {
		if equal(val, tag) {
			return index
		}
	}
	return -1
}

func IndexAllFunc[T any](src []T, tag T, equal EqualFunc[T]) []int {
	var indexSlice []int
	for index, val := range src {
		if equal(val, tag) {
			indexSlice = append(indexSlice, index)
		}
	}
	return indexSlice
}
