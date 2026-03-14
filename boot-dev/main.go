package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, nonoWord := range badWords {
			if word == nonoWord {
				return i
			}
		}
	}
	return -1
}

