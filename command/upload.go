package command

import (
	"github.com/nathenapse/telegram-commander/config"
	"github.com/nathenapse/telegram-commander/custom"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Upload to download files
type Upload struct {
	*Base
}

// Exec to locate file and send
func (u *Upload) Exec(b *tb.Bot, m *tb.Message, conf *config.Config) {

	path := conf.User.HomeDir + "/Downloads/" + m.Document.FileName
	// TODO: file already exists do u want to replace it

	if err := b.Download(&m.Document.File, path); err != nil {
		custom.SendLong(b)(m.Sender, err.Error(), &tb.SendOptions{
			ReplyTo: m,
		})
		return
	}
	b.Send(m.Sender, "Saved Successfully to "+path, &tb.SendOptions{
		ReplyTo: m,
	})

}

// IsInline return if command is /command somethingelse
func (u *Upload) IsInline() bool {
	return false
}

// GetExample return the how to use of this command
func (u *Upload) GetExample() string {
	return "send documents"
}
