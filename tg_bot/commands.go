package tg_bot

import (
	"bytes"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"os/exec"
	"strings"
)

type Command struct {
	bot *tgbotapi.BotAPI
	update tgbotapi.Update
}

func (c Command) Start() {
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, "Welcome to the SmartEmpire1920 System")
	msg.ReplyMarkup = getGeneralKeyboard()
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
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, getStatusMsg())
	msg.ParseMode = "Markdown"
	c.bot.Send(msg)
}

func (c Command) SwitchLight() {
	lightMsg := switchLight()
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, lightMsg)
	msg.ParseMode = "Markdown"
	c.bot.Send(msg)
}

func (c Command) SwitchSiren() {
	sirenMsg := switchSiren()
	sendToActiveUsers(c.bot, sirenMsg)
}

func (c Command) SwitchDefMode() {
	msg := tgbotapi.NewMessage(c.update.Message.Chat.ID, "Select DefMode:")
	msg.ReplyMarkup = defModeKeyboard
	c.bot.Send(msg)
}

func (c Command) DefMode(mode string) {
	defModeMsg := switchDefMode(mode)
	sendToActiveUsers(c.bot, defModeMsg)
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
