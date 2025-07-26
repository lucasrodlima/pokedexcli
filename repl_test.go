package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  leading and trailing spaces  ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "Some More         tests",
			expected: []string{"some", "more", "tests"},
		},
		{
			input:    "CHARMANDER  Bulbasaur PikaCHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput output length does not match expected")
		}

		for i := range actual {
			expectedWord := c.expected[i]
			actualWord := actual[i]

			if actualWord != expectedWord {
				t.Errorf("cleanInput output word does not match expected")
			}
		}
	}
}
