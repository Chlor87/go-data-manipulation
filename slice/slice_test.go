package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapIntToString(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		fn       func(int) string
		expected []string
	}{
		{
			name:     "simple",
			input:    []int{1, 2, 3},
			fn:       strconv.Itoa,
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "empty input",
			input:    nil,
			fn:       strconv.Itoa,
			expected: []string{},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Map(tt.fn, tt.input))
		})
	}
}

func TestFilterInt(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		fn       func(int) bool
		expected []int
	}{
		{
			name:  "simple",
			input: []int{1, 2, 3},
			fn: func(i int) bool {
				return i > 1
			},
			expected: []int{2, 3},
		},
		{
			name:     "empty input",
			input:    nil,
			fn:       nil,
			expected: []int{},
		},
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Filter(tt.fn, tt.input))
		})
	}
}
