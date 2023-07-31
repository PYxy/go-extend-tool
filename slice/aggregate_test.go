package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		src       []int
		wantIndex int
		wantVal   int
		wantError error
		name      string
	}{
		{
			src:       []int{1, 2, 3, 5},
			wantIndex: 3,
			wantVal:   5,
			wantError: nil,
			name:      "正常排序1",
		},
		{
			src:       []int{},
			wantIndex: 0,
			wantVal:   0,
			wantError: ZEROSliceLength,
			name:      "正常排序2",
		},
		{
			src:       nil,
			wantIndex: 0,
			wantVal:   0,
			wantError: SourceSliceIsNil,
			name:      "正常排序3",
		},
		{
			src:       []int{5, 2, 3, 5},
			wantIndex: 0,
			wantVal:   5,
			wantError: nil,
			name:      "正常排序4",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, index, err := Max[int](test.src)
			if err != nil {
				assert.Equal(t, test.wantError, err)

			} else {
				assert.Equal(t, test.wantIndex, index)
				assert.Equal(t, test.wantVal, val)
				assert.Equal(t, test.wantError, err)
			}

		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		src       []int
		wantIndex int
		wantVal   int
		wantError error
		name      string
	}{
		{
			src:       []int{1, 2, 3, 5},
			wantIndex: 0,
			wantVal:   1,
			wantError: nil,
			name:      "正常排序1",
		},
		{
			src:       []int{},
			wantIndex: 0,
			wantVal:   0,
			wantError: ZEROSliceLength,
			name:      "正常排序2",
		},
		{
			src:       nil,
			wantIndex: 0,
			wantVal:   0,
			wantError: SourceSliceIsNil,
			name:      "正常排序3",
		},
		{
			src:       []int{5, 2, 2, 5},
			wantIndex: 1,
			wantVal:   2,
			wantError: nil,
			name:      "正常排序4",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, index, err := Min[int](test.src)
			if err != nil {
				assert.Equal(t, test.wantError, err)

			} else {
				assert.Equal(t, test.wantIndex, index)
				assert.Equal(t, test.wantVal, val)
				assert.Equal(t, test.wantError, err)
			}

		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		src       []int
		wantVal   int
		wantError error
		name      string
	}{
		{
			src:       []int{1, 2, 3, 5},
			wantVal:   11,
			wantError: nil,
			name:      "正常1",
		},
		{
			src:       []int{},
			wantVal:   0,
			wantError: nil,
			name:      "正常2",
		},
		{
			src:       nil,
			wantVal:   0,
			wantError: SourceSliceIsNil,
			name:      "异常1",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := Sum[int](test.src)
			if err != nil {
				assert.Equal(t, test.wantError, err)

			} else {
				assert.Equal(t, test.wantVal, val)
				assert.Equal(t, test.wantError, err)
			}

		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		src    []int
		other  []int
		result []int
		name   string
	}{
		{
			result: []int{},
			name:   "src跟other都是nil",
		},
		{
			src:    []int{},
			other:  []int{},
			result: []int{},
			name:   "src跟other都是长度为0的切片",
		},
		{
			src:    []int{1, 2, 3},
			result: []int{},
			name:   "other是nil",
		},
		{
			other:  []int{1, 2, 3},
			result: []int{},
			name:   "src是nil",
		},
		{
			src:    []int{1, 3},
			other:  []int{1, 2, 3},
			result: []int{1, 3},
			name:   "交集正常测试",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := Intersection(test.src, test.other)
			assert.Equal(t, test.result, val)
		})
	}
}

func TestIntersectionFunc(t *testing.T) {
	type student struct {
		name  string
		hobby []string
	}
	tests := []struct {
		src    []student
		other  []student
		result []student
		getKey GetKey[student, string]
		name   string
	}{
		{
			result: []student{},
			name:   "src跟other都是nil",
		},
		{
			src:    []student{},
			result: []student{},
			name:   "other是nil",
		},
		{
			other:  []student{},
			result: []student{},
			name:   "src是nil",
		},
		{
			src:    []student{},
			other:  []student{},
			result: []student{},
			name:   "src跟other都是长度为0的切片",
		},
		{
			src: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小红",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞", "跳", "rap"},
				},
			},
			other: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小红",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞2", "跳", "rap2"},
				},
			},
			result: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				//{
				//	name:  "小明",
				//	hobby: []string{"飞", "跳", "rap"},
				//},
			},
			getKey: func(student student) string {
				return student.hobby[0]
			},
			name: "正常测试",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := IntersectionFunc(test.src, test.other, test.getKey)
			assert.Equal(t, test.result, val)
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		src    []int
		other  []int
		result []int
		name   string
	}{
		{
			result: []int{},
			name:   "src跟other都是nil",
		},
		{
			src:    []int{},
			other:  []int{},
			result: []int{},
			name:   "src跟other都是长度为0的切片",
		},
		{
			src:    []int{1, 2, 3},
			result: []int{1, 2, 3},
			name:   "other是nil",
		},
		{
			src:    []int{1, 3},
			other:  []int{1, 2, 3},
			result: []int{1, 2, 3},
			name:   "并集正常测试",
		},
		{
			src:    []int{7},
			other:  []int{1, 2, 3},
			result: []int{1, 2, 3, 7},
			name:   "并集正常测试2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := Union(test.src, test.other)
			for i := range val {
				assert.NotEqual(t, -1, Index(test.result, val[i]))
			}

		})
	}
}

func TestUnionFunc(t *testing.T) {
	type student struct {
		name  string
		hobby []string
	}
	tests := []struct {
		src       []student
		other     []student
		result    []student
		getKey    GetKey[student, string]
		equalfunc EqualFunc[student]
		name      string
	}{
		{
			result: []student{},
			name:   "src跟other都是nil",
		},
		{
			src:    []student{},
			result: []student{},
			name:   "other是nil",
		},
		{
			other:  []student{},
			result: []student{},
			name:   "src是nil",
		},
		{
			src:    []student{},
			other:  []student{},
			result: []student{},
			name:   "src跟other都是长度为0的切片",
		},
		{
			src: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小红",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞", "跳", "rap"},
				},
			},
			other: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小红",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞2", "跳", "rap2"},
				},
			},
			result: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞2", "跳", "rap2"},
				},
			},
			getKey: func(student student) string {
				return student.hobby[0]
			},
			equalfunc: func(a student, b student) bool {
				return a.hobby[0] == b.hobby[0]
			},
			name: "正常测试",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := UnionFunc(test.src, test.other, test.getKey)
			for i := range val {
				assert.NotEqual(t, -1, IndexFunc(test.result, val[i], test.equalfunc))
			}

			//for i := range val {
			//	assert.NotEqual(t, -1, IndexFunc(test.result, val[i]))
			//}

		})
	}
}

func TestDifferenceSet(t *testing.T) {
	tests := []struct {
		src    []int
		other  []int
		result []int
		name   string
	}{
		{
			result: []int{},
			name:   "src跟other都是nil",
		},
		{
			src:    []int{},
			other:  []int{},
			result: []int{},
			name:   "src跟other都是长度为0的切片",
		},
		{
			src:    []int{1, 2, 3},
			result: []int{1, 2, 3},
			name:   "other是nil",
		},
		{
			other:  []int{1, 2, 3},
			result: []int{},
			name:   "src是nil",
		},
		{
			src:    []int{1, 3},
			other:  []int{1, 2, 3},
			result: []int{1, 3},
			name:   "交集正常测试",
		},
		{
			src:    []int{4, 5, 6},
			other:  []int{1, 2, 3},
			result: []int{4, 5, 6},
			name:   "交集正常测试2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := DifferenceSet(test.src, test.other)
			for i := range val {
				assert.NotEqual(t, -1, Index(test.result, val[i]))
			}

		})
	}
}

func TestDifferenceSetFun(t *testing.T) {
	type student struct {
		name  string
		hobby []string
	}

	tests := []struct {
		src       []student
		other     []student
		result    []student
		getKey    GetKey[student, string]
		equalfunc EqualFunc[student]
		name      string
	}{
		{
			result: []student{},
			name:   "src跟other都是nil",
		},
		{
			src:    []student{},
			result: []student{},
			name:   "other是nil",
		},
		{
			other:  []student{},
			result: []student{},
			name:   "src是nil",
		},
		{
			src:    []student{},
			other:  []student{},
			result: []student{},
			name:   "src跟other都是长度为0的切片",
		},
		{
			src: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小红",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞", "跳", "rap"},
				},
			},
			other: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小红",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞2", "跳", "rap2"},
				},
			},
			result: []student{
				{
					name:  "小白",
					hobby: []string{"唱", "跳", "rap"},
				},
				{
					name:  "小明",
					hobby: []string{"飞", "跳", "rap"},
				},
			},
			getKey: func(student student) string {
				return student.hobby[0]
			},
			equalfunc: func(a student, b student) bool {
				return a.hobby[0] == b.hobby[0]
			},
			name: "正常测试",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := DifferenceSetFunc(test.src, test.other, test.getKey)
			for i := range val {
				assert.NotEqual(t, -1, IndexFunc(test.result, val[i], test.equalfunc))
			}

			//for i := range val {
			//	assert.NotEqual(t, -1, IndexFunc(test.result, val[i]))
			//}

		})
	}
}
