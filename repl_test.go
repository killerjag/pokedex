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
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "krl, que foda",
			expected: []string{"krl,", "que", "foda"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "GOOOOOOOOL!",
			expected: []string{"gooooooool!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Wrong length")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words doesn't match: %s vs %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
