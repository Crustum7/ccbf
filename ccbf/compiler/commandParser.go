package compiler

import (
	"fmt"
	"regexp"

	"martinjonson.com/ccbf/ccbf/operations"
	"martinjonson.com/ccbf/ccbf/utils"
)

type CommandParser struct {
	reMap utils.RegexMap[operations.Operation]
}

type ParsedCommand struct {
	Pattern   string
	Match     string
	Groups    []string
	Operation *operations.Operation
}

func InitCommandParser(patterns []string, ops []operations.Operation) CommandParser {
	m := make(map[string]operations.Operation)
	for i := range patterns {
		pattern := addStartAnchor(patterns[i])
		m[pattern] = ops[i]
	}

	reMap := utils.InitRegexMap(m)
	return CommandParser{reMap: reMap}
}

func addStartAnchor(str string) string {
	return fmt.Sprintf("^%s", str)
}

func (parser *CommandParser) FindLongest(feed string) ParsedCommand {
	pattern := parser.reMap.FindLongestMatchPattern(feed)
	operation := parser.reMap.GetValueFromPattern(pattern)
	match, groups := getMatchAndGroups(feed, pattern)

	if pattern != "" {
		pattern = pattern[1:]
	}
	parsed := ParsedCommand{
		Pattern:   pattern,
		Operation: operation,
		Match:     match,
		Groups:    groups,
	}

	return parsed
}

func getMatchAndGroups(str, pattern string) (string, []string) {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(str)
	return match[0], match[1:]
}
