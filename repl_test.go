package main

import (
	"testing"
) 

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Pikachu   BULBASAUR charmander",
			expected: []string{"pikachu", "bulbasaur", "charmander"},
		},
		{
			input: "ThePrimeagen",
			expected: []string{"theprimeagen"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slice lengths don't match %v != %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words aren't equal: %v != %v", word, expectedWord)
			}
		}
	}
}

