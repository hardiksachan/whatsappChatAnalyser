package analyzer

import "github.com/hardiksachan/whatsappChatAnalyser/types"

func CountMessagesOf(sender string, chat types.Chat) (count int) {
	for _, message := range chat {
		if message.Sender == sender {
			count++
		}
	}
	return
}

func ListAllSenders(chat types.Chat) (senders []string) {
	senderSet := make(map[string]bool)
	for _, message := range chat {
		if _, ok := senderSet[message.Sender]; !ok {
			senderSet[message.Sender] = true
			senders = append(senders, message.Sender)
		}
	}
	return
}
