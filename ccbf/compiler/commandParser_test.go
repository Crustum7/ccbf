package compiler_test

import (
	"fmt"
	"testing"

	"martinjonson.com/ccbf/ccbf/compiler"
	"martinjonson.com/ccbf/ccbf/operations"
)

func TestCommandParser(t *testing.T) {
	testcases := []struct {
		commands        []string
		ops             []operations.Operation
		feed            string
		expectedPattern string
		expectedMatch   string
		expectedGroups  []string
	}{
		{
			commands:        []string{">"},
			feed:            "abcde",
			expectedPattern: "",
			expectedMatch:   "",
		},
		{
			commands:        []string{">"},
			feed:            "abcde",
			expectedPattern: "",
			expectedMatch:   "",
		},
		{
			commands:        []string{">"},
			feed:            ">>>>",
			expectedPattern: ">",
			expectedMatch:   ">",
		},
		{
			commands:        []string{">"},
			feed:            "a>>>>",
			expectedPattern: "",
			expectedMatch:   "",
		},
		{
			commands:        []string{`\+`},
			feed:            "++++",
			expectedPattern: `\+`,
			expectedMatch:   "+",
		},
		{
			commands:        []string{`\++`},
			feed:            "++++",
			expectedPattern: `\++`,
			expectedMatch:   "++++",
		},
		{
			commands:        []string{`\[-(>+)\+(<+)\]`},
			feed:            "[->>+<<]",
			expectedPattern: `\[-(>+)\+(<+)\]`,
			expectedMatch:   "[->>+<<]",
			expectedGroups:  []string{">>", "<<"},
		},
		{
			commands:        []string{">", ">>", ">"},
			feed:            ">>",
			expectedPattern: ">>",
			expectedMatch:   ">>",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%v, %s, %s", tc.commands, tc.feed, tc.expectedPattern), func(t *testing.T) {
			parser := compiler.InitCommandParser(tc.commands, make([]operations.Operation, 5))

			parsedCommand := parser.FindLongest(tc.feed)

			if parsedCommand.Pattern != tc.expectedPattern {
				t.Fatalf("FindLongest should have found %s but found %s", tc.expectedPattern, parsedCommand.Pattern)
			}

			if parsedCommand.Match != tc.expectedMatch {
				t.Fatalf("FindLongest should have found match %s but found %s", tc.expectedMatch, parsedCommand.Match)
			}

			if len(parsedCommand.Groups) != len(tc.expectedGroups) {
				t.Fatalf("FindLongest should have found groups %v but found %v", tc.expectedGroups, parsedCommand.Groups)
			}

			for i, group := range tc.expectedGroups {
				if parsedCommand.Groups[i] != group {
					t.Fatalf("FindLongest should have found groups %v but found %v", tc.expectedGroups, parsedCommand.Groups)
				}
			}
		})
	}
}
