package command

import (
	"bytes"
	"os/exec"

	"github.com/nathenapse/telegram-commander/config"
	"github.com/nathenapse/telegram-commander/custom"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Downloader to download files
type Downloader struct {
	*Base
}

// Exec to locate file and send
func (d *Downloader) Exec(b *tb.Bot, m *tb.Message, conf *config.Config) {

	go b.Reply(m, conf.Downloader+" "+m.Payload)

	// TODO: confirm to overwrite
	cmd := exec.Command("sh", "-c", conf.Downloader+" "+m.Payload)

	var out bytes.Buffer
	var errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb

	cmd = SetUser(cmd, conf)

	go b.Reply(m, "Download, starting")
	err := cmd.Run()

	if err != nil && errb.String() != "" {
		go b.Send(m.Sender, "Download, Stoped with error", &tb.SendOptions{
			ReplyTo: m,
		})
		go custom.SendLong(b)(m.Sender, err.Error()+":"+errb.String(), &tb.SendOptions{
			ReplyTo: m,
		})
	} else {
		go b.Send(m.Sender, "Download, Finished", &tb.SendOptions{
			ReplyTo: m,
		})
		go custom.SendLong(b)(m.Sender, out.String(), &tb.SendOptions{
			ReplyTo: m,
		})
	}

}

// IsInline return if command is /command somethingelse
func (d *Downloader) IsInline() bool {
	return true
}

// GetExample return the how to use of this command
func (d *Downloader) GetExample() string {
	return "please use /download url"
}
