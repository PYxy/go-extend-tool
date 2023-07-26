package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
			wartError:  ElementNotExist,
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := DeleteByIndex[int](test.src, test.indexSlice...)

			assert.Equal(t, test.wantSlice, result)
			assert.Equal(t, test.wartError, err)
		})
	}
}

func TestDeleteByValueAsc(t *testing.T) {
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
			name:      "first one",
		},
		{
			src:       []int{},
			tagVal:    1,
			count:     -1,
			wantSlice: []int{},
			wartError: ElementNotExist,
			name:      "first one",
		},
		{
			src:       []int{1, 2, 1, 5},
			tagVal:    1,
			count:     -1,
			wantSlice: []int{2, 5},
			wartError: nil,
			name:      "delete all 1",
		},
		{
			src:       []int{5, 2, 1, 5},
			tagVal:    5,
			count:     3,
			wantSlice: []int{2, 1},
			wartError: nil,
			name:      "删除111",
		},
		{
			src:       []int{5, 2, 1, 5},
			tagVal:    2,
			count:     -1,
			wantSlice: []int{5, 1, 5},
			wartError: nil,
			name:      "删除6996",
		},
		{
			src:       []int{1, 2, 1, 5},
			tagVal:    5,
			count:     10,
			wantSlice: []int{1, 2, 1},
			wartError: nil,
			name:      "删除222",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := DeleteByValueAsc[int](test.src, test.tagVal, test.count)

			assert.Equal(t, test.wantSlice, result)
			assert.Equal(t, test.wartError, err)
		})
	}
}
