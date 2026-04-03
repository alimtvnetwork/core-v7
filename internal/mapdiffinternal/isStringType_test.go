package mapdiffinternal

import "testing"

func Test_isStringType(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"string value returns true", "hello", true},
		{"int value returns false", 42, false},
		{"bool value returns false", true, false},
		{"nil value returns false", nil, false},
		{"empty string returns true", "", true},
		{"float value returns false", 3.14, false},
		{"slice value returns false", []int{1, 2}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isStringType(tc.input)
			if result != tc.expected {
				t.Errorf("isStringType(%v) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
