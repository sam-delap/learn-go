package main

func maxMessages(thresh int) int {
	i := 0
	totalCost := 100
	for ; totalCost <= thresh; i++ {
		totalCost += 100 + i
	}

	return i
}
