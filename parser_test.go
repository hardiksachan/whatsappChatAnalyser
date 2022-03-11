package whatsappChatAnalyser_test

import (
	"fmt"
	"github.com/hardiksachan/whatsappChatAnalyser"
	"strings"
	"testing"
	"time"
)

type testableMessage struct {
	in                string
	expectedSender    string
	expectedMessage   string
	expectedTimestamp time.Time
}

func TestParseSingleLineMessage(t *testing.T) {
	cases := []testableMessage{
		{"2/24/22, 02:04 - Nicole: call me back!", "Nicole", "call me back!", simpleDate(2022, 2, 24, 2, 4)},
		{"1/30/22, 02:45 - Chris: alright!", "Chris", "alright!", simpleDate(2022, 01, 30, 2, 45)},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("succesfully extract sender from %#v", c.in), func(t *testing.T) {
			got := whatsappChatAnalyser.ParseChat(strings.NewReader(c.in))[0].Sender
			want := c.expectedSender

			if got != want {
				t.Errorf("want %v, got %v", want, got)
			}
		})

		t.Run(fmt.Sprintf("succesfully extract message from %#v", c.in), func(t *testing.T) {
			got := whatsappChatAnalyser.ParseChat(strings.NewReader(c.in))[0].Content()
			want := c.expectedMessage

			if got != want {
				t.Errorf("want %v, got %v", want, got)
			}
		})

		t.Run(fmt.Sprintf("succesfully extract timestamp from %#v", c.in), func(t *testing.T) {
			got := whatsappChatAnalyser.ParseChat(strings.NewReader(c.in))[0].Timestamp
			want := c.expectedTimestamp

			if !got.Equal(want) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}

func TestParseMultiLineMessage(t *testing.T) {
	cases := []testableMessage{
		{
			`2/24/22, 02:04 - Nicole: call me back!
I really need to talk to you`,
			"Nicole",
			`call me back!
I really need to talk to you`,
			simpleDate(2022, 02, 24, 2, 4),
		},
		{
			`2/24/22, 02:04 - Nicole: call me back!
I really need to talk to you
Another Line`,
			"Nicole",
			`call me back!
I really need to talk to you
Another Line`,
			simpleDate(2022, 02, 24, 2, 4),
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("succesfully extract sender from %#v", c.in), func(t *testing.T) {
			got := whatsappChatAnalyser.ParseChat(strings.NewReader(c.in))[0].Sender
			want := c.expectedSender

			if got != want {
				t.Errorf("want %v, got %v", want, got)
			}
		})

		t.Run(fmt.Sprintf("succesfully extract message from %#v", c.in), func(t *testing.T) {
			got := whatsappChatAnalyser.ParseChat(strings.NewReader(c.in))[0].Content()
			want := c.expectedMessage

			if got != want {
				t.Errorf(`want "%v", got "%v"`, want, got)
			}
		})

		t.Run(fmt.Sprintf("succesfully extract timestamp from %#v", c.in), func(t *testing.T) {
			got := whatsappChatAnalyser.ParseChat(strings.NewReader(c.in))[0].Timestamp
			want := c.expectedTimestamp

			if !got.Equal(want) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}

func TestParseMultipleMessages(t *testing.T) {
	input := `2/24/22, 02:04 - Nicole: call me back!
I really need to talk to you
1/30/22, 02:45 - Chris: alright!
1/30/22, 02:52 - Nicole: awesome!`

	got := whatsappChatAnalyser.ParseChat(strings.NewReader(input))

	if len(got) != 3 {
		t.Fatalf(`did not read all messages, input: "%#v", length recieved "%d"`, input, len(got))
	}
}

func simpleDate(year int, month time.Month, day, hour, min int) time.Time {
	return time.Date(year, month, day, hour, min, 0, 0, time.UTC)
}
