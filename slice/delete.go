package slice

import (
	"fmt"
	"sort"
)

// DeleteByIndex 根据index 删除,可一次性删除多个
// src 源切片
// indexSlice 索引切片
func DeleteByIndex[T any](src []T, indexSlice ...int) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	if len(indexSlice) == 0 || len(src) == 0 {
		return src, ZEROSliceLength
	}
	srcLength := len(src)
	//排序之后 直接拿最后一个,如果比
	sort.Ints(indexSlice)
	if indexSlice[0] < 0 {
		return src, IndexError
	}

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

// DeleteByValue 根据val 删除,可一次性删除 count个
// val  删除的值
// count -1  表示全部删除命中的 1 删除命中的第一个
func DeleteByValue[T comparable](src []T, val T, count int) ([]T, error) {
	if len(src) == 0 {
		return src, ElementNotExist
	}
	if count == 0 {
		return src, nil
	}
	matchSlice := IndexAll(src, val)
	if matchSlice == nil {
		return src, ElementNotExist
	}
	//fmt.Println(len(matchSlice), count)
	//比较count 与 matchSlice 的长短
	if count >= len(matchSlice) || count == -1 {
		//倒叙全删除
		for i := len(matchSlice) - 1; i >= 0; i-- {
			if matchSlice[i] == len(src)-1 {
				src = src[:len(src)-1]
			} else {
				src = append(src[0:matchSlice[i]], src[matchSlice[i]+1:]...)
			}

		}
	} else {
		//只删除 count 个
		var loop int
		for i := 0; i < count; i++ {
			src = append(src[0:matchSlice[i-loop]], src[matchSlice[i-loop]+1:]...)
			loop += 1
		}
	}
	return src, nil
}

func DeleteByValueFunc[T any](src []T, val T, count int, equal EqualFunc[T]) ([]T, error) {
	if len(src) == 0 {
		return src, ElementNotExist
	}
	if count == 0 {
		return src, nil
	}
	matchSlice := IndexAllFunc(src, val, equal)
	if matchSlice == nil {
		return src, ElementNotExist
	}

	//比较count 与 matchSlice 的长短
	if count >= len(matchSlice) || count == -1 {
		//倒叙全删除
		for i := len(matchSlice) - 1; i >= 0; i-- {
			if matchSlice[i] == len(src)-1 {
				src = src[:len(src)-1]
			} else {
				src = append(src[0:matchSlice[i]], src[matchSlice[i]+1:]...)
			}

		}
	} else {
		//只删除 count 个
		var loop int
		for i := 0; i < count; i++ {
			src = append(src[0:matchSlice[i-loop]], src[matchSlice[i-loop]+1:]...)
			fmt.Println(src)
			loop += 1
		}
	}

	return src, nil
}
