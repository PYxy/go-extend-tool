package list

import "errors"

var (
	QueueFull       = errors.New("队列已经满了")
	QueueEmpty      = errors.New("队列已空")
	QueueIsNotIni   = errors.New("队列没有初始化")
	IndexOutOfRange = errors.New("索引超标")
)

const (
// IndexOutOfRange = "索引超标: 总长度:%d,索引下标:%d"
)
