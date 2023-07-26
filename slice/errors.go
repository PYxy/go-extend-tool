package slice

import "errors"

var (
	ElementNotExist  = errors.New("元素不存在") //切片为空 或者 索引位置越界
	SourceSliceIsNil = errors.New("源切片未初始化")
)
