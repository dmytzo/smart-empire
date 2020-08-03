package main

import (
	"smart_empire/mqtt_client"
	"smart_empire/tg_bot"
)

func main() {
	mqtt_client.Run()
	tg_bot.Run()
}
