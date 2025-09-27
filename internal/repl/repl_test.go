package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input			string
		expected 	[]string
	}{
		// basic
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// multiple spaces
	{
		input:    "foo     bar",
		expected: []string{"foo", "bar"},
	},
	// tabs and newlines
	{
		input:    "\tGo\nLang\tRocks\n",
		expected: []string{"go", "lang", "rocks"},
	},
	// mixed case
	{
		input:    "HeLLo WoRLd",
		expected: []string{"hello", "world"},
	},
	// single word
	{
		input:    "   single   ",
		expected: []string{"single"},
	},
	// empty string
	{
		input:    "",
		expected: []string{},
	},
	// only spaces
	{
		input:    "     ",
		expected: []string{},
	},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("slice len does not match: expected %d, got %d", len(c.expected), len(actual))
		}

		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word does not match: expected %q, got %q", expectedWord, word)
			}
		}
	}

}


