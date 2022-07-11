package function

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipe2IntString(t *testing.T) {
	table := []struct {
		name     string
		fn1      func(int) int
		fn2      func(int) string
		input    int
		expected string
	}{
		{
			name:     "simple",
			fn1:      func(i int) int { return i + 1 },
			fn2:      strconv.Itoa,
			input:    1,
			expected: "2",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Pipe2(tt.fn1, tt.fn2)(tt.input))
		})
	}
}
