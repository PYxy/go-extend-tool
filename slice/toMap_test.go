package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMap(t *testing.T) {
	tests := []struct {
		src     []int
		wantMap map[int]struct{}
		name    string
	}{
		{
			src: []int{1, 2, 3, 4},
			wantMap: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
				4: struct{}{},
			},
			name: "正常测试",
		},
		{
			src: []int{1, 2, 3, 4},
			wantMap: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
				4: struct{}{},
			},
			name: "正常测试1",
		},
		{
			src: []int{1, 1, 3, 4},
			wantMap: map[int]struct{}{
				1: struct{}{},
				3: struct{}{},
				4: struct{}{},
			},
			name: "正常测试2",
		},
		{
			wantMap: map[int]struct{}{},
			name:    "空切片转为map",
		},
		{
			src:     []int{},
			wantMap: map[int]struct{}{},
			name:    "正常测试3 空map",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := ToMap[int](test.src)

			assert.Equal(t, test.wantMap, val)

		})
	}
}

func TestToMapFunc(t *testing.T) {
	type student struct {
		name string
		mark map[string]int
	}
	tests := []struct {
		src     []student
		wantMap map[int]student
		getkey  GetKey[student, int]
		name    string
	}{
		{
			src: []student{student{
				name: "小宏",
				mark: map[string]int{
					"数学": 15,
					"语文": 16,
				},
			},
				student{
					name: "小白",

					mark: map[string]int{
						"数学": 17,
						"语文": 18,
					},
				},
				student{
					name: "小白2",

					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
				student{
					name: "小东",

					mark: map[string]int{
						"数学": 17,
						"语文": 18,
					},
				},
			},
			wantMap: map[int]student{
				15: {
					name: "小宏",

					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
				17: student{
					name: "小白",

					mark: map[string]int{
						"数学": 17,
						"语文": 18,
					},
				},
			},
			getkey: func(s student) int {
				return s.mark["数学"]
			},
			name: "根据学生成绩中的数学成绩去重 并转为map",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := ToMapFunc[student, int](test.src, test.getkey)
			fmt.Println(val)
			assert.Equal(t, test.wantMap, val)

		})
	}
}
