package config

import (
	"errors"
	"os"
	"os/user"
	"strconv"
)

// Config for the app to user got from .env file
type Config struct {
	Token      string
	Admin      int
	Username   string
	User       *user.User
	Userid     uint32
	Groupid    uint32
	Downloader string
}

// Set from .env
func (c *Config) Set() error {
	token, found := os.LookupEnv("TELEGRAM_TOKEN")
	if found == false {
		return errors.New("Telegram Token not set, Please use go Build to build the app")
	}

	s, found := os.LookupEnv("TELEGRAM_USER_ID")
	if found == false {
		return errors.New("Telegram user not set, Please use go Build to build the app")
	}

	admin, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	username, found := os.LookupEnv("RUNAS")
	if found == false {
		return errors.New("Username not set to run the app as, Please use go Build to build the app")
	}

	usr, err := user.Lookup(username)
	if err != nil {
		return errors.New("User not found with set username")
	}

	downloader, found := os.LookupEnv("DEFAULT_DOWNLOADER")
	if found == false {
		return errors.New("Downloader not set on the app not found")
	}

	c.Token = token
	c.Admin = admin
	c.User = usr
	c.Username = username
	c.Downloader = downloader
	c.Userid = converStringToInt32(usr.Uid)
	c.Groupid = converStringToInt32(usr.Gid)

	return nil
}

func converStringToInt32(str string) (result uint32) {

	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		panic(err)
	}
	result = uint32(i)
	return
}
