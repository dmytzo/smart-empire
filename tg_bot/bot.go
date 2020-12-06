package tg_bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

var generalKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/status")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/switch_light")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/switch_def_mode")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/switch_siren")),
)

var defModeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/def_mode_home")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/def_mode_on")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/def_mode_off")),
)

func getGeneralKeyboard() tgbotapi.ReplyKeyboardMarkup{
	return generalKeyboard
}

func createTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(auth.getApiToken())
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func RunTelegramBot() {
	bot := createTelegramBot()
	log.Printf(bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	go EventsHandler(bot)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		_, err := auth.getUserByUsername(update.Message.From.UserName)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
		}
		if update.Message.IsCommand() {
			command := update.Message.Command()
			c := Command{bot, update}
			switch command {
			case "start":
				c.Start()
			case "stop":
				c.Stop()
			case "status":
				c.Status()
			case "switch_siren":
				c.SwitchSiren()
			case "switch_light":
				c.SwitchLight()
			case "switch_def_mode":
				c.SwitchDefMode()
			case "def_mode_home":
				c.DefMode("HOME")
			case "def_mode_on":
				c.DefMode("ON")
			case "def_mode_off":
				c.DefMode("OFF")
			case "tech_info":
				c.TechInfo()
			}
		}
	}
}
