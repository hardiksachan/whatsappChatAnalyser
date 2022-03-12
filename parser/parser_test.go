package parser_test

import (
	analyser "github.com/hardiksachan/whatsappChatAnalyser/parser"
	"github.com/hardiksachan/whatsappChatAnalyser/types"
	"reflect"
	"strings"
	"testing"
	"time"
)

type testCase struct {
	name         string
	data         string
	expectedChat types.Chat
}

var emptyChat types.Chat

func TestParseChat(t *testing.T) {
	cases := []testCase{
		{
			name: "parse single line message single message 1",
			data: "2/24/22, 02:04 - Nicole: call me back!",
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name:         "should ignore system message if no message exist",
			data:         "2/24/22, 02:04 - Messages and calls are end-to-end encrypted. No one outside of this chat, not even WhatsApp, can read or listen to them. Tap to learn more",
			expectedChat: emptyChat,
		},
		{
			name: "should ignore system message in the beginning",
			data: `2/24/22, 02:04 - Messages and calls are end-to-end encrypted. No one outside of this chat, not even WhatsApp, can read or listen to them. Tap to learn more
2/24/22, 02:04 - Nicole: call me back!`,
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name: "should ignore system message in the end",
			data: `2/24/22, 02:04 - Nicole: call me back!
2/24/22, 02:04 - Messages and calls are end-to-end encrypted. No one outside of this chat, not even WhatsApp, can read or listen to them. Tap to learn more`,
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name: "should ignore system message in the middle",
			data: `2/24/22, 02:04 - Nicole: call me back!
2/24/22, 02:04 - Messages and calls are end-to-end encrypted. No one outside of this chat, not even WhatsApp, can read or listen to them. Tap to learn more
2/24/22, 02:04 - Nicole: call me back!`,
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
				types.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name: "should not ignore multiline message containing `-`",
			data: `2/24/22, 02:04 - Nicole: call me back!
this includes a - as well!`,
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole",
					Content: `call me back!
this includes a - as well!`,
					Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name: "parse single line single message 2",
			data: "1/30/22, 02:45 - Chris: alright!",
			expectedChat: types.Chat{
				types.Message{
					Sender: "Chris", Content: "alright!", Timestamp: simpleDate(2022, 01, 30, 2, 45)},
			},
		},
		{
			name: "parse two line single message",
			data: `2/24/22, 02:04 - Nicole: call me back!
I really need to talk to you`,
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole",
					Content: `call me back!
I really need to talk to you`,
					Timestamp: simpleDate(2022, 02, 24, 2, 4),
				},
			},
		},
		{
			name: "parse three line single message",
			data: `2/24/22, 02:04 - Nicole: call me back!
I really need to talk to you
Another Line`,
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole",
					Content: `call me back!
I really need to talk to you
Another Line`,
					Timestamp: simpleDate(2022, 02, 24, 2, 4),
				},
			},
		},
		{
			name: "parse multi line three message chat",
			data: `2/24/22, 02:04 - Nicole: call me back!
I really need to talk to you
1/30/22, 02:45 - Chris: alright!
1/30/22, 02:52 - Nicole: awesome!`,
			expectedChat: types.Chat{
				types.Message{
					Sender: "Nicole",
					Content: `call me back!
I really need to talk to you`,
					Timestamp: simpleDate(2022, 02, 24, 2, 4),
				},
				types.Message{
					Sender:    "Chris",
					Content:   "alright!",
					Timestamp: simpleDate(2022, 1, 30, 2, 45),
				},
				types.Message{
					Sender:    "Nicole",
					Content:   "awesome!",
					Timestamp: simpleDate(2022, 1, 30, 2, 52),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := analyser.ParseChat(strings.NewReader(test.data))
			want := test.expectedChat

			assertChat(t, got, want)
		})
	}

}

func assertChat(t testing.TB, got types.Chat, want types.Chat) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("expected length %d, got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`want:
	"%#v"
got:
	"%#v"`, want, got)
	}
}

func simpleDate(year int, month time.Month, day, hour, min int) time.Time {
	return time.Date(year, month, day, hour, min, 0, 0, time.UTC)
}
