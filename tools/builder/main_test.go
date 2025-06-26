package main

import "testing"

func TestParseHeader(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "# h1 heading",
			Expected: "<h1>h1 heading</h1>",
		},
		{
			Input:    "## h2 heading",
			Expected: "<h2>h2 heading</h2>",
		},
		{
			Input:    "### h3 heading",
			Expected: "<h3>h3 heading</h3>",
		},
	}

	for _, testCase := range cases {
		output, _ := ParseHeading(testCase.Input)
		if output != testCase.Expected {
			t.Errorf("Wrong heading output, expected: %s, got: %s", testCase.Expected, output)
		}
	}
}

func TestParseBold(t *testing.T) {
	input := "**bold**"
	expected := "<b>bold</b>"

	output, _ := ParseBold(input)
	if output != expected {
		t.Errorf("Wrong bold output, expected: %s, got: %s", expected, output)
	}
}

func TestParse(t *testing.T) {
	input := "# test **bold heading**"
	expected := "<h1>test <b>bold heading</b></h1>"
	output, _ := Parse(input)

	if output != expected {
		t.Errorf("Wrong bold output, expected: %s, got: %s", expected, output)
	}
}
