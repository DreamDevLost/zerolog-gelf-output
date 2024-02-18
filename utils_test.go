package zerologgelfoutput

import (
	"testing"
)

// Test_getStringFromMap tests the getStringFromMap function with various scenarios.
func Test_getStringFromMap(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		m        map[string]interface{}
		key      string
		def      string
		expected string
	}{
		{
			name:     "Key exists with string value",
			m:        map[string]interface{}{"testKey": "testValue"},
			key:      "testKey",
			def:      "defaultValue",
			expected: "testValue",
		},
		{
			name:     "Key exists with non-string value",
			m:        map[string]interface{}{"testKey": 123},
			key:      "testKey",
			def:      "defaultValue",
			expected: "defaultValue",
		},
		{
			name:     "Key does not exist",
			m:        map[string]interface{}{"otherKey": "testValue"},
			key:      "testKey",
			def:      "defaultValue",
			expected: "defaultValue",
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getStringFromMap(tt.m, tt.key, tt.def)
			if result != tt.expected {
				t.Errorf("getStringFromMap(%v, %s, %s) = %s; want %s", tt.m, tt.key, tt.def, result, tt.expected)
			}
		})
	}
}
