package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

func TestBase(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	a = a[:2]

	fmt.Println(a, " ", len(a), "  ", cap(a)) //[1 2]   2    3
	b = append(b[:1], b[2:]...)
	fmt.Println(b, " ", len(b), "  ", cap(b)) //[1 3]   2    3

	c := []int{4, 2, 10}
	sort.Ints(c)
	fmt.Println(c)
	for i := range c {
		fmt.Println(i)
	}

	d := make([]int, 0, 10)
	d = append(d, 1)
	d = append(d, 5)
	fmt.Println(d[0:])
}

func TestBase1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//for i, v := range a {
	//	fmt.Println(i, v)
	//	a = a[i:]
	//}

	//倒序输出
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i])
	}

}

func TestDeleteByIndex(t *testing.T) {
	tests := []struct {
		src        []int
		indexSlice []int
		wantSlice  []int
		wartError  error
		name       string
	}{
		{
			src:        []int{1, 2, 3, 5},
			indexSlice: []int{0, 1},
			wantSlice:  []int{3, 5},
			wartError:  nil,
			name:       "first one",
		},
		{
			src:        []int{},
			indexSlice: []int{0, 1},
			wantSlice:  []int{},
			wartError:  ZEROSliceLength,
			name:       "sorce slice length is 0",
		},
		{

			indexSlice: []int{0, 1},
			wantSlice:  nil,
			wartError:  SourceSliceIsNil,
			name:       "sorce slice is nil",
		},
		{
			src:        []int{1, 2, 3, 5},
			indexSlice: []int{0, 1, 5},
			wantSlice:  []int{1, 2, 3, 5},
			wartError:  ElementNotExist,
			name:       "indexSlice out of range",
		},
		{
			src:        []int{1, 2, 3, 5},
			indexSlice: []int{0, 3},
			wantSlice:  []int{2, 3},
			wartError:  nil,
			name:       "last one",
		},
		{
			src:        []int{1, 2, 3, 5},
			indexSlice: []int{2},
			wantSlice:  []int{1, 2, 5},
			wartError:  nil,
			name:       "middle one",
		},
		{
			src:        []int{1, 2, 3, 5},
			indexSlice: []int{1, 3},
			wantSlice:  []int{1, 3},
			wartError:  nil,
			name:       "middle two",
		},
		{
			src:        []int{1, 2, 3, 5},
			indexSlice: []int{0, 1},
			wantSlice:  []int{3, 5},
			wartError:  nil,
			name:       "middle two2",
		},
		{
			src:        []int{1, 2, 3, 5},
			indexSlice: []int{3},
			wantSlice:  []int{1, 2, 3},
			wartError:  nil,
			name:       "delete the last",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := DeleteByIndex[int](test.src, test.indexSlice...)

			assert.Equal(t, test.wantSlice, result)
			assert.Equal(t, test.wartError, err)
		})
	}
}

func TestDeleteByValue(t *testing.T) {

	tests := []struct {
		src       []int
		tagVal    int
		count     int
		wantSlice []int
		wartError error
		name      string
	}{
		{
			src:       []int{1, 2, 1, 5},
			tagVal:    1,
			count:     1,
			wantSlice: []int{2, 1, 5},
			wartError: nil,
			name:      "删除第一个1",
		},
		{
			src:       []int{},
			tagVal:    1,
			count:     -1,
			wantSlice: []int{},
			wartError: ElementNotExist,
			name:      "查询失败",
		},
		{
			src:       []int{1, 2, 1, 5},
			tagVal:    1,
			count:     -1,
			wantSlice: []int{2, 5},
			wartError: nil,
			name:      "删除全部1",
		},
		{
			src:       []int{5, 2, 1, 5},
			tagVal:    5,
			count:     3,
			wantSlice: []int{2, 1},
			wartError: nil,
			name:      "删除3个5",
		},
		{
			src:       []int{5, 2, 1, 5},
			tagVal:    2,
			count:     -1,
			wantSlice: []int{5, 1, 5},
			wartError: nil,
			name:      "删除全部2",
		},
		{
			src:       []int{1, 2, 1, 5},
			tagVal:    5,
			count:     10,
			wantSlice: []int{1, 2, 1},
			wartError: nil,
			name:      "删除222",
		},
		{
			src:       []int{1, 2, 1, 5},
			tagVal:    5,
			count:     0,
			wantSlice: []int{1, 2, 1, 5},
			wartError: nil,
			name:      "不删除元素",
		},
		{
			src:       []int{1, 2, 2, 5},
			tagVal:    2,
			count:     2,
			wantSlice: []int{1, 5},
			wartError: nil,
			name:      "连续删除",
		},
		{
			src:       []int{1, 2, 2, 2, 5},
			tagVal:    2,
			count:     2,
			wantSlice: []int{1, 2, 5},
			wartError: nil,
			name:      "有连续多个2(连续删除2个2)",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := DeleteByValue[int](test.src, test.tagVal, test.count)
			fmt.Println(len(result), cap(result))
			assert.Equal(t, test.wantSlice, result)
			assert.Equal(t, test.wartError, err)
		})
	}
}

func TestDeleteByValueFunc(t *testing.T) {
	type student struct {
		name  string
		hobby []string
		mark  map[string]int
	}

	tests := []struct {
		src       []student
		targetVal student
		equal     func(student, student) bool
		result    []student
		err       error
		count     int
		name      string
	}{
		{
			src: []student{student{
				name:  "小宏",
				hobby: []string{"唱", "跳", "rapper"},
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			},
				student{
					name:  "小白",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
			},
			targetVal: student{
				name:  "小宏",
				hobby: []string{"唱1", "跳1", "rapper1"},
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			},
			count: -1,
			equal: func(i student, i2 student) bool {
				//return i.name == i2.name
				return i.mark["语文"] == i2.mark["语文"]
			},
			result: []student{},
			name:   "全删除",
		},

		{
			src: []student{student{
				name:  "小宏",
				hobby: []string{"唱1", "跳", "rapper"},
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			},
				student{
					name:  "小白",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
			},

			targetVal: student{
				name:  "小宏",
				hobby: []string{"唱1", "跳1", "rapper1"},
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			},
			equal: func(i student, i2 student) bool {
				//return i.name == i2.name
				return i.hobby[0] == i2.hobby[0]
			},
			result: []student{
				student{
					name:  "小白",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
			},
			err:   nil,
			count: 1,
			name:  "only two",
		},
		{
			src: []student{student{
				name:  "小宏",
				hobby: []string{"唱", "跳", "rapper"},
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			},
				student{
					name:  "小白",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
			},
			targetVal: student{
				name:  "小宏",
				hobby: []string{"唱1", "跳1", "rapper1"},
				mark: map[string]int{
					"数学": 15,
					"语文": 17,
				},
			},
			count: -1,
			equal: func(i student, i2 student) bool {
				//return i.name == i2.name
				return i.mark["语文"] == i2.mark["语文"]
			},
			result: nil,
			err:    ElementNotExist,
			name:   "没有匹配对象",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := DeleteByValueFunc[student](test.src, test.targetVal, test.count, test.equal)
			//fmt.Println(reflect.DeepEqual(test.result,result))

			if err != nil {
				fmt.Println(test.result, "肉眼对比:", result)
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, true, reflect.DeepEqual(test.result, result))
				assert.Equal(t, test.err, err)
			}

		})
	}
}

// 缩容测试
func TestDeleteAndUnexpansion(t *testing.T) {
	unExpansionSlice := make([]int, 0, 400)
	for i := 0; i <= 100; i++ {
		unExpansionSlice = append(unExpansionSlice, i)
	}
	tests := []struct {
		src          []int
		tagVal       int
		count        int
		wantSliceCap int
		wantSliceLen int
		wartError    error
		name         string
	}{
		{
			src:          unExpansionSlice,
			tagVal:       2,
			count:        -1,
			wantSliceCap: 200,
			wantSliceLen: 100,
			name:         "只删除一个,并进行了缩容",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := DeleteByValue[int](test.src, test.tagVal, test.count)
			assert.Equal(t, test.wantSliceCap, cap(result))
			assert.Equal(t, test.wantSliceLen, len(result))
			assert.Equal(t, test.wartError, err)
		})
	}
}
