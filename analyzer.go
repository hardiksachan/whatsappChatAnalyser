package whatsappChatAnalyser

func CountMessagesOf(sender string, chat Chat) (count int) {
	for _, message := range chat {
		if message.Sender == sender {
			count++
		}
	}
	return
}

func ListAllSenders(chat Chat) (senders []string) {
	senderSet := make(map[string]bool)
	for _, message := range chat {
		if _, ok := senderSet[message.Sender]; !ok {
			senderSet[message.Sender] = true
			senders = append(senders, message.Sender)
		}
	}
	return
}
