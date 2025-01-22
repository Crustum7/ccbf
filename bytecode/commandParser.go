package bytecode

type commandNode struct {
	final bool
	next  map[byte]*commandNode
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

// func (node *commandNode) printAll(prev string) {
// 	if node.final {
// 		fmt.Println(prev)
// 	}

// 	for key, val := range node.next {
// 		val.printAll(prev + string(key))
// 	}
// }

type CommandParser struct {
	topNode *commandNode
}

func InitCommandParser(patterns []string) CommandParser {
	var parser CommandParser
	parser.topNode = &commandNode{final: false, next: make(map[byte]*commandNode)}

	for _, pattern := range patterns {
		parser.topNode.addPattern(pattern)
	}

	return parser
}

func (parser *CommandParser) FindPattern(pattern string) string {
	return parser.topNode.findPattern(pattern)
}
