package go121

import (
	"testing"
)

func TestMinInteger(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		label    string
	}{
		{[]int{1, 2, 3}, 1, "positive"},
		{[]int{-1, -2, -3}, -3, "negative"},
		{[]int{1, 1, 1}, 1, "all equal"},
		{[]int{1, 2, 1}, 1, "first and last"},
		{[]int{2, 1, 1}, 1, "first and second"},
	}

	for _, test := range tests {
		result := min(test.input[0], test.input[1], test.input[2])
		if result != test.expected {
			t.Errorf("min(%v) = %d, want %d (%s)", test.input, result, test.expected, test.label)
		}
	}
}

func TestMinString(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
		label    string
	}{
		{[]string{"a", "b", "c"}, "a", "positive"},
		{[]string{"", "foo", "bar"}, "", "empty"},
	}

	for _, test := range tests {
		result := min(test.input[0], test.input[1], test.input[2])
		if result != test.expected {
			t.Errorf("min(%v) = %s, want %s (%s)", test.input, result, test.expected, test.label)
		}
	}
}

func TestMaxInteger(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		label    string
	}{
		{[]int{1, 2, 3}, 3, "positive"},
		{[]int{-1, -2, -3}, -1, "negative"},
		{[]int{1, 1, 1}, 1, "all equal"},
		{[]int{1, 2, 1}, 2, "first and last"},
		{[]int{2, 1, 1}, 2, "first and second"},
	}

	for _, test := range tests {
		result := max(test.input[0], test.input[1], test.input[2])
		if result != test.expected {
			t.Errorf("max(%v) = %d, want %d (%s)", test.input, result, test.expected, test.label)
		}
	}
}

func TestMaxString(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
		label    string
	}{
		{[]string{"a", "b", "c"}, "c", "positive"},
		{[]string{"", "foo", "bar"}, "foo", "empty"},
	}

	for _, test := range tests {
		result := max(test.input[0], test.input[1], test.input[2])
		if result != test.expected {
			t.Errorf("max(%v) = %s, want %s (%s)", test.input, result, test.expected, test.label)
		}
	}
}

func TestClearMap(t *testing.T) {
	tests := []struct {
		input    map[string]any
		expected map[string]any
		label    string
	}{
		{map[string]any{"a": 1, "b": 2, "c": 3}, map[string]any{}, "empty map[string]int"},
		{map[string]any{"a": "b", "c": "d", "e": "f"}, map[string]any{}, "empty map[string]string"},
	}

	for _, test := range tests {
		clear(test.input)
		if len(test.input) != len(test.expected) {
			t.Errorf("clear(%v) = %v, want %v (%s)", test.input, test.input, test.expected, test.label)
		}
	}
}

func TestClearSlice(t *testing.T) {
	tests := []struct {
		input    []any
		expected []any
		label    string
	}{
		{[]any{1, 2, 3}, []any{nil, nil, nil}, "empty []int"},
		{[]any{"a", "b", "c"}, []any{nil, nil, nil}, "empty []string"},
		{[]any{1.1, 2.2, 3.3}, []any{nil, nil, nil}, "empty []float"},
	}

	for _, test := range tests {
		clear(test.input)
		if len(test.input) != len(test.expected) {
			t.Errorf("clear(%v) = %v, want %v (%s)", test.input, test.input, test.expected, test.label)
		}
	}
}
