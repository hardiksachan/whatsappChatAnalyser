package whatsappChatAnalyser

import (
	"bufio"
	"io"
	"regexp"
	"strings"
	"time"
)

const msgLineWithSenderLayout = `(.*)-\s(\w+):\s(.*)`

type Message struct {
	Sender    string
	content   strings.Builder
	Timestamp time.Time
}

func (m Message) Content() string {
	return m.content.String()
}

func (m *Message) writeContent(line string) {
	if m.content.Len() != 0 {
		m.content.WriteByte(byte('\n'))
	}
	m.content.WriteString(line)
}

type Chat []Message

func ParseChat(in io.Reader) (chat Chat) {
	reader := bufio.NewReader(in)

	for {
		line, _ := reader.ReadBytes(byte('\n'))
		if len(line) == 0 {
			break
		}
		if match, _ := regexp.MatchString(msgLineWithSenderLayout, string(line)); match {
			chat = append(chat, Message{})
			parseFirstLine(line, &chat[len(chat)-1])
		} else {
			readMessageLine(line, &chat[len(chat)-1])
		}
	}

	return
}

func readMessageLine(line []byte, message *Message) {
	content := string(line)
	if content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}
	message.writeContent(content)
}

func parseFirstLine(in []byte, message *Message) {
	stdMessageRe := regexp.MustCompile(msgLineWithSenderLayout)
	data := stdMessageRe.FindAllSubmatch(in, 1)[0]
	message.Sender = string(data[2])
	t, _ := parseTimestamp(data[1])
	message.Timestamp = t
	message.writeContent(string(data[3]))
}

func parseTimestamp(unparsed []byte) (time.Time, error) {
	return time.Parse("1/2/06, 15:04", strings.TrimSpace(string(unparsed)))
}
