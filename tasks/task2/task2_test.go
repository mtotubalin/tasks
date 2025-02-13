package task2

import (
	"fmt"
	"tasks/pkg/slice_compr"
	"testing"
)

func TestMakeSliceOfNumbersReverse(t *testing.T) {
	var tests = []struct {
		n        int
		expected []int
	}{
		{n: 0, expected: []int{}},
		{n: 1, expected: []int{1}},
		{n: 3, expected: []int{3, 2, 1}},
		{n: 5, expected: []int{5, 4, 3, 2, 1}},
		{n: 8, expected: []int{8, 7, 6, 5, 4, 3, 2, 1}},
		{n: 10, expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("n = %d", tt.n)
		t.Run(testname, func(t *testing.T) {
			actual := makeSliceOfNumbersReverse(tt.n)
			if !slice_compr.Equal(actual, tt.expected) {
				t.Errorf("got %d, expected %d", actual, tt.expected)
			}
		})
	}
}
