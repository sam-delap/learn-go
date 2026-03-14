package main

import "unicode/utf8"

func getNameCounts(names []string) map[rune]map[string]int {
	nameCountMap := make(map[rune]map[string]int)
	for _, name := range names {
		firstRune, _ := utf8.DecodeRuneInString(name)
		if _, ok := nameCountMap[firstRune]; !ok {
			nameCountMap[firstRune] = make(map[string]int)
			nameCountMap[firstRune][name] = 1
		} else if val, ok := nameCountMap[firstRune][name]; !ok {
			nameCountMap[firstRune][name] = 1
		} else {
			nameCountMap[firstRune][name] = val + 1
		}
	}
	
	return nameCountMap
}
