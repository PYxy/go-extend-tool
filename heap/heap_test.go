package heap

import (
	"fmt"
	"testing"
)

// 大顶堆
func TestNewHeapBig(t *testing.T) {
	//fmt.Println(int(math.Floor(float64(2-1) / 2)))
	//return
	mq := NewHeap(10, func(t int, t2 int) bool {
		return t > t2
	})
	sourceIndex := []int{20, 30, 15, 49, 10, 62}
	for _, val := range sourceIndex {
		mq.Enqueue(val)
	}
	fmt.Println(mq.data)
	for {
		tmpData, err := mq.DeQueue()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("出队元素:", tmpData)
	}
}

// 小顶堆
func TestNewHeapSmall(t *testing.T) {
	//fmt.Println(int(math.Floor(float64(2-1) / 2)))
	//return
	mq := NewHeap(10, func(t int, t2 int) bool {
		return t < t2
	})
	sourceIndex := []int{20, 30, 15, 49, 10, 62}
	for _, val := range sourceIndex {
		mq.Enqueue(val)
	}
	fmt.Println(mq.data)
	for {
		tmpData, err := mq.DeQueue()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("出队元素:", tmpData)
	}
}
