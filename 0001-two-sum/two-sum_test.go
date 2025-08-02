package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "case1 from leetcode testcase 1",
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "case2 from leetcode testcase 2",
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "case 3 from leetcode testcase 3",
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "case 4 no solution",
			input:    []int{1, 2, 3, 4},
			target:   10,
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.input, tc.target)
			assert.ElementsMatch(t, tc.expected, result)
		})
	}
}
