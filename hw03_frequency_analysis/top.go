package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordsCount struct {
	Key   string
	Value int
}

var regularExpression = regexp.MustCompile(`[\p{L}\d]*[-]*[$\p{L}\d]+`)

func Top10(inputString string) []string {
	input := strings.ToLower(inputString)
	if len(input) == 0 {
		return make([]string, 0)
	}
	matches := regularExpression.FindAllString(input, -1)

	counts := make(map[string]int)
	for _, v := range matches {
		counts[v]++
	}

	countsStruct := make([]wordsCount, 0)
	for i, v := range counts {
		countsStruct = append(countsStruct, wordsCount{i, v})
	}
	sort.SliceStable(countsStruct, func(i, j int) bool {
		if countsStruct[i].Value == countsStruct[j].Value {
			return countsStruct[i].Key < countsStruct[j].Key
		}
		return countsStruct[i].Value > countsStruct[j].Value
	})
	words := make([]string, 0)
	for _, v := range countsStruct[:10] {
		words = append(words, v.Key)
	}
	return words
}
