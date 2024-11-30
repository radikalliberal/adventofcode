package utils

import (
	"os"
	"testing"
)

func TestReadPipe(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Multiple Lines",
			input:    "line1\nline2\nline3\n",
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "Empty Input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Single Line",
			input:    "singleline\n",
			expected: []string{"singleline"},
		},
		{
			name:     "Input with Empty Lines",
			input:    "line1\n\nline2\n\nline3\n",
			expected: []string{"line1", "", "line2", "", "line3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate stdin input
			r, w, _ := os.Pipe()
			_, _ = w.WriteString(tt.input)
			w.Close()
			os.Stdin = r

			result, e := ReadPipe()
			if e != nil {
				t.Errorf("Error reading pipe: %s", e)
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d lines, got %d", len(tt.expected), len(result))
			}

			for i, line := range result {
				if line != tt.expected[i] {
					t.Errorf("Expected line %d to be %s, got %s", i, tt.expected[i], line)
				}
			}
		})
	}
}
