package bytecode

import "testing"

func testNext(t *testing.T, parser *ProgramParser, expectedPattern string, expectedRepetitions int) {
	pattern, repetitions := parser.next()
	if pattern != expectedPattern {
		t.Fatalf("Pattern \"%s\" does not match \"%s\"", pattern, expectedPattern)
	}

	if repetitions != expectedRepetitions {
		t.Fatalf("Repetitions \"%d\" does not match \"%d\"", repetitions, expectedRepetitions)
	}
}

func expectHasNext(t *testing.T, parser ProgramParser) {
	if !parser.hasNext() {
		t.Fatalf("Program should not be completed at this point")
	}
}

func TestProgramParser(t *testing.T) {
	patterns := []string{"a", "b", "cc"}
	program := "aabccbaa"
	commandParser := InitCommandParser(patterns)
	parser := InitProgramParser(program, commandParser)

	expectHasNext(t, parser)
	testNext(t, &parser, "a", 2)
	parser.skipRepetitions(2)

	expectHasNext(t, parser)
	testNext(t, &parser, "b", 1)
	parser.skipRepetitions(1)

	expectHasNext(t, parser)
	testNext(t, &parser, "cc", 1)
	parser.skipRepetitions(2)

	expectHasNext(t, parser)
	testNext(t, &parser, "b", 1)
	parser.skipRepetitions(1)

	expectHasNext(t, parser)
	testNext(t, &parser, "a", 2)
	parser.skipRepetitions(2)

	if parser.hasNext() {
		t.Fatalf("Program should be completed at this point")
	}
}
