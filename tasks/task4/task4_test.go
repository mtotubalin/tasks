package task4

import (
	"fmt"
	"testing"
)

func TestIsEven(t *testing.T) {
	var tests = []struct {
		number   int
		expected bool
	}{
		{-2, true},
		{-1, false},
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, true},
		{7, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("number=%d", tt.number)
		t.Run(testname, func(t *testing.T) {
			actual := IsEven(tt.number)
			if actual != tt.expected {
				t.Errorf("got %t, expected %t", actual, tt.expected)
			}
		})
	}
}
