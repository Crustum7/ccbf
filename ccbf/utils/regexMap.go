package utils

import (
	"regexp"
)

// RegexMap is a simple map wrapper that makes it easy to find the
// longest match and its corresponding value
type RegexMap[T any] struct {
	m map[string]T
}

func InitRegexMap[T any](m map[string]T) RegexMap[T] {
	reMap := RegexMap[T]{m: make(map[string]T)}
	for key, val := range m {
		regexp.MustCompile(key)

		reMap.m[key] = val
	}
	return reMap
}

func (reMap RegexMap[T]) FindLongestMatchPattern(str string) string {
	longest, bestRegex := "", ""

	for key := range reMap.m {
		re := regexp.MustCompile(key)
		match := re.FindString(str)

		shorterRegex := len(match) == len(longest) && len(key) < len(bestRegex)
		longerMatch := len(match) > len(longest)
		if longerMatch || shorterRegex {
			longest = match
			bestRegex = key
		}
	}

	return bestRegex
}

func (reMap RegexMap[T]) GetValueFromPattern(pattern string) *T {
	value, wasRetrieved := reMap.m[pattern]
	if !wasRetrieved {
		return nil
	}

	return &value
}
