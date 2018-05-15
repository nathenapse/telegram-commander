package command

import (
	"bytes"
	"os/exec"

	"github.com/nathenapse/telegram-commander/config"
	"github.com/nathenapse/telegram-commander/custom"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Command the basic command
type Command struct {
	*Base
}

// Exec executes the command
func (c *Command) Exec(b *tb.Bot, m *tb.Message, conf *config.Config) {

	cmd := exec.Command("sh", "-c", m.Text)

	var out bytes.Buffer
	var errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb
	cmd = SetUser(cmd, conf)

	err := cmd.Run()

	if err != nil {
		custom.SendLong(b)(m.Sender, err.Error()+":"+errb.String(), &tb.SendOptions{
			ReplyTo: m,
		})
	} else {
		custom.SendLong(b)(m.Sender, out.String(), &tb.SendOptions{
			ReplyTo: m,
		})
	}

}

// IsInline return if command is /command somethingelse
func (c *Command) IsInline() bool {
	return false
}

// GetExample return the how to use of this command
func (c *Command) GetExample() string {
	return "Just use it inline"
}
