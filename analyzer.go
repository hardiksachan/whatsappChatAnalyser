package whatsappChatAnalyser

func CountMessagesOf(sender string, chat Chat) (count int) {
	for _, message := range chat {
		if message.Sender == sender {
			count++
		}
	}
	return
}
