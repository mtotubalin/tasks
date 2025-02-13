package task1

import (
	"fmt"
	"tasks/pkg/slice_compr"
	"testing"
)

func TestMakeSliceOfNumbers(t *testing.T) {
	var tests = []struct {
		n        int
		expected []int
	}{
		{n: 0, expected: []int{}},
		{n: 1, expected: []int{1}},
		{n: 3, expected: []int{1, 2, 3}},
		{n: 5, expected: []int{1, 2, 3, 4, 5}},
		{n: 8, expected: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{n: 10, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("n = %d", tt.n)
		t.Run(testname, func(t *testing.T) {
			actual := makeSliceOfNumbers(tt.n)
			if !slice_compr.Equal(actual, tt.expected) {
				t.Errorf("got %d, expected %d", actual, tt.expected)
			}
		})
	}
}
