package whatsappChatAnalyser_test

import (
	analyser "github.com/hardiksachan/whatsappChatAnalyser"
	"testing"
)

func TestCountMessageOfPerson(t *testing.T) {
	cases := []struct {
		name           string
		sender         string
		chat           analyser.Chat
		expectedResult int
	}{
		{"count a single message", "Nicole", analyser.Chat{
			analyser.Message{Sender: "Nicole", Content: "doesn't matter"},
		}, 1},
		{"find count of chat of same sender", "Nicole", analyser.Chat{
			analyser.Message{Sender: "Nicole", Content: "doesn't matter"},
			analyser.Message{Sender: "Nicole", Content: "doesn't matter"},
			analyser.Message{Sender: "Nicole", Content: "doesn't matter"},
		}, 3},
		{"find count of chat of same sender", "Nicole", analyser.Chat{
			analyser.Message{Sender: "Nicole", Content: "doesn't matter"},
			analyser.Message{Sender: "Derek", Content: "doesn't matter"},
			analyser.Message{Sender: "Nicole", Content: "doesn't matter"},
		}, 2},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := analyser.CountMessagesOf(test.sender, test.chat)

			if test.expectedResult != got {
				t.Errorf("want %d, got %d", test.expectedResult, got)
			}
		})
	}
}
