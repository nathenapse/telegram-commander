package custom

import (
	"bytes"

	tb "gopkg.in/tucnak/telebot.v2"
)

// SendLong is used to send long messages
func SendLong(b *tb.Bot) func(to tb.Recipient, what interface{}, options ...interface{}) ([]*tb.Message, []error) {
	return func(to tb.Recipient, what interface{}, options ...interface{}) ([]*tb.Message, []error) {
		messages := []*tb.Message{}
		errors := []error{}
		switch object := what.(type) {
		case string:
			strings := splitString(object, 4000)
			for _, str := range strings {
				message, error := b.Send(to, str, options...)
				messages = append(messages, message)
				errors = append(errors, error)
			}
			return messages, errors
		default:
			return messages, errors
		}
	}
}

func splitString(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
