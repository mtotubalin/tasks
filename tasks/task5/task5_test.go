package task5

import (
	"fmt"
	"testing"
)

func TestCountEvenNumbers(t *testing.T) {
	var tests = []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{1, 2, 2, 3}, 2},
		{[]int{1, 3}, 0},
		{[]int{}, 0},
		{[]int{2, 4, 6, 8}, 4},
		{[]int{1, 2, 1, 2, 2, 5, 5, 4, 6}, 5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("numbers=%v", tt.numbers)
		t.Run(testname, func(t *testing.T) {
			actual := countEvenNumbers(tt.numbers)
			if actual != tt.expected {
				t.Errorf("got %d, expected %d", actual, tt.expected)
			}
		})
	}
}
