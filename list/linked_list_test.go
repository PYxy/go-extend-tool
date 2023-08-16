package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListBase(t *testing.T) {
	a := NewLinkedList[int]()
	a.Append([]int{1, 2, 7, 8, 6}...)

	fmt.Println("长度:", a.length)
	head := a.head.next
	for i := 0; i < a.length; i++ {
		fmt.Println(head.data)
		head = head.next
	}
	fmt.Println("反序")
	tail := a.tail.prev
	for i := 0; i < a.length; i++ {
		fmt.Println(tail.data)
		tail = tail.prev
	}
	fmt.Println("按索引查找")
	varl, err := a.Get(0)
	fmt.Println(varl, err)
	//fmt.Println(a.head.data)
	//fmt.Println(a.tail.data)
	//fmt.Println(a.tail.prev.data)

	a.Add(3, 99)
	a.Set(3, 101)
	fmt.Println(a.Delete(3))
	fmt.Println("反序")
	tail = a.tail.prev
	for i := 0; i < a.length; i++ {
		fmt.Println(tail.data)
		tail = tail.prev
	}
	fmt.Println("排序")
	head = a.head.next
	for i := 0; i < a.length; i++ {
		fmt.Println(head.data)
		head = head.next
	}
}

func TestLinkedList_Get(t *testing.T) {

	linkList, err := NewLinkedBySlice[int]([]int{1, 2, 3, 4, 5, 6, 7})
	if err != nil {
		panic("双向链表初始化失败")
	}
	//initSlice := func() {
	//
	//}
	tests := []struct {
		src         *LinkedList[int]
		searchIndex int
		wartResurt  int
		wartError   error
		name        string
	}{
		{
			src:         linkList,
			searchIndex: 0,
			wartResurt:  1,
			wartError:   nil,
			name:        "获取第一个元素",
		},
		{
			src:         linkList,
			searchIndex: linkList.length - 1,
			wartResurt:  7,
			wartError:   nil,
			name:        "获取最后一个元素",
		},
		{
			src:         linkList,
			searchIndex: linkList.length - 2,
			wartResurt:  6,
			wartError:   nil,
			name:        "获取中间元素",
		},
		{
			src:         linkList,
			searchIndex: linkList.length,
			wartResurt:  0,
			wartError:   IndexOutOfRange,
			name:        "索引超标",
		},
	}
	for _, test := range tests {
		//初始化测试数据
		//initSlice()
		//tmpSLice = make([]int, 0, 10)
		//tmpSLice = append(tmpSLice, []int{1, 2, 3, 4}...)
		t.Run(test.name, func(t *testing.T) {
			result, err := linkList.Get(test.searchIndex)
			assert.Equal(t, test.wartError, err)
			assert.Equal(t, test.wartResurt, result)

		})

	}

}

func TestLinkedList_Append(t *testing.T) {
	tests := []struct {
		src         *LinkedList[int]
		AppendSlice []int
		wartSlice   []int
		wartError   error
		name        string
	}{
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			AppendSlice: []int{1},
			wartSlice:   []int{1, 2, 1},
			wartError:   nil,
			name:        "添加一个元素",
		},
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			AppendSlice: []int{},
			wartSlice:   []int{1, 2},
			wartError:   nil,
			name:        "添加0个元素",
		},
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			AppendSlice: []int{7, 8},
			wartSlice:   []int{1, 2, 7, 8},
			wartError:   nil,
			name:        "添加多个元素",
		},
		{
			src: func() *LinkedList[int] {
				return &LinkedList[int]{}
			}(),
			AppendSlice: []int{1},
			wartSlice:   []int{},
			wartError:   QueueIsNotIni,
			name:        "空链表添加元素",
		},
	}
	for _, test := range tests {
		//初始化测试数据
		//initSlice()
		//tmpSLice = make([]int, 0, 10)
		//tmpSLice = append(tmpSLice, []int{1, 2, 3, 4}...)
		t.Run(test.name, func(t *testing.T) {
			err := test.src.Append(test.AppendSlice...)
			assert.Equal(t, test.wartError, err)
			if err == nil {
				assert.Equal(t, test.src.AsSliceAsc(), test.wartSlice)

			}

		})

	}

}

func TestLinkedList_Add(t *testing.T) {
	tests := []struct {
		src         *LinkedList[int]
		AppendIndex int
		AppendValue int
		wartSlice   []int
		wartError   error
		name        string
	}{
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			AppendIndex: 1,
			AppendValue: 10,
			wartSlice:   []int{1, 10, 2},
			wartError:   nil,
			name:        "中间添加一个元素",
		},
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			AppendIndex: 0,
			AppendValue: 10,
			wartSlice:   []int{10, 1, 2},
			wartError:   nil,
			name:        "头部添加一个元素",
		},
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			AppendIndex: 10,
			AppendValue: 10,
			wartSlice:   []int{},
			wartError:   IndexOutOfRange,
			name:        "索引超标",
		},
		{
			src: func() *LinkedList[int] {
				return &LinkedList[int]{}
			}(),
			AppendIndex: 0,
			wartSlice:   []int{},
			wartError:   IndexOutOfRange,
			name:        "空链表添加元素",
		},
	}
	for _, test := range tests {
		//初始化测试数据
		//initSlice()
		//tmpSLice = make([]int, 0, 10)
		//tmpSLice = append(tmpSLice, []int{1, 2, 3, 4}...)
		t.Run(test.name, func(t *testing.T) {
			err := test.src.Add(test.AppendIndex, test.AppendValue)
			assert.Equal(t, test.wartError, err)
			if err == nil {
				assert.Equal(t, test.src.AsSliceAsc(), test.wartSlice)

			}

		})

	}

}

func TestLinkedList_Delete(t *testing.T) {
	tests := []struct {
		src        *LinkedList[int]
		deletendex int
		wartVal    int
		wartSlice  []int
		wartError  error
		name       string
	}{
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			deletendex: 1,
			wartVal:    2,
			wartSlice:  []int{1},
			wartError:  nil,
			name:       "删除最后一个元素",
		},
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			deletendex: 0,
			wartVal:    1,
			wartSlice:  []int{2},
			wartError:  nil,
			name:       "删除第一个元素",
		},
		{
			src: func() *LinkedList[int] {
				linkList, err := NewLinkedBySlice[int]([]int{1, 2})
				if err != nil {
					panic("双向链表初始化失败")
				}
				return linkList
			}(),
			deletendex: 3,
			wartVal:    0,
			wartSlice:  []int{},
			wartError:  IndexOutOfRange,
			name:       "索引超标",
		},
		{
			src: func() *LinkedList[int] {
				return &LinkedList[int]{}
			}(),
			deletendex: 0,
			wartSlice:  []int{},
			wartError:  IndexOutOfRange,
			name:       "空链表添加元素",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := test.src.Delete(test.deletendex)
			assert.Equal(t, test.wartError, err)
			if err == nil {

				assert.Equal(t, test.wartVal, val)
				assert.Equal(t, test.src.AsSliceAsc(), test.wartSlice)

			}

		})

	}

}
