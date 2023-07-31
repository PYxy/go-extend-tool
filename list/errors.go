package list

import "errors"

var (
	QueueFull  = errors.New("队列已经满了")
	QueueEmpty = errors.New("队列已空")
)
