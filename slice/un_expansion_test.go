package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnExpansion(t *testing.T) {
	a := make([]int, 0, 10)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)

	b := make([]int, 0, 10)
	b = append(b, 1)
	b = append(b, 2)
	b = append(b, 3)

	c := make([]int, 0, 10)
	c = append(c, 1)
	c = append(c, 2)
	c = append(c, 3)
	c = append(c, 4)

	e := make([]int, 0, 10)
	tests := []struct {
		src       []int
		hard      bool
		wantSlice []int
		wantCap   int
		wartError error
		name      string
	}{
		{
			src:       []int{},
			hard:      false,
			wantSlice: []int{},
			wantCap:   0,
			wartError: nil,
			name:      "src length is 0",
		},
		{
			src:       nil,
			hard:      false,
			wantSlice: nil,
			wantCap:   0,
			wartError: SourceSliceIsNil,
			name:      "src length is 0",
		},
		{
			src:       a,
			hard:      false,
			wantCap:   5,
			wantSlice: nil,
			wartError: nil,
			name:      "cut half",
		},
		{
			src:       b,
			hard:      true,
			wantCap:   3,
			wantSlice: nil,
			wartError: nil,
			name:      "cap == len",
		},
		{
			src:       c,
			hard:      false,
			wantCap:   10,
			wantSlice: nil,
			wartError: nil,
			name:      "pass",
		},
		{
			src:       e,
			hard:      true,
			wantCap:   0,
			wantSlice: nil,
			wartError: nil,
			name:      "cap 10 len 0  change to cap 0 len 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := UnExpansion[int](test.src, test.hard)

			//assert.Equal(t, test.wantSlice, result)
			assert.Equal(t, test.wantCap, cap(result))
			assert.Equal(t, test.wartError, err)
		})
	}
}
