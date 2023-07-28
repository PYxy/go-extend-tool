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
