package whatsappChatAnalyser

import (
	"bufio"
	"io"
	"regexp"
	"strings"
	"time"
)

const msgLineWithSenderLayout = `(.*)-\s(\w+):\s(.*)`

func ParseChat(in io.Reader) (chat Chat) {
	reader := bufio.NewReader(in)

	for {
		line, _ := reader.ReadBytes(byte('\n'))
		if len(line) == 0 {
			break
		}
		if match, _ := regexp.MatchString(msgLineWithSenderLayout, string(line)); match {
			sender, content, timestamp := parseFirstLine(line)
			chat = append(chat, Message{sender, content, timestamp})
		} else {
			chat.last().addContent(parseContentOnlyLine(line))
		}
	}

	return
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
