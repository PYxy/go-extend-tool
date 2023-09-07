package heap

import (
	"errors"
	"math"
)

type CompareFunc[T any] func(T, T) bool

type MyQueue[T any] struct {
	data        []T
	compareFunc CompareFunc[T]
}

func NewHeap[T any](cap int, compareFunc CompareFunc[T]) *MyQueue[T] {
	return &MyQueue[T]{
		data:        make([]T, 0, cap),
		compareFunc: compareFunc,
	}
}

func (m *MyQueue[T]) Enqueue(obj T) {
	m.data = append(m.data, obj)
	length := len(m.data)
	node := length - 1
	if node > 0 {
		for {
			//先找父 (这里的操作都是针对 索引来操作的)
			parent := int(math.Floor(float64(node-1) / 2))
			//只需要跟父比较
			if parent >= 0 && m.compareFunc(m.data[node], m.data[parent]) {
				m.data[node], m.data[parent] = m.data[parent], m.data[node]
				node = parent
			} else {
				break
			}
		}

	}
	return
}

func (m *MyQueue[T]) DeQueue() (T, error) {
	var t T
	if len(m.data) <= 0 {
		return t, errors.New("数组为空")
	}

	t = m.data[0]

	m.inHeap()
	return t, nil
}

func (m *MyQueue[T]) inHeap() {
	//删除头元素之后 先把最后一个元素塞到头位置
	//按需缩容
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[0 : len(m.data)-1]
	length := len(m.data)
	start := 0
	targetIndex := start
	//这里是以父亲节点往下比较
	/*
		parent(i) = floor((i - 1)/2)
		left(i)   = 2i + 1
		right(i)  = 2i + 2 // right(i) # 就是简单的 left(i) + 1。左右节点总是处于相邻的位置
	*/
	for {
		//先找左节点
		if left := 2*start + 1; left <= length-1 && m.compareFunc(m.data[left], m.data[targetIndex]) {

			targetIndex = left
		}
		//在比较右节点
		if right := 2*start + 2; right <= length-1 && m.compareFunc(m.data[right], m.data[targetIndex]) {

			targetIndex = right
		}
		if targetIndex == start {
			break
		}
		//选出最大的 或者最小的进行替换
		m.data[targetIndex], m.data[start] = m.data[start], m.data[targetIndex]
		start = targetIndex
	}
}

func main() {

}
