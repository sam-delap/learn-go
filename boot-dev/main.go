package main

import "strings"

func countDistinctWords(messages []string) int {
	wordCounter := make(map[string]struct{})
	for _, msg := range messages {
		words := strings.Fields(strings.ToLower(msg))
		for _, word := range words {
			_, ok := wordCounter[word]; if !ok {
				wordCounter[word] = struct{}{}
			}
		}
	}

	return len(wordCounter)
}

