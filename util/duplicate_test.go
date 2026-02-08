package util

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	type MyString string
	type MyStrings []MyString
	tests := []struct {
		name     string
		input    any
		remove   any
		expected any
	}{
		{
			name:     "Test_-_string_slice",
			input:    []string{"a", "b", "a", "c", "a"},
			remove:   "a",
			expected: []string{"b", "c"},
		},
		{
			name:     "Test_-_named_param_string_slice",
			input:    MyStrings{"a", "b", "a", "c", "a"},
			remove:   MyString("a"),
			expected: MyStrings{"b", "c"},
		},
		{
			name:     "Test_-_int_slice",
			input:    []int{1, 2, 3, 2, 4, 2},
			remove:   2,
			expected: []int{1, 3, 4},
		},
		{
			name:     "Test_-_float_slice",
			input:    []float32{1.1, 2.2, 1.1, 3.3},
			remove:   float32(1.1),
			expected: []float32{2.2, 3.3},
		},
		{
			name:     "Test_-_double_slice",
			input:    []float64{1.1, 2.2, 1.1, 3.3},
			remove:   1.1,
			expected: []float64{2.2, 3.3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result any

			switch input := tt.input.(type) {
			case []string:
				result = RemoveDuplicate(input, tt.remove.(string))
			case []int:
				result = RemoveDuplicate(input, tt.remove.(int))
			case []float32:
				result = RemoveDuplicate(input, tt.remove.(float32))
			case []float64:
				result = RemoveDuplicate(input, tt.remove.(float64))
			case MyStrings:
				result = RemoveDuplicate(input, tt.remove.(MyString))
			default:
				t.Fatalf("unsupported type: %T", tt.input)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v %T, got %v %T", tt.expected, tt.expected, result, result)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	type MyString string
	type MyStrings []MyString
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{
			name:     "Test_-_string_slice",
			input:    []string{"a", "b", "a", "c", "a"},
			expected: true,
		},
		{
			name:     "Test_-_string_slice_false",
			input:    []string{"a", "b", "d", "c", "e"},
			expected: false,
		},
		{
			name:     "Test_-_named_param_string_slice",
			input:    MyStrings{"a", "b", "a", "c", "a"},
			expected: true,
		},
		{
			name:     "Test_-_int_slice",
			input:    []int{1, 2, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Test_-_int_slice_false",
			input:    []int{1, 2, 3, 5},
			expected: false,
		},
		{
			name:     "Test_-_float_slice",
			input:    []float32{1.1, 2.2, 1.1, 3.3},
			expected: true,
		},
		{
			name:     "Test_-_float_slice_false",
			input:    []float32{1.1, 2.2, 3.3},
			expected: false,
		},
		{
			name:     "Test_-_double_slice",
			input:    []float64{1.1, 2.2, 1.1, 3.3},
			expected: true,
		},
		{
			name:     "Test_-_double_slice_false",
			input:    []float64{1.15, 2.20, 3.33},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result any

			switch input := tt.input.(type) {
			case []string:
				result = ContainsDuplicate(input)
			case []int:
				result = ContainsDuplicate(input)
			case []float32:
				result = ContainsDuplicate(input)
			case []float64:
				result = ContainsDuplicate(input)
			case MyStrings:
				result = ContainsDuplicate(input)
			default:
				t.Fatalf("unsupported type: %T", tt.input)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v %T, got %v %T", tt.expected, tt.expected, result, result)
			}
		})
	}
}
