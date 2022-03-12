package whatsappChatAnalyser_test

import (
	"fmt"
	analyser "github.com/hardiksachan/whatsappChatAnalyser"
	"reflect"
	"testing"
)

func TestListAllSenders(t *testing.T) {
	cases := []struct {
		chat           analyser.Chat
		expectedResult []string
	}{
		{analyser.Chat{
			analyser.Message{Sender: "Nicole"},
		}, []string{"Nicole"}},
		{analyser.Chat{
			analyser.Message{Sender: "Nicole"},
			analyser.Message{Sender: "Nick"},
		}, []string{"Nicole", "Nick"}},
		{analyser.Chat{
			analyser.Message{Sender: "Nicole"},
			analyser.Message{Sender: "Nick"},
			analyser.Message{Sender: "Nick"},
		}, []string{"Nicole", "Nick"}},
	}

	for _, test := range cases {
		t.Run(
			fmt.Sprintf("list %d unique senders in %d messages", len(test.expectedResult), len(test.chat)),
			func(t *testing.T) {
				got := analyser.ListAllSenders(test.chat)

				if !reflect.DeepEqual(got, test.expectedResult) {
					t.Errorf(`
want
	%#v
got
	%#v`, test.expectedResult, got)
				}
			},
		)
	}
}

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
