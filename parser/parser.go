package parser

import (
	"bufio"
	"github.com/hardiksachan/whatsappChatAnalyser/types"
	"io"
	"regexp"
	"strings"
	"time"
)

const (
	msgLineWithSenderLayout = `(.*)-\s(\w+):\s(.*)`
	systemMessagesLayout    = `(.*)-\s+(.*)Tap to\s(.*)`
)

func ParseChat(in io.Reader) (chat types.Chat) {
	reader := bufio.NewReader(in)

	for {
		line, _ := reader.ReadBytes(byte('\n'))
		if len(line) == 0 {
			break
		}
		if isMessageWithSenderTag(line) {
			sender, content, timestamp := parseFirstLine(line)
			chat = append(chat, types.Message{sender, content, timestamp})
			continue
		}
		if isSystemMessage(line) || len(chat) == 0 {
			continue
		}
		chat[len(chat)-1].AddContent(parseContentOnlyLine(line))
	}

	return
}

func isSystemMessage(line []byte) bool {
	return regexp.MustCompile(systemMessagesLayout).MatchString(string(line))
}

func isMessageWithSenderTag(line []byte) bool {
	return regexp.MustCompile(msgLineWithSenderLayout).MatchString(string(line))
}

func parseContentOnlyLine(line []byte) string {
	content := string(line)
	if content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}
	return content
}

func parseFirstLine(in []byte) (sender string, content string, timestamp time.Time) {
	stdMessageRe := regexp.MustCompile(msgLineWithSenderLayout)
	data := stdMessageRe.FindAllSubmatch(in, 1)[0]
	sender = string(data[2])
	timestamp, _ = parseTimestamp(data[1])
	content = string(data[3])
	return
}

func parseTimestamp(unparsed []byte) (time.Time, error) {
	return time.Parse("1/2/06, 15:04", strings.TrimSpace(string(unparsed)))
}
