package main

import (
	"testing"
)

func TestAddParseTimeParam(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard MySQL DSN",
			input:    "user:password@tcp(localhost:3306)/dbname",
			expected: "user:password@tcp(localhost:3306)/dbname?parseTime=true",
		},
		{
			name:     "DSN with existing params",
			input:    "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4",
			expected: "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=true",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := addParseTimeParam(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
