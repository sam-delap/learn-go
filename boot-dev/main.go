package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func analyzeMessage(analysis *Analytics, msg Message) {
	analysis.MessagesTotal += 1
	if msg.Success {
		analysis.MessagesSucceeded +=1
	} else {
		analysis.MessagesFailed += 1
	}
}

