package bytecode

import "testing"

func testFindPattern(t *testing.T, parser CommandParser, pattern string, expectedResult string) {
	result := parser.FindPattern(pattern)
	if result != expectedResult {
		t.Fatalf("FindPatterns should have returned \"%s\", but resturned \"%s\" instead, with input \"%s\"", expectedResult, result, pattern)
	}
}

func TestCommandParser(t *testing.T) {
	patterns := []string{">", "<", "<<", "<.."}
	parser := InitCommandParser(patterns)

	testFindPattern(t, parser, "", "")
	testFindPattern(t, parser, ">", ">")
	testFindPattern(t, parser, ">.", ">")
	testFindPattern(t, parser, "<>", "<")
	testFindPattern(t, parser, "<<<", "<<")
	testFindPattern(t, parser, "<<<<", "<<")
	testFindPattern(t, parser, "<.", "<")
	testFindPattern(t, parser, "<...", "<..")
}
