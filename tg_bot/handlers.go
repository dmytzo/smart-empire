package tg_bot

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"smart_empire/sensors"
)

func SensorsHandler(bot *tgbotapi.BotAPI) {
	for {
		select {
		case ds := <- sensors.DoorSensor.MsgChan:
			doorSensorHandler(bot, ds)
		}
	}
}

func doorSensorHandler(bot *tgbotapi.BotAPI, dsMsg sensors.DoorSensorMsg) {
	event := "door is closed"
	if dsMsg.Contact == false {
		event = "door is opened"
	}
	msg := fmt.Sprintf("*DoorSensor:* %s", event)
	for _, user := range GetAuth().getActiveUsers() {
		msg := tgbotapi.NewMessage(user.ChatId, msg)
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}
