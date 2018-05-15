package handlers

import (
	"strings"

	"github.com/nathenapse/telegram-commander/command"
	"github.com/nathenapse/telegram-commander/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type commands interface {
	Exec(*tb.Bot, *tb.Message, *config.Config)
	IsInline() bool
	GetExample() string
	Validate(m *tb.Message) bool
}

// Handlers all command handlers
func Handlers(b *tb.Bot, conf *config.Config) {
	// mutex := new(sync.Mutex)
	// add new commands here
	payloadCommands := map[string]commands{
		"/download":     &command.Downloader{},
		"/listcommands": &command.ListCommands{},
		"/file":         &command.Download{},
		"/help":         &command.Help{},
	}

	b.Handle("/start", func(m *tb.Message) {
		if ValidateUser(b, m, conf.Admin) {
			go b.Send(m.Sender, "Welcome "+m.Sender.FirstName+" "+m.Sender.LastName)
		}
	})

	for command, inter := range payloadCommands {
		go func(command string, inter commands) {
			attachHandlers(b, command, inter, conf)
		}(command, inter)
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		if ValidateUser(b, m, conf.Admin) {
			go func() {
				command := &command.Command{}
				command.Exec(b, m, conf)
			}()
		}
	})

	b.Handle(tb.OnEdited, func(m *tb.Message) {
		isCommand := false
		if ValidateUser(b, m, conf.Admin) {
			for command, inter := range payloadCommands {
				if strings.HasPrefix(m.Text, command) {
					isCommand = true
					m.Payload = strings.TrimLeft(m.Text, command+" ")
					if inter.Validate(m) {
						go inter.Exec(b, m, conf)
					} else {
						go b.Reply(m, inter.GetExample())
					}
				}
			}
			if isCommand == false {
				go func() {
					command := &command.Command{}
					command.Exec(b, m, conf)
				}()
			}
		}
	})

	// Handle Document upload
	b.Handle(tb.OnDocument, func(m *tb.Message) {
		if ValidateUser(b, m, conf.Admin) {
			go func() {
				command := &command.Upload{}
				command.Exec(b, m, conf)
			}()
		}
	})

}

// ValidateUser if message is sent from the user
func ValidateUser(b *tb.Bot, m *tb.Message, user int) bool {
	if m.Sender.ID == user {
		return true
	}
	b.Reply(m, "You do not have permission to access this bot.")
	return false
}

func attachHandlers(b *tb.Bot, command string, inter commands, conf *config.Config) {

	b.Handle(command, func(m *tb.Message) {
		if ValidateUser(b, m, conf.Admin) {
			if inter.Validate(m) {
				go inter.Exec(b, m, conf)
			} else {
				go b.Reply(m, inter.GetExample())
			}
		}
	})
	return
}
