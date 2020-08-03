package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"smart_empire/mqtt_client/sensors"
)

func SensorsHandler(bot *tgbotapi.BotAPI, chatId int64) {
	for {
		select {
		case ds := <- sensors.DoorSensorChan:
			msg := "Closed"
			if ds.Contact == false {
				msg = "Opened"
			}
			bot.Send(tgbotapi.NewMessage(chatId, msg))
		}
	}
}
