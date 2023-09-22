package go121

import (
	"slices"
	"testing"
)

func TestSliceContains(t *testing.T) {
	tests := []struct {
		input    []any
		search   any
		expected any
		label    string
	}{
		{[]any{1, 2, 3}, 2, true, "[]int success"},
		{[]any{"a", "b", "c"}, "c", true, "[]string success"},
		{[]any{1, 2, 3}, 4, false, "[]int failure"},
		{[]any{"a", "b", "c"}, "d", false, "[]string failure"},
	}

	for _, test := range tests {
		result := slices.Contains(test.input, test.search)
		if result != test.expected {
			t.Errorf("sliceContains(%v) = %t, want %t (%s)", test.input, result, test.expected, test.label)
		}
	}
}

func TestSliceContainsFunc(t *testing.T) {
	tests := []struct {
		input    []any
		f        func(any) bool
		expected any
		label    string
	}{
		{[]any{1, 2, -3}, func(a any) bool { return a.(int) < 0 }, true, "[]int has negative success"},
		{[]any{"a", "b", ""}, func(a any) bool { return a.(string) == "" }, true, "[]string has empty success"},
		{[]any{1, 2, 3}, func(a any) bool { return a.(int) < 0 }, false, "[]int has negative failure"},
		{[]any{"a", "b", "c"}, func(a any) bool { return a.(string) == "" }, false, "[]string has empty failure"},
	}

	for _, test := range tests {
		result := slices.ContainsFunc(test.input, test.f)
		if result != test.expected {
			t.Errorf("sliceContains(%v) = %t, want %t (%s)", test.input, result, test.expected, test.label)
		}
	}
}
