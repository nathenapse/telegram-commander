package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kardianos/osext"
	"github.com/kardianos/service"
	"github.com/nathenapse/telegram-commander/config"
	"github.com/nathenapse/telegram-commander/handlers"
	tb "gopkg.in/tucnak/telebot.v2"
)

var logger service.Logger

// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit    chan struct{}
	service service.Service
	*config.Config
	cmd *exec.Cmd
}

func (p *program) Start(s service.Service) error {
	// Look for exec.
	// Verify home directory.

	go p.run()
	return nil
}

func (p *program) run() error {
	logger.Info("Starting Telegram Commander")
	defer func() {
		if service.Interactive() {
			p.Stop(p.service)
		} else {
			p.service.Stop()
		}
	}()

	b, err := tb.NewBot(tb.Settings{
		Token:  p.Config.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	handlers.Handlers(b, p.Config)

	b.Start()

	return nil
}
func (p *program) Stop(s service.Service) error {
	close(p.exit)
	logger.Info("Stopping Telegram Commander")
	if p.cmd.ProcessState.Exited() == false {
		p.cmd.Process.Kill()
	}
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}

func main() {
	envpath, err := getConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	godotenv.Load(envpath)
	godotenv.Load()

	conf := &config.Config{}

	err = conf.Set()

	if err != nil {
		log.Fatal(err)
	}

	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        "telegram-commander",
		DisplayName: "Telegram Commander",
		Description: "Control your pc through telegram",
	}

	prg := &program{
		exit: make(chan struct{}),
	}
	prg.Config = conf

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	prg.service = s

	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func getConfigPath() (string, error) {
	fullexecpath, err := osext.Executable()
	if err != nil {
		return "", err
	}

	dir, _ := filepath.Split(fullexecpath)

	return filepath.Join(dir, ".env"), nil
}
