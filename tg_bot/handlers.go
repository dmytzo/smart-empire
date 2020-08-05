package tg_bot

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"smart_empire/sensors"
)

var TemperatureSensorCurrentState string
var DoorSensorCurrentState string

func SensorsHandler(bot *tgbotapi.BotAPI) {
	for {
		select {
		case ds := <- sensors.DoorSensor.MsgChan:
			doorSensorHandler(bot, ds)
		case ts := <- sensors.TemperatureSensor.MsgChan:
			thermometerSensorHandler(bot, ts)
		}
	}
}

func doorSensorHandler(bot *tgbotapi.BotAPI, dsMsg sensors.DoorSensorMsg) {
	event := "closed"
	if dsMsg.Contact == false {
		event = "opened"
	}
	msg := fmt.Sprintf("*DoorSensor:* \n%s", event)
	DoorSensorCurrentState = msg
	for _, user := range GetAuth().getActiveUsers() {
		msg := tgbotapi.NewMessage(user.ChatId, msg)
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}

func thermometerSensorHandler(bot *tgbotapi.BotAPI, dsMsg sensors.TemperatureSensorMsg) {
	msg := fmt.Sprintf(
		"*TemperatureSensor:* \n" +
			"%.1f ℃ / %.1f %%",
			dsMsg.Temperature, dsMsg.Humidity,
		)
	TemperatureSensorCurrentState = msg
}
