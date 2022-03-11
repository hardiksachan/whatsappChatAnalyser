package whatsappChatAnalyser_test

import (
	parser "github.com/hardiksachan/whatsappChatAnalyser"
	"reflect"
	"strings"
	"testing"
	"time"
)

type testCase struct {
	name         string
	data         string
	expectedChat parser.Chat
}

var emptyChat parser.Chat

func TestParseChat(t *testing.T) {
	cases := []testCase{
		{
			name: "parse single line message single message 1",
			data: "2/24/22, 02:04 - Nicole: call me back!",
			expectedChat: parser.Chat{
				parser.Message{
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
			expectedChat: parser.Chat{
				parser.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name: "should ignore system message in the end",
			data: `2/24/22, 02:04 - Nicole: call me back!
2/24/22, 02:04 - Messages and calls are end-to-end encrypted. No one outside of this chat, not even WhatsApp, can read or listen to them. Tap to learn more`,
			expectedChat: parser.Chat{
				parser.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name: "should ignore system message in the middle",
			data: `2/24/22, 02:04 - Nicole: call me back!
2/24/22, 02:04 - Messages and calls are end-to-end encrypted. No one outside of this chat, not even WhatsApp, can read or listen to them. Tap to learn more
2/24/22, 02:04 - Nicole: call me back!`,
			expectedChat: parser.Chat{
				parser.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
				parser.Message{
					Sender: "Nicole", Content: "call me back!", Timestamp: simpleDate(2022, 2, 24, 2, 4),
				},
			},
		},
		{
			name: "should not ignore multiline message containing `-`",
			data: `2/24/22, 02:04 - Nicole: call me back!
this includes a - as well!`,
			expectedChat: parser.Chat{
				parser.Message{
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
			expectedChat: parser.Chat{
				parser.Message{
					Sender: "Chris", Content: "alright!", Timestamp: simpleDate(2022, 01, 30, 2, 45)},
			},
		},
		{
			name: "parse two line single message",
			data: `2/24/22, 02:04 - Nicole: call me back!
I really need to talk to you`,
			expectedChat: parser.Chat{
				parser.Message{
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
			expectedChat: parser.Chat{
				parser.Message{
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
			expectedChat: parser.Chat{
				parser.Message{
					Sender: "Nicole",
					Content: `call me back!
I really need to talk to you`,
					Timestamp: simpleDate(2022, 02, 24, 2, 4),
				},
				parser.Message{
					Sender:    "Chris",
					Content:   "alright!",
					Timestamp: simpleDate(2022, 1, 30, 2, 45),
				},
				parser.Message{
					Sender:    "Nicole",
					Content:   "awesome!",
					Timestamp: simpleDate(2022, 1, 30, 2, 52),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := parser.ParseChat(strings.NewReader(test.data))
			want := test.expectedChat

			assertChat(t, got, want)
		})
	}

}

func assertChat(t testing.TB, got parser.Chat, want parser.Chat) {
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
