package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexFirst(t *testing.T) {
	tests := []struct {
		src       []int
		targetVal int
		wantIndex int
		name      string
	}{
		{
			src:       []int{1, 2, 3, 5},
			targetVal: 1,
			wantIndex: 0,
			name:      "find the first val",
		},
		{
			src:       []int{},
			targetVal: 1,
			wantIndex: -1,
			name:      "the length of src is 0",
		},
		{
			targetVal: 1,
			wantIndex: -1,
			name:      "the src nil",
		},
		{
			src:       []int{1, 4, 6},
			targetVal: 7,
			wantIndex: -1,
			name:      "index out of range",
		},
		{
			src:       []int{1, 3, 4, 2, 0},
			targetVal: 0,
			wantIndex: 4,
			name:      "find the src last val",
		},
		{
			src:       []int{1, 3, 3, 2, 0},
			targetVal: 3,
			wantIndex: 1,
			name:      "find the val in the sce has double val",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.wantIndex, Index[int](test.src, test.targetVal))
		})
	}
}

func TestIndexAll(t *testing.T) {
	tests := []struct {
		src       []int
		targetVal int
		resultSrc []int
		name      string
	}{
		{
			src:       []int{1, 1, 3, 5},
			targetVal: 1,
			resultSrc: []int{0, 1},
			name:      "first one",
		},
		{
			src:       []int{},
			targetVal: 1,
			resultSrc: nil,
			name:      "the length of src is 0",
		},
		{
			targetVal: 1,
			resultSrc: nil,
			name:      "src is nil",
		},
		{
			src:       []int{1, 4, 6},
			targetVal: 7,
			resultSrc: nil,
			name:      "src has not targetVal",
		},
		{
			src:       []int{1, 3, 4, 2, 0},
			targetVal: 0,
			resultSrc: []int{4},
			name:      "find the last val index",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.resultSrc, IndexAll[int](test.src, test.targetVal))
		})
	}
}

//不可比较类型  map  slice func  只能跟nil 对比

func TestIndexFunc(t *testing.T) {
	tests := []struct {
		src         []int
		targetVal   int
		equal       func(int, int) bool
		targetIndex int
		name        string
	}{
		{
			src:       []int{1, 2, 3, 5},
			targetVal: 1,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			targetIndex: 0,
			name:        "first one",
		},
		{
			src:       []int{},
			targetVal: 1,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			targetIndex: -1,
			name:        "the length of src is 0",
		},
		{
			targetVal: 1,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			targetIndex: -1,
			name:        "src nil",
		},
		{
			src:       []int{1, 4, 6},
			targetVal: 7,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			targetIndex: -1,
			name:        "dst not exist",
		},
		{
			src:       []int{1, 3, 4, 2, 0},
			targetVal: 0,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			targetIndex: 4,
			name:        "last one",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.targetIndex, IndexFunc[int](test.src, test.targetVal, test.equal))
		})
	}
}

func TestIndexFirstFunc2(t *testing.T) {
	type student struct {
		name  string
		hobby []string
		mark  map[string]int
	}

	tests := []struct {
		src         []student
		targetVal   student
		equal       func(student, student) bool
		targetIndex int
		name        string
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
			equal: func(i student, i2 student) bool {
				//return i.name == i2.name
				return i.mark["语文"] == i2.mark["语文"]
			},
			targetIndex: 0,
			name:        "only one",
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
					"语文": 16,
				},
			},
			equal: func(i student, i2 student) bool {
				//return i.name == i2.name
				return i.hobby[0] == i2.hobby[0]
			},
			targetIndex: -1,
			name:        "only two",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.targetIndex, IndexFunc[student](test.src, test.targetVal, test.equal))
		})
	}
}

func TestIndexAllFunc(t *testing.T) {
	type student struct {
		name  string
		hobby []string
		mark  map[string]int
	}

	tests := []struct {
		src         []student
		targetVal   student
		equal       func(student, student) bool
		targetIndex []int
		name        string
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
			equal: func(i student, i2 student) bool {
				//return i.name == i2.name
				return i.mark["语文"] == i2.mark["语文"]
			},
			targetIndex: []int{0, 1},
			name:        "has two",
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
					"语文": 16,
				},
			},
			equal: func(i student, i2 student) bool {
				//return i.name == i2.name
				return i.hobby[0] == i2.hobby[0]
			},
			targetIndex: nil,
			name:        "only two",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.targetIndex, IndexAllFunc[student](test.src, test.targetVal, test.equal))
		})
	}
}
