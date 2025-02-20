package compiler

type CommandParser struct {
	topNode *commandNode
}

type commandNode struct {
	final bool
	next  map[byte]*commandNode
}

func InitCommandParser(patterns []string) CommandParser {
	var parser CommandParser
	parser.topNode = &commandNode{final: false, next: make(map[byte]*commandNode)}

	for _, pattern := range patterns {
		parser.topNode.addPattern(pattern)
	}

	return parser
}

func (parser *CommandParser) FindPattern(text string) string {
	return parser.topNode.findPattern(text)
}

func (parser *CommandParser) FindPatternReapetions(text string) (string, int) {
	pattern := parser.FindPattern(text)
	if pattern == "" {
		return "", 1
	}

	repetitionsAfterFirst := findRepetition(text[len(pattern):], pattern)
	return pattern, repetitionsAfterFirst + 1
}

func findRepetition(text string, pattern string) int {
	pattLen := len(pattern)
	for i := 0; i < len(text); i++ {
		offset := i * pattLen
		if offset+pattLen > len(text) {
			return i
		}

		comp := text[offset : offset+pattLen]
		if pattern != comp {
			return i
		}
	}
	return len(text) / len(pattern)
}

func (node *commandNode) findPattern(pattern string) string {
	if pattern == "" {
		return ""
	}

	char := pattern[0]
	nextNode, nextExists := node.next[char]
	if !nextExists {
		return ""
	}

	nextPattern := nextNode.findPattern(pattern[1:])
	patternFound := nextPattern != "" || nextNode.final
	if patternFound {
		return string(char) + nextPattern
	}

	return ""
}

func (node *commandNode) addPattern(pattern string) {
	if len(pattern) < 1 {
		node.final = true
		return
	}

	char := pattern[0]
	nextNode, nextExists := node.next[char]
	if !nextExists {
		node.next[char] = &commandNode{final: false, next: make(map[byte]*commandNode)}
		nextNode = node.next[char]
	}

	nextNode.addPattern(pattern[1:])
}
