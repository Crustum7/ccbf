package utils

import "regexp"

type RegexMap[T any] struct {
	m map[*regexp.Regexp]T
}

func InitRegexMap[T any](m map[string]T) (RegexMap[T], error) {
	reMap := RegexMap[T]{m: make(map[*regexp.Regexp]T)}
	for key, val := range m {
		regex, err := regexp.Compile(key)
		if err != nil {
			return RegexMap[T]{}, err
		}

		reMap.m[regex] = val
	}
	return reMap, nil
}

func (reMap RegexMap[T]) FindLongestMatch(str string) (string, *T) {
	longest := []byte{}
	var bestVal *T = nil
	for key, val := range reMap.m {
		match := key.Find([]byte(str))
		if match != nil && len(match) > len(longest) {
			longest = match
			bestVal = &val
		}
	}
	return string(longest), bestVal
}
