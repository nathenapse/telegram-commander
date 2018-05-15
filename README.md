# telegram-commander
Self deployed telegram bot that will help with running simple commands



- [DESCRIPTION](#description)
- [INSTALLATION](#installation)
- [DEVELOPMENT](#development)
- [TODO](#todo)
- [LIMITATIONS](#limitations)
- [LICENSE](#license)

## DESCRIPTION

**telegram-commander** is a self deployed telegram bot that let's you run commands on a remote pc through [telegram](https://telegram.org). There are default commands that will help you run simple commands like Upload a file to pc, Download a file from pc, Download file from internet and Run simple commands and more to come.

## REQUIREMENTS

- GoLang > 1.9.1
- One of terminal file downloaders wget, curl, axel


## INSTALLATION

- Download the latest release binary zip
- unzip the file
- inside there is a .env.example file
- To install it first clone the project on your pc 
- Create a telegram bot using [@BotFather](https://t.me/botfather)
	- Set the name of the bot
	- Set the username of the bot
- Get your telegram user id using [@IDBot](https://t.me/myidbot)
- Copy .env.example to .env
```bash
cp .env.example .env
```
- Set your bot and telegram user id on .env
```YAML
TELEGRAM_TOKEN=<Telegram bot token at @BotFather>
TELEGRAM_USER_ID=<Telegram user id @getidsbot>
# DEFAULT_DOWNLOADER put the full command 
# example 
# DEFAULT_DOWNLOADER=curl -O
DEFAULT_DOWNLOADER=wget --content-disposition
# Username of the user u want to run command on as
RUNAS=<username of the pc>
```
- Run the executable 
```bash 
./telegram-commander
```
- If you want to run it as a service 
- First install the service
```bash 
sudo ./telegram-commander -service install
```
- you can use your systems service manager to handle the service
```bash
sudo service telegram-commander start
sudo service telegram-commander stop
```

## DEVELOPMENT
- get the package
```bash
go get github.com/nathenapse/telegram-commander
```

## TODO

- [ ] Torrent Downloads
- [ ] Youtube Downloads
- [ ] Better Documentation

## LIMITATIONS

- Works on linux and mac

## LICENSE

MIT