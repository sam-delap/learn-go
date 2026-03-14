package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, username := range messagedUsers {
		val, ok := validUsers[username]; if ok {
			validUsers[username] = val + 1
		}
	}
}

