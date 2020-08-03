package tg_bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"smart_empire/config"
)

var cfg = config.Cfg.TgBot

func setUpTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(cfg.ApiToken)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func Run() {
	bot := setUpTelegramBot()
	log.Printf(bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	go SensorsHandler(bot)

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
			c := Command{bot: bot, update: update}
			switch command {
			case "start":
				c.Start()
			case "stop":
				c.Stop()
			}
		}
	}
}

type Command struct {
	bot *tgbotapi.BotAPI
	update tgbotapi.Update
}

func (c Command) Start() {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("EmpireStatus"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("EmpireAttack"),
		),
	)
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, "Welcome to SmartEmpire1920 System")
	msg.ReplyMarkup = keyboard

	GetAuth().updateUserStatus(c.update.Message.From.UserName, true)
	c.bot.Send(msg)
}

func (c Command) Stop() {
	GetAuth().updateUserStatus(c.update.Message.From.UserName, false)
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, "You are logged out from the SmartEmpire1920 System")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	c.bot.Send(msg)
}
