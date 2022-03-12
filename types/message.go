package types

import (
	"strings"
	"time"
)

type Chat []Message

func (c Chat) last() *Message {
	return &c[len(c)-1]
}

type Message struct {
	Sender    string
	Content   string
	Timestamp time.Time
}

func (m *Message) AddContent(content string) {
	contentBuilder := strings.Builder{}
	contentBuilder.WriteString(m.Content)
	contentBuilder.WriteByte(byte('\n'))
	contentBuilder.WriteString(content)
	m.Content = contentBuilder.String()
}
