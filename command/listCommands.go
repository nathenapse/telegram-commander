package command

import (
	"github.com/nathenapse/telegram-commander/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

// ListCommands run on help
type ListCommands struct {
	*Base
}

// Setup where u set the commands
func (l *ListCommands) Setup() (commands []string, descriptions []string) {
	commands = []string{
		"start",
		"file",
		"download",
		"listcommands",
		"help",
	}
	descriptions = []string{
		"Prints a welcome messge",
		"Sends the file found on the pc to telegram /file path/to/file",
		"Starts downloading the url with the set Downloader /download http://url",
		"Prints a string that u can use to setcommand on @BotFather",
		"Prints the help message",
	}
	return commands, descriptions
}

// Exec to locate file and send
func (l *ListCommands) Exec(b *tb.Bot, m *tb.Message, conf *config.Config) {
	commands, descriptions := l.Setup()
	message := ""

	for i := 0; i < len(commands); i++ {
		message = message + commands[i] + " - " + descriptions[i] + "\n"
	}

	go b.Reply(m, message)
}

// IsInline return if command is /command somethingelse
func (l *ListCommands) IsInline() bool {
	return false
}

// GetExample return the how to use of this command
func (l *ListCommands) GetExample() string {
	return ""
}

// Validate Default implementation of Validate
func (l *ListCommands) Validate(m *tb.Message) bool {
	return true
}
