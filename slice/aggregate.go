package slice

import "reflect"

/*
在使用 float32 或者 float64 的时候要小心精度问题
*/

// Max 返回切片的最大值
func Max[T RealNumber](src []T) (max T, index int, err error) {
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
func Min[T RealNumber](src []T) (min T, index int, err error) {
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

func Sum[T Number](src []T) (sum T, err error) {
	if src == nil {
		err = SourceSliceIsNil
		return
	}
	for _, n := range src {
		sum += n
	}
	return sum, nil

}

// Intersection 可比较类型(交集)
func Intersection[T comparable](src []T, other []T) (res []T) {
	//先去重
	srcMap := ToMap(src)
	otherMap := ToMap(other)

	//取小的map
	if len(src) > len(otherMap) {
		res = make([]T, 0, len(otherMap))
	} else {
		res = make([]T, 0, len(src))
	}

	for val := range srcMap {
		if _, ok := otherMap[val]; ok {
			res = append(res, val)
		}
	}

	return res
}

// IntersectionFunc 交集
func IntersectionFunc[T any, key comparable](src []T, other []T, getKey GetKey[T, key]) (res []T) {
	//先去重
	srcMap := ToMapFunc(src, getKey)
	otherMap := ToMapFunc(other, getKey)

	//取小的map
	if len(src) > len(otherMap) {
		res = make([]T, 0, len(otherMap))
	} else {
		res = make([]T, 0, len(src))
	}

	for val := range srcMap {
		if _, ok := otherMap[val]; ok {
			if reflect.DeepEqual(srcMap[val], otherMap[val]) {
				res = append(res, srcMap[val])
			}

		}
	}

	return res
}

// Union 可比较类型(并集)
func Union[T comparable](src []T, other []T) (res []T) {
	//先去重
	srcMap := ToMap(src)
	otherMap := ToMap(other)
	res = make([]T, 0, len(otherMap)+len(srcMap))
	for val := range srcMap {
		if _, ok := otherMap[val]; ok {
			//删除重复的部分
			delete(otherMap, val)
		}
		//在不在都加
		res = append(res, val)
	}

	for val := range otherMap {
		res = append(res, val)
	}

	return res
}

// UnionFunc 并集
func UnionFunc[T any, key comparable](src []T, other []T, getKey GetKey[T, key]) (res []T) {
	//先去重
	srcMap := ToMapFunc(src, getKey)
	otherMap := ToMapFunc(other, getKey)
	res = make([]T, 0, len(otherMap)+len(srcMap))
	for val := range srcMap {
		if _, ok := otherMap[val]; ok {
			//删除重复的部分
			if reflect.DeepEqual(srcMap[val], otherMap[val]) {
				delete(otherMap, val)
			}
		}
		//在不在都加
		res = append(res, srcMap[val])
	}

	for val := range otherMap {
		res = append(res, otherMap[val])
	}

	return res
}

// DifferenceSet 可比较类型(差集)
func DifferenceSet[T comparable](src []T, other []T) (res []T) {
	//先去重
	srcMap := ToMap(src)
	otherMap := ToMap(other)
	//取大的map
	if len(src) > len(otherMap) {
		res = make([]T, 0, len(src))
	} else {
		res = make([]T, 0, len(otherMap))
	}
	for val := range srcMap {
		if _, ok := otherMap[val]; ok {
			//跳过重复部分
			continue
		}
		//在不在都加
		res = append(res, val)
	}
	return res
}

// DifferenceSetFunc 差集
func DifferenceSetFunc[T any, key comparable](src []T, other []T, getKey GetKey[T, key]) (res []T) {
	//先去重
	srcMap := ToMapFunc(src, getKey)
	otherMap := ToMapFunc(other, getKey)
	res = make([]T, 0, len(otherMap)+len(srcMap))
	for val := range srcMap {
		if _, ok := otherMap[val]; ok {
			//跳过重复部分
			if reflect.DeepEqual(srcMap[val], otherMap[val]) {
				continue
			}
		}
		//在不在都加
		res = append(res, srcMap[val])
	}

	return res
}
