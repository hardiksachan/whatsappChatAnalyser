package analyzer

import "github.com/hardiksachan/whatsappChatAnalyser/types"

type Analyzer struct {
	chat types.Chat
}

func NewAnalyzer(chat types.Chat) Analyzer {
	return Analyzer{chat}
}

func (a Analyzer) CountMessagesOf(sender string) (count int) {
	for _, message := range a.chat {
		if message.Sender == sender {
			count++
		}
	}
	return
}

func (a Analyzer) ListAllSenders() (senders []string) {
	senderSet := make(map[string]bool)
	for _, message := range a.chat {
		if _, ok := senderSet[message.Sender]; !ok {
			senderSet[message.Sender] = true
			senders = append(senders, message.Sender)
		}
	}
	return
}
