package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexFirst(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want int
		name string
	}{
		{
			src:  []int{1, 2, 3, 5},
			dst:  1,
			want: 0,
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: -1,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
			name: "last one",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexFirst[int](test.src, test.dst))
		})
	}
}

func TestIndexAll(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want []int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: []int{0, 1},
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: nil,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: nil,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: nil,
			name: "dst not exist",
		},
		{
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: []int{4},
			name: "last one",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexAll[int](test.src, test.dst))
		})
	}
}

//不可比较类型  map  slice func  只能跟nil 对比

func TestIndexFirstFunc(t *testing.T) {
	type student struct {
		age   string
		hobby []string
		mark  map[string]int
	}

	tests := []struct {
		src   []int
		dst   int
		equal func(int, int) bool
		want  int
		name  string
	}{
		{
			src: []int{1, 2, 3, 5},
			dst: 1,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			want: 0,
			name: "first one",
		},
		{
			src: []int{},
			dst: 1,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst: 1,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			want: -1,
			name: "src nil",
		},
		{
			src: []int{1, 4, 6},
			dst: 7,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			want: -1,
			name: "dst not exist",
		},
		{
			src: []int{1, 3, 4, 2, 0},
			dst: 0,
			equal: func(i int, i2 int) bool {
				return i == i2
			},
			want: 4,
			name: "last one",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexFirstFunc[int](test.src, test.dst, test.equal))
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
		src   []student
		dst   student
		equal func(student, student) bool
		want  int
		name  string
	}{
		{
			src: []student{student{
				name:  "小宏",
				hobby: []string{"唱", "跳", "rapper"},
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			}},
			dst: student{
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
			want: 0,
			name: "only one",
		},

		{
			src: []student{student{
				name:  "小宏",
				hobby: []string{"唱", "跳", "rapper"},
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			}},
			dst: student{
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
			want: -1,
			name: "only two",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexFirstFunc[student](test.src, test.dst, test.equal))
		})
	}
}
