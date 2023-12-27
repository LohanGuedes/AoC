package day1

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		fn       func(string) int
	}{
		{
			expected: 142,
			input:    "test1.txt",
			fn:       part1,
		},
	}
	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)), test.input)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		fn       func(string) int
	}{
		{
			expected: 89,
			input:    "input.txt",
			fn:       part2,
		},
	}
	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)), test.input)
	}
}
