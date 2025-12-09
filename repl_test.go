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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "ThIs Is A tEsT",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    " gotta  catch  them  all",
			expected: []string{"gotta", "catch", "them", "all"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("failed test: actual length (%v) does not equal expected (%v)", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("failed test: word (%v) does not match expected (%v)", word, expectedWord)
			}
		}
	}
}
