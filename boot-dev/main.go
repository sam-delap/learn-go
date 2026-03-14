package main

func maxMessages(thresh int) int {
	totalCost := 0
	if totalCost >= thresh {
		return 0
	}

	i := 0

	for ; totalCost <= thresh; i++ {
		totalCost += 100 + i
	}

	return i - 1
}
