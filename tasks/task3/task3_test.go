package task3

import (
	"fmt"
	"testing"
)

func TestCountNumberInSlice(t *testing.T) {
	var tests = []struct {
		numbers  []int
		search   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 1},
		{[]int{3, 2, 3, 3, 1}, 3, 3},
		{[]int{2, 2, 3, 3, 1}, 3, 2},
		{[]int{2, 4, 4, 3, 1}, 1, 1},
		{[]int{}, 1, 0},
		{[]int{1}, 2, 0},
		{[]int{1}, 1, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("slice=%d,search=%d", tt.numbers, tt.search)
		t.Run(testname, func(t *testing.T) {
			actual := countNumberInSlice(tt.numbers, tt.search)
			if actual != tt.expected {
				t.Errorf("got %d, expected %d", actual, tt.expected)
			}
		})
	}
}
