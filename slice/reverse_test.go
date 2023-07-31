package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		src         []int
		resultSlice []int
		name        string
	}{
		{
			resultSlice: []int{},
			name:        "源切片为nil",
		},
		{
			src:         []int{1, 2, 2, 3},
			resultSlice: []int{3, 2, 2, 1},
			name:        "双数 切片反转",
		},
		{
			src:         []int{1, 2, 2, 3, 4},
			resultSlice: []int{4, 3, 2, 2, 1},
			name:        "单数 切片反转",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			assert.Equal(t, test.resultSlice, Reverse(test.src))
		})
	}
}

func TestReverseOnSelf(t *testing.T) {
	tests := []struct {
		src         []int
		resultSlice []int
		name        string
	}{
		{
			resultSlice: []int{},
			name:        "源切片为nil",
		},
		{
			src:         []int{1, 2, 2, 3},
			resultSlice: []int{3, 2, 2, 1},
			name:        "双数 切片反转",
		},
		{
			src:         []int{1, 2, 2, 3, 4},
			resultSlice: []int{4, 3, 2, 2, 1},
			name:        "单数 切片反转",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			assert.Equal(t, test.resultSlice, ReverseOnSelf(test.src))
		})
	}
}
