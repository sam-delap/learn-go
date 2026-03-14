package main

func getMessageCosts(messages []string) []float64 {
	numMessages := len(messages)
	msgCosts := make([]float64, numMessages, numMessages)

	for i := 0; i < numMessages; i++ {
		msgCosts[i] = float64(len(messages[i])) * 0.01
	}

	return msgCosts
}

