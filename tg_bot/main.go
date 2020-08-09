package tg_bot

import (
	"bytes"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os/exec"
	"smart_empire/config"
	"strings"
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
			case "status":
				c.Status()
			case "attack":
				c.Attack()
			case "tech_info":
				c.TechInfo()
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
			tgbotapi.NewKeyboardButton("/status"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/attack"),
		),
	)
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, "Welcome to the SmartEmpire1920 System")
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

func (c Command) Status() {
	status := fmt.Sprintf("%s\n ---------- \n%s", TemperatureSensorCurrentState, DoorSensorCurrentState)
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, status)
	msg.ParseMode = "Markdown"
	c.bot.Send(msg)
}

func (c Command) Attack() {
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, "ATTACK!!!")
	c.bot.Send(msg)
}

func (c Command) TechInfo() {
	cmd := exec.Command("vcgencmd", "measure_temp")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf(err.Error())
	}
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, out.String())
	c.bot.Send(msg)
}
