package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "case1 from leetcode testcase 1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "case2 from leetcode testcase 2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
		{
			name:     "case3 when an empty slice",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := removeDuplicates(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
