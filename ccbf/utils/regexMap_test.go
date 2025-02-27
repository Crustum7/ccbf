package utils_test

import (
	"fmt"
	"testing"

	"martinjonson.com/ccbf/ccbf/utils"
)

func TestRegexMap(t *testing.T) {
	testcases := []struct {
		m               map[string]int
		str             string
		expectedPattern string
	}{
		{
			m: map[string]int{
				`\+`: 5,
			},
			str:             "+++",
			expectedPattern: `\+`,
		},
		{
			m: map[string]int{
				`-+`: 5,
				`\+`: 3,
			},
			str:             "+++------",
			expectedPattern: "-+",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%v", tc.m), func(t *testing.T) {
			regexMap, _ := utils.InitRegexMap(tc.m)
			match := regexMap.FindLongestMatchPattern(tc.str)

			if match != tc.expectedPattern {
				t.Fatalf(
					"FindLongestMatch should have found match %s but found %s",
					tc.expectedPattern, match,
				)
			}
		})
	}
}

func TestRegexMapGetValue(t *testing.T) {
	testcases := []struct {
		m             map[string]int
		str           string
		expectedValue int
	}{
		{
			m: map[string]int{
				`\+`: 5,
			},
			str:           "+++",
			expectedValue: 5,
		},
		{
			m: map[string]int{
				`-+`: 5,
				`\+`: 3,
			},
			str:           "+++------",
			expectedValue: 5,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%v", tc.m), func(t *testing.T) {
			regexMap, _ := utils.InitRegexMap(tc.m)
			pattern := regexMap.FindLongestMatchPattern(tc.str)
			val := regexMap.GetValueFromPattern(pattern)

			if val == nil {
				t.Fatalf(
					"FindLongestMatch should have found value %d but found nil",
					tc.expectedValue,
				)
			}

			if *val != tc.expectedValue {
				t.Fatalf(
					"FindLongestMatch should have found value %d but found %d",
					tc.expectedValue, *val,
				)
			}
		})
	}
}

func TestRegexMapIncorrectRegex(t *testing.T) {
	testcases := []struct {
		m map[string]int
	}{
		{
			m: map[string]int{
				"[": 5,
			},
		},
		{
			m: map[string]int{
				"-+": 5,
				"[":  3,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%v", tc.m), func(t *testing.T) {
			_, err := utils.InitRegexMap(tc.m)

			if err == nil {
				t.Fatalf("Incorrect regex %v should have returned an error", tc.m)
			}
		})
	}
}
