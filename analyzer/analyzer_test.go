package analyzer_test

import (
	"fmt"
	"github.com/hardiksachan/whatsappChatAnalyser/analyzer"
	"github.com/hardiksachan/whatsappChatAnalyser/types"
	"reflect"
	"testing"
)

func TestListAllSenders(t *testing.T) {
	cases := []struct {
		chat           types.Chat
		expectedResult []string
	}{
		{types.Chat{
			types.Message{Sender: "Nicole"},
		}, []string{"Nicole"}},
		{types.Chat{
			types.Message{Sender: "Nicole"},
			types.Message{Sender: "Nick"},
		}, []string{"Nicole", "Nick"}},
		{types.Chat{
			types.Message{Sender: "Nicole"},
			types.Message{Sender: "Nick"},
			types.Message{Sender: "Nick"},
		}, []string{"Nicole", "Nick"}},
	}

	for _, test := range cases {
		t.Run(
			fmt.Sprintf("list %d unique senders in %d messages", len(test.expectedResult), len(test.chat)),
			func(t *testing.T) {
				got := analyzer.NewAnalyzer(test.chat).ListAllSenders()

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
		chat           types.Chat
		expectedResult int
	}{
		{"count a single message", "Nicole", types.Chat{
			types.Message{Sender: "Nicole", Content: "doesn't matter"},
		}, 1},
		{"find count of chat of same sender", "Nicole", types.Chat{
			types.Message{Sender: "Nicole", Content: "doesn't matter"},
			types.Message{Sender: "Nicole", Content: "doesn't matter"},
			types.Message{Sender: "Nicole", Content: "doesn't matter"},
		}, 3},
		{"find count of chat of same sender", "Nicole", types.Chat{
			types.Message{Sender: "Nicole", Content: "doesn't matter"},
			types.Message{Sender: "Derek", Content: "doesn't matter"},
			types.Message{Sender: "Nicole", Content: "doesn't matter"},
		}, 2},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := analyzer.NewAnalyzer(test.chat).CountMessagesOf(test.sender)

			if test.expectedResult != got {
				t.Errorf("want %d, got %d", test.expectedResult, got)
			}
		})
	}
}
