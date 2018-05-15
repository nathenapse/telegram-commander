package command

import (
	"github.com/nathenapse/telegram-commander/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Help run on help
type Help struct {
	*ListCommands
}

// Exec to locate file and send
func (h *Help) Exec(b *tb.Bot, m *tb.Message, conf *config.Config) {
	commands, descriptions := h.Setup()
	message := `This Bot is used to do simple things to your unix based pc from the confort of telegram
	like start downloading a file, run commands, get files, send files .. etc
	
	Commands To use
	`

	for i := 0; i < len(commands); i++ {
		message = message + "/" + commands[i] + " - " + descriptions[i] + "\n"
	}

	go b.Reply(m, message)
}
