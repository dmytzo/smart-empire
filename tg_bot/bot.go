package tg_bot

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

const (
	Start         = "start"
	Stop          = "stop"
	TechInfo      = "tech_info"
	Status        = "status"
	SwitchLight   = "switch_light"
	SwitchDefMode = "switch_def_mode"
	SwitchSiren   = "switch_siren"
	DefModeHome   = "def_mode_home"
	DefModeOn     = "def_mode_on"
	DefModeOff    = "def_mode_off"
)

var generalKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", Status))),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", SwitchLight))),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", SwitchDefMode))),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", SwitchSiren))),
)

var defModeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", DefModeHome))),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", DefModeOn))),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", DefModeOff))),
)

func getGeneralKeyboard() tgbotapi.ReplyKeyboardMarkup {
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
			case Start:
				c.Start()
			case Stop:
				c.Stop()
			case Status:
				c.Status()
			case SwitchSiren:
				c.SwitchSiren()
			case SwitchLight:
				c.SwitchLight()
			case SwitchDefMode:
				c.SwitchDefMode()
			case DefModeHome:
				c.DefMode("HOME")
			case DefModeOn:
				c.DefMode("ON")
			case DefModeOff:
				c.DefMode("OFF")
			case TechInfo:
				c.TechInfo()
			}
		}
	}
}
