package slice

import (
	"sort"
)

// DeleteByIndex 根据index 删除,可一次性删除多个
func DeleteByIndex[T any](src []T, indexSlice ...int) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	if len(indexSlice) == 0 || len(src) == 0 {
		return src, ElementNotExist
	}
	srcLength := len(src)
	//排序之后 直接拿最后一个
	sort.Ints(indexSlice)
	if indexSlice[len(indexSlice)-1] > srcLength {
		return src, ElementNotExist
	}
	var Loop int
	// 执行删除操作这个会影响性能
	for _, v := range indexSlice {
		v -= Loop
		if v == srcLength-1 {
			src = src[0:v]
			Loop += 1
			continue
		}
		src = append(src[0:v], src[v+1:]...) //源数组一直在变必须要记下来
		Loop += 1
	}
	return src, nil
}

// DeleteByValueAsc 根据vla 删除,可一次性删除 count个
// count -1  表示全部删除命中的
// count 1 删除第一个
func DeleteByValueAsc[T comparable](src []T, val T, count int) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	if len(src) == 0 {
		return src, ElementNotExist
	}
	var Loop int
	srcLength := len(src)
	// 执行删除操作这个会影响性能
	for index, v := range src {
		if v == val {
			index -= Loop
			if index == srcLength-1 {

				src = src[0:index]
				Loop += 1
				continue
			}
			src = append(src[0:index], src[index+1:]...) //源数组一直在变必须要记下来
			Loop += 1
			srcLength -= 1
			if count > 0 && Loop >= count { //
				break
			}
		}
	}
	return src, nil
}
