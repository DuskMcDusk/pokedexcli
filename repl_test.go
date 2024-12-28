package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO hi",
			expected: []string{"hello", "hi"},
		},
	}

	for _, c := range cases {
		actual := cleaninput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Mismatching results length")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expecting %v actual %v", expectedWord, word)
			}
		}
	}
}
