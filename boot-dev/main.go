package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgArray := [3]string{primary, secondary, tertiary}
	costArray := [3]int{0, 0, 0}

	msgCost := 0
	for i := 0; i < len(msgArray); i++ {
		msgCost += len(msgArray[i])
		costArray[i] = msgCost
	}

	return msgArray, costArray
}

