package tg_bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"smart_empire/gateway"
	"smart_empire/gateway/devices"
)

func EventsHandler(bot *tgbotapi.BotAPI) {
	for {
		select {
		case ds := <- gateway.DoorSensor.DoorEventChan:
			doorEventsHandler(bot, ds)
		case _ = <- gateway.Siren.AlarmChan:
			alarmEventsHandler(bot)
		}
	}
}

func doorEventsHandler(bot *tgbotapi.BotAPI, msg devices.DoorSensorMsg) {
	textMessage := getDoorEventMessage(msg)
	sendToActiveUsers(bot, textMessage)
}

func alarmEventsHandler(bot *tgbotapi.BotAPI) {
	sendToActiveUsers(bot, getAlarmMessage())
}

func sendToActiveUsers(bot *tgbotapi.BotAPI, msg string) {
	for _, user := range GetAuth().getActiveUsers() {
		msg := tgbotapi.NewMessage(user.ChatId, msg)
		msg.ReplyMarkup = getGeneralKeyboard()
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}
