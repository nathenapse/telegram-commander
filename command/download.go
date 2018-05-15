package command

import (
	"os"

	"github.com/nathenapse/telegram-commander/config"
	"github.com/nathenapse/telegram-commander/custom"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Download to download files
type Download struct {
	*Base
}

// Exec to locate file and send
func (d *Download) Exec(b *tb.Bot, m *tb.Message, conf *config.Config) {

	path := conf.User.HomeDir + "/" + m.Payload
	_, err := os.Stat(path)

	if err == nil {
		file := &tb.Document{File: tb.FromDisk(path), Caption: path}
		b.Send(m.Sender, file, &tb.SendOptions{
			ReplyTo: m,
		})
	} else {
		custom.SendLong(b)(m.Sender, err.Error(), &tb.SendOptions{
			ReplyTo: m,
		})
	}

}

// IsInline return if command is /command somethingelse
func (d *Download) IsInline() bool {
	return true
}

// GetExample return the how to use of this command
func (d *Download) GetExample() string {
	return "please use /file relative/path/to/home"
}
