package bytecode

import "testing"

func testFindPattern(t *testing.T, parser CommandParser, text string, expectedPattern string) {
	pattern := parser.FindPattern(text)
	if pattern != expectedPattern {
		t.Fatalf("FindPatterns should have returned \"%s\", but resturned \"%s\" instead, with input \"%s\"", expectedPattern, pattern, text)
	}
}

func testFindPatternRepetitions(t *testing.T, parser CommandParser, text string, expectedPattern string, expectedRepetions int) {
	pattern, repetitions := parser.FindPatternReapetions(text)
	if pattern != expectedPattern {
		t.Fatalf("FindPatterns should have returned \"%s\", but resturned \"%s\" instead, with input \"%s\"", expectedPattern, pattern, text)
	}
	if repetitions != expectedRepetions {
		t.Fatalf("FindPatterns should have matched pattern \"%s\" %d times, but only found %d for input \"%s\"", pattern, repetitions, expectedRepetions, expectedPattern)
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

	testFindPatternRepetitions(t, parser, "<..<...", "<..", 2)
	testFindPatternRepetitions(t, parser, "<..<..", "<..", 2)
	testFindPatternRepetitions(t, parser, "<<<<<<<<<", "<<", 4)
	testFindPatternRepetitions(t, parser, ".", "", 1)
}
