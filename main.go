package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/gloob/irc_bot/lib"
	"github.com/tucnak/telebot"
)

var (
	// Global configuration object.
	globalConfig irc_bot.GlobalConfig

	help *bool   = flag.Bool("help", false, "Display usage information")
	host *string = flag.String("host", "irc.irc-hispano.org", "The host to connect to")
	port *int    = flag.Int("port", 6667, "The port to connect to")
)

func PrintUsage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	// Parse commandline options
	flag.Parse()
	if *help {
		PrintUsage()
		return
	}

	ircConfig := irc_bot.ProxyConfig{
		Host:     *host,
		Port:     *port,
		Password: "",
		Nick:     "irc_bot",
		RealName: "The Telegram Irc Bot.",
	}

	proxy, err := irc_bot.Connect(ircConfig)
	if err != nil {
		log.Fatal(err)
	}

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

	go proxy.Run()

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
