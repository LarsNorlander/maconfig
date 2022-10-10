package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNumberList(t *testing.T) {
	tests := []struct {
		name      string
		arg       string
		expected  []int
		shouldErr bool
	}{
		{
			name:     "single number",
			arg:      "1",
			expected: []int{1},
		},
		{
			name:     "single descending range",
			arg:      "5-1",
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "single ascending range",
			arg:      "1-5",
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "single number repetition",
			arg:      "1*3",
			expected: []int{1, 1, 1},
		},
		{
			name:     "range repetition",
			arg:      "1-2*3",
			expected: []int{1, 2, 1, 2, 1, 2},
		},
		{
			name:      "invalid range",
			arg:       "1-2-3",
			shouldErr: true,
		},
		{
			name:     "singe range of same number",
			arg:      "1-1",
			expected: []int{1},
		},
		{
			name:     "multiple numbers ranges and repetitions",
			arg:      "1,5-7,6,10*3,1-2*2",
			expected: []int{1, 5, 6, 7, 6, 10, 10, 10, 1, 2, 1, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ParseNumberList(test.arg)
			assert.Equal(t, test.shouldErr, err != nil)
			assert.Equal(t, test.expected, result)
		})
	}
}
