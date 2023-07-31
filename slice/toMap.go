package slice

/*
用于辅助 聚合操作
*/

// ToMap  切片转map
func ToMap[T comparable](src []T) map[T]struct{} {
	var tmpMap map[T]struct{}
	if src == nil || len(src) == 0 {
		tmpMap = make(map[T]struct{}, 0)
		return tmpMap
	}
	tmpMap = make(map[T]struct{}, len(src))
	for _, val := range src {
		if _, ok := tmpMap[val]; !ok {
			tmpMap[val] = struct{}{}
		}
	}
	return tmpMap
}

func ToMapFunc[T any, key comparable](src []T, getKey GetKey[T, key]) map[key]T {
	var tmpMap map[key]T
	if src == nil || len(src) == 0 {
		tmpMap = make(map[key]T, 0)
		return tmpMap
	}
	tmpMap = make(map[key]T, len(src))
	for index, val := range src {
		tmpKey := getKey(val)
		if _, ok := tmpMap[tmpKey]; !ok {
			tmpMap[tmpKey] = src[index]
		}
	}
	return tmpMap
}
