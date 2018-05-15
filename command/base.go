package command

import (
	"os/exec"
	"os/user"
	"syscall"

	"github.com/nathenapse/telegram-commander/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Base is where all structs inherit
type Base struct {
}

// Validate Default implementation of Validate
func (b *Base) Validate(m *tb.Message) bool {
	return m.Payload != ""
}

// SetUser is to set the user for the command
func SetUser(cmd *exec.Cmd, conf *config.Config) *exec.Cmd {
	// TODO: handle error
	usr, _ := user.Current()

	if usr.Uid == "0" {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: conf.Userid, Gid: conf.Groupid}
		cmd.Dir = conf.User.HomeDir
	} else {
		cmd.Dir = usr.HomeDir
	}

	return cmd
}
