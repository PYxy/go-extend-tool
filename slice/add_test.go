package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBase(t *testing.T) {

	a := make([]int, 0, 10)
	for i := 0; i < 6; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	insertIndex := 3
	insertSlice := []int{98, 97, 96, 95}
	tmp1 := a[0:insertIndex]
	tmp2 := a[insertIndex:]
	fmt.Println(tmp1)

	fmt.Println(tmp2)

	//法1
	a = append(a, insertSlice...)
	copy(a[insertIndex+len(insertSlice):], a[insertIndex:])
	copy(a[insertIndex:insertIndex+len(insertSlice)], insertSlice)
	fmt.Println(a) //[0 1 2 98 97 96 95 3 4 5]

	//法2
	//tmpSLice := make([]int, 0, len(a)+len(insertSlice))
	//
	//tmpSLice = append(tmpSLice, a[0:insertIndex]...)
	//tmpSLice = append(tmpSLice, insertSlice...)
	//tmpSLice = append(tmpSLice, a[insertIndex:]...)
	//fmt.Println(tmpSLice) //[0 1 2 98 97 96 95 3 4 5]
	//copy(a[insertIndex:insertIndex+len(insertSlice)], insertSlice)
	//a[insertIndex] = 99
	//fmt.Println(a)
}

func TestAddIndex(t *testing.T) {

	//initSlice := func() {
	//
	//}
	tests := []struct {
		src         func() []int
		insertIndex int
		insertSlice []int
		wartResurt  []int
		wartError   error
		name        string
	}{
		{
			src: func() []int {
				tmpSLice := make([]int, 0, 10)
				tmpSLice = append(tmpSLice, []int{1, 2, 3, 4}...)
				return tmpSLice
			},
			insertIndex: 0,
			insertSlice: []int{6, 7},
			wartResurt:  []int{6, 7, 1, 2, 3, 4},
			wartError:   nil,
			name:        "在头插入多个数据",
		},
		{
			src: func() []int {
				tmpSLice := make([]int, 0, 10)
				tmpSLice = append(tmpSLice, []int{1, 2, 3, 4}...)
				return tmpSLice
			},
			insertIndex: 4,
			insertSlice: []int{6, 7},
			wartResurt:  []int{1, 2, 3, 4, 6, 7},
			wartError:   nil,
			name:        "在尾部插入多个数据",
		},
		{
			src: func() []int {
				tmpSLice := make([]int, 0, 10)
				tmpSLice = append(tmpSLice, []int{1, 2, 3, 4}...)
				return tmpSLice
			},
			insertIndex: 3,
			insertSlice: []int{6, 7},
			wartResurt:  []int{1, 2, 3, 6, 7, 4},
			wartError:   nil,
			name:        "在中部插入多个数据",
		},
		{
			src: func() []int {
				return nil
			},
			insertIndex: 3,
			insertSlice: []int{6, 7},
			wartResurt:  []int{1, 2, 3, 6, 7, 4},
			wartError:   SourceSliceIsNil,
			name:        "向nil 切片插入数据",
		},
		{
			src: func() []int {
				return []int{}
			},
			insertIndex: 3 - 1,
			insertSlice: []int{6, 7},
			wartResurt:  []int{1, 2, 3, 6, 7, 4},
			wartError:   IndexError,
			name:        "长度为0的切片插入数据",
		},
		{
			src: func() []int {
				return []int{1}
			},
			insertIndex: 0,
			insertSlice: []int{6, 7},
			wartResurt:  []int{6, 7, 1},
			wartError:   nil,
			name:        "需要扩容的切片",
		},
	}
	for _, test := range tests {
		//初始化测试数据
		//initSlice()
		//tmpSLice = make([]int, 0, 10)
		//tmpSLice = append(tmpSLice, []int{1, 2, 3, 4}...)
		t.Run(test.name, func(t *testing.T) {
			result, err := AddIndex[int](test.src(), test.insertIndex, test.insertSlice...)
			//fmt.Println(result)
			if err != nil {
				assert.Equal(t, test.wartError, err)
			} else {
				assert.Equal(t, test.wartResurt, result)
				assert.Equal(t, test.wartError, err)
			}

		})

	}
}
