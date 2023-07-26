package slice

// IndexFirst 查询第一个命中的索引值
// -1 表示没找到
func IndexFirst[T comparable](src []T, tag T) int {
	for index, val := range src {
		if val == tag {
			return index
		}
	}
	return -1
}

// IndexAll 查询第一个命中的索引值
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

type equalFunc[T any] func(a T, b T) bool

func IndexFirstFunc[T any](src []T, tag T, equal equalFunc[T]) int {
	for index, val := range src {
		if equal(val, tag) {
			return index
		}
	}
	return -1
}

func IndexFirstAll[T any](src []T, tag T, equal equalFunc[T]) []int {
	var indexSlice []int
	for index, val := range src {
		if equal(val, tag) {
			return append(indexSlice, index)
		}
	}
	return indexSlice
}
