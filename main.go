package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
	"smart_empire/mqtt_client"
)

var botApi = os.Getenv("API_TOKEN")
var password = os.Getenv("PASSWORD")


func setUpTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botApi)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func main() {
	bot := setUpTelegramBot()
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			command := update.Message.Command()
			switch command {
			case "start":
				arguments := update.Message.CommandArguments()
				if arguments != password {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong password")
					bot.Send(msg)
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to SmartEmpire1920")
				bot.Send(msg)
				mqtt_client.SetUpClient()
				SensorsHandler(bot, update.Message.Chat.ID)
			case "/stop":
				mqtt_client.Disconnect()
			}
		}
	}
}


