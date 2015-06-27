package main

import (
	"os/user"
	"time"

	"github.com/gloob/irc_bot/lib"
	"github.com/tucnak/telebot"
)

var (
	// Global configuration object.
	globalConfig irc_bot.Config
)

func main() {
	// Load main configuration.
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	configFile := usr.HomeDir + "/.config/irc_bot/config.toml"

	irc_bot.LoadConfig(configFile, &globalConfig)

	bot, err := telebot.Create(globalConfig.Token)
	if err != nil {
		return
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if message.Text == "/start" {
			bot.SendMessage(message.Chat, "TODO: Implement start..")
		}
		if message.Text == "/stop" {
			bot.SendMessage(message.Chat, "TODO: Implement stop.")
		}
		if message.Text == "/help" {
			bot.SendMessage(message.Chat, "This is a IRC bot for connecting telegram with your irc server / channel favorite and relay content between them.")
		}
		if message.Text == "/settings" {
			bot.SendMessage(message.Chat, "TODO: Implement settings reply.")
		}
		if message.Text == "/stats" {
			bot.SendMessage(message.Chat, "TODO: Implement stop.")
		}
	}
}
