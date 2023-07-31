package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		src        []int
		wartResurt []int
		wartError  error
		name       string
	}{
		{
			src:        []int{1, 1, 1, 2, 3, 4, 4, 5, 6},
			wartResurt: []int{1, 2, 3, 4, 5, 6},
			wartError:  nil,
			name:       "正常操作",
		},
		{
			src:        nil,
			wartResurt: nil,
			wartError:  SourceSliceIsNil,
			name:       "nil 切片",
		},
		{
			src:        []int{},
			wartResurt: []int{},
			wartError:  nil,
			name:       "长度为0 的切片",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Unique[int](test.src)
			//fmt.Println(reflect.DeepEqual(test.result,result))
			if err != nil {
				assert.Equal(t, test.wartError, err)
			} else {
				assert.Equal(t, test.wartError, err)
				assert.Equal(t, test.wartResurt, result)
			}

		})
	}
}

func TestUniqueFuncFirst(t *testing.T) {
	type student struct {
		name  string
		hobby []string
		mark  map[string]int
	}

	tests := []struct {
		src       []student
		getKey    GetKey[student, int]
		wantSLice []student
		wantErr   error
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
						"数学": 17,
						"语文": 18,
					},
				},
				student{
					name:  "小白2",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
				student{
					name:  "小东",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 17,
						"语文": 18,
					},
				},
			},
			wantSLice: []student{
				student{
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
						"数学": 17,
						"语文": 18,
					},
				},
			},
			getKey: func(s student) int {
				return s.mark["数学"]
			},
			wantErr: nil,
			name:    "根据学生成绩中的数学 去重",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := UniqueFuncFirst(test.src, test.getKey)
			fmt.Println(res)
			assert.Equal(t, test.wantSLice, res)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestUniqueFuncLast(t *testing.T) {
	type student struct {
		name  string
		hobby []string
		mark  map[string]int
	}

	tests := []struct {
		src       []student
		getKey    GetKey[student, int]
		wantSLice []student
		wantErr   error
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
						"数学": 17,
						"语文": 18,
					},
				},
				student{
					name:  "小白2",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
				student{
					name:  "小东",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 17,
						"语文": 18,
					},
				},
			},
			wantSLice: []student{
				student{
					name:  "小白2",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
				student{
					name:  "小东",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 17,
						"语文": 18,
					},
				},
			},
			getKey: func(s student) int {
				return s.mark["数学"]
			},
			wantErr: nil,
			name:    "根据学生成绩中的数学 去重",
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
						"数学": 17,
						"语文": 18,
					},
				},

				student{
					name:  "小东",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 19,
						"语文": 18,
					},
				},
				student{
					name:  "小白2",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
			},
			wantSLice: []student{
				student{
					name:  "小白",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 17,
						"语文": 18,
					},
				},

				student{
					name:  "小东",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 19,
						"语文": 18,
					},
				},
				student{
					name:  "小白2",
					hobby: []string{"唱", "飞", "装"},
					mark: map[string]int{
						"数学": 15,
						"语文": 16,
					},
				},
			},
			getKey: func(s student) int {
				return s.mark["数学"]
			},
			wantErr: nil,
			name:    "根据学生成绩中的数学去重 去除第一个",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := UniqueFuncLast(test.src, test.getKey)
			fmt.Println(res)
			assert.Equal(t, test.wantSLice, res)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
