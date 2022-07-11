package function

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Chlor87/go-data-manipulation/slice"
)

func TestCurry2MapIntString(t *testing.T) {
	curried := Curry2(slice.Map[int, string])
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
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, curried(tt.fn)(tt.input))
		})
	}
}
