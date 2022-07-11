package data

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Chlor87/go-data-manipulation/function"
	"github.com/Chlor87/go-data-manipulation/slice"
)

func TestGet(t *testing.T) {
	table := []struct {
		name     string
		key      any
		data     any
		ok       bool
		expected any
	}{
		{
			name:     "simple",
			key:      "a",
			data:     M{"a": "b"},
			ok:       true,
			expected: "b",
		},
		{
			name:     "nil map",
			key:      "a",
			data:     nil,
			ok:       false,
			expected: nil,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			res, ok := Get[any](tt.key, tt.data)
			assert.Equal(tt.ok, ok)
			assert.Equal(tt.expected, res)
		})
	}
}

func TestGetPath(t *testing.T) {
	table := []struct {
		name     string
		path     A
		data     any
		ok       bool
		expected any
	}{
		{
			name:     "simple map",
			path:     A{"a", "b"},
			data:     M{"a": M{"b": "c"}},
			ok:       true,
			expected: "c",
		},
		{
			name:     "not found",
			path:     A{"nothing", "b"},
			data:     M{"a": M{"b": "c"}},
			ok:       false,
			expected: nil,
		},
		{
			name:     "map and slice",
			path:     A{"a", 0, "b", 0},
			data:     M{"a": A{M{"b": A{"c"}}}},
			ok:       true,
			expected: "c",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			res, ok := GetPath[any](tt.path, tt.data)
			assert.Equal(tt.ok, ok)
			assert.Equal(tt.expected, res)
		})
	}
}

func TestModifyIntString(t *testing.T) {
	table := []struct {
		name     string
		key      any
		fn       func(int) string
		data     any
		expected any
	}{
		{
			name:     "simple",
			key:      "a",
			fn:       strconv.Itoa,
			data:     M{"a": 1},
			expected: M{"a": "1"},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, Modify("a", tt.fn, tt.data))
			assert.NotEqual(tt.expected, tt.data)
		})
	}
}

func TestModifyPathIntString(t *testing.T) {
	table := []struct {
		name     string
		fn       func(int) string
		path     A
		data     any
		expected any
	}{
		{
			name: "simple",
			fn: func(i int) string {
				return strconv.Itoa(i + 1)
			},
			path:     A{"a", "b", "c"},
			data:     M{"a": M{"b": M{"c": 1}}},
			expected: M{"a": M{"b": M{"c": "2"}}},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, ModifyPath(tt.path, tt.fn, tt.data))
			assert.NotEqual(tt.expected, tt.data)
		})
	}
}

func TestModifyPathSlice(t *testing.T) {
	curriedMap := function.Curry2(slice.Map[any, any])
	table := []struct {
		name     string
		fn       func(A) A
		path     A
		data     any
		expected any
	}{
		{
			name:     "map over nested slice",
			fn:       curriedMap(func(i any) any { return strconv.Itoa(i.(int) + 1) }),
			path:     A{"nested", "arr"},
			data:     M{"nested": M{"arr": A{1, 2, 3}}},
			expected: M{"nested": M{"arr": A{"2", "3", "4"}}},
		},
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, ModifyPath(tt.path, tt.fn, tt.data))
			assert.NotEqual(tt.expected, tt.data)
		})
	}
}

func TestSet(t *testing.T) {
	table := []struct {
		name     string
		key      any
		value    any
		data     any
		expected any
	}{
		{
			name:     "simple map",
			key:      "a",
			value:    "z",
			data:     M{"a": "b"},
			expected: M{"a": "z"},
		},
		{
			name:     "nil map",
			key:      "a",
			value:    "z",
			data:     nil,
			expected: M{"a": "z"},
		},
		{
			name:     "simple slice",
			key:      1,
			value:    "a",
			data:     A{},
			expected: A{nil, "a"},
		},
		{
			name:     "nil slice",
			key:      1,
			value:    "a",
			data:     nil,
			expected: A{nil, "a"},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, Set(tt.key, tt.value, tt.data))
			assert.NotEqual(tt.expected, tt.data)
		})
	}
}

func TestSetPath(t *testing.T) {
	table := []struct {
		name     string
		path     A
		value    any
		data     any
		expected any
	}{
		{
			name:     "simple map",
			path:     A{"a", "b", "c"},
			value:    "d",
			data:     M{},
			expected: M{"a": M{"b": M{"c": "d"}}},
		},
		{
			name:     "nil map",
			path:     A{"a", "b"},
			value:    "c",
			data:     nil,
			expected: M{"a": M{"b": "c"}},
		},
		{
			name:     "simple slice",
			path:     A{0, 0},
			value:    "a",
			data:     A{},
			expected: A{A{"a"}},
		},
		{
			name:     "wrong key type",
			path:     A{true},
			value:    "a",
			data:     A{},
			expected: nil,
		},
		{
			name:     "map and slice",
			path:     A{"a", 0, "b", 0},
			value:    "c",
			data:     nil,
			expected: M{"a": A{M{"b": A{"c"}}}},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, SetPath(tt.path, tt.value, tt.data))
			assert.NotEqual(tt.expected, tt.data)
		})
	}
}
