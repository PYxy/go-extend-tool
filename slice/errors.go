package slice

import "errors"

var (
	ElementNotExist  = errors.New("元素不存在") //元素不存在
	ZEROSliceLength  = errors.New("切片长度为0")
	SourceSliceIsNil = errors.New("源切片未初始化")
	IndexError       = errors.New("索引值必须大于0")
)
