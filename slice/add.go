package slice

// AddIndex 允许越界(超标就直接放在最后)
// src 目标切片
// val  插入值 可多个
// index  索引值
func AddIndex[T any](src []T, index int, valSLice ...T) ([]T, error) {
	if src == nil {
		return src, SourceSliceIsNil
	}
	//索引值不能是负数 而且 必须小于目标切片的cap
	if index < 0 || index > cap(src) {
		return src, IndexError
	}
	if cap(src) >= len(src)+len(valSLice) {
		//先补齐不够的空间
		src = append(src, valSLice...)
		//后移
		copy(src[index+len(valSLice):], src[index:])
		//再填补
		copy(src[index:index+len(valSLice)], valSLice)
		return src, nil
	} else {
		tmpSLice := make([]T, 0, len(src)+len(valSLice))
		tmpSLice = append(tmpSLice, src[0:index]...)
		tmpSLice = append(tmpSLice, valSLice...)
		tmpSLice = append(tmpSLice, src[index:]...)
		return tmpSLice, nil
	}
}
