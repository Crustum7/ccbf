package compiler

import (
	"fmt"

	"martinjonson.com/ccbf/ccbf/operations"
	"martinjonson.com/ccbf/ccbf/utils"
)

type CommandParser2 struct {
	reMap utils.RegexMap[operations.Operation]
}

func InitCommandParser2(patterns []string, ops []operations.Operation) (CommandParser2, error) {
	m := make(map[string]operations.Operation)
	for i := range patterns {
		pattern := addStartAnchor(patterns[i])
		m[pattern] = ops[i]
	}

	reMap, err := utils.InitRegexMap(m)
	return CommandParser2{reMap: reMap}, err
}

func addStartAnchor(str string) string {
	return fmt.Sprintf("^%s", str)
}

func (parser *CommandParser2) FindLongest(feed string) (string, *operations.Operation) {
	return parser.reMap.FindLongestMatch(feed)
}
