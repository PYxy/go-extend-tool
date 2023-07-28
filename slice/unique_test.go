package slice

import (
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
