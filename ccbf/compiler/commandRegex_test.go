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
		expectedCommand string
	}{
		{
			commands:        []string{">"},
			feed:            "abcde",
			expectedCommand: "",
		},
		{
			commands:        []string{">"},
			feed:            "abcde",
			expectedCommand: "",
		},
		{
			commands:        []string{">"},
			feed:            ">>>>",
			expectedCommand: ">",
		},
		{
			commands:        []string{">"},
			feed:            "a>>>>",
			expectedCommand: "",
		},
		{
			commands:        []string{"\\+"},
			feed:            "++++",
			expectedCommand: "+",
		},
		{
			commands:        []string{"\\++"},
			feed:            "++++",
			expectedCommand: "++++",
		},
		{
			commands:        []string{"\\[->+\\+<+\\]"},
			feed:            "[->>+<<]",
			expectedCommand: "[->>+<<]",
		},
		{
			commands:        []string{">", ">>", ">"},
			feed:            ">>",
			expectedCommand: ">>",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%v, %s, %s", tc.commands, tc.feed, tc.expectedCommand), func(t *testing.T) {
			parser, err := compiler.InitCommandParser2(tc.commands, make([]operations.Operation, 5))

			if err != nil {
				t.Fatalf("InitCommandParser should not have thrown error for commands %v", tc.commands)
			}

			command, _ := parser.FindLongest(tc.feed)

			if command != tc.expectedCommand {
				t.Fatalf("FindLongest should have found %s but found %s", tc.expectedCommand, command)
			}
		})
	}
}
