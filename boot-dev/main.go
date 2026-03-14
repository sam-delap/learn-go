package main

import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, msg := range messages {
		messages[i].tags = tagger(msg)
	}

	return messages
}

func tagger(msg sms) []string {
	tags := []string{}

	lowerMsg := strings.ToLower(msg.content)

	if strings.Contains(lowerMsg, "urgent") {
		tags = append(tags, "Urgent")
	}

	if strings.Contains(lowerMsg, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

