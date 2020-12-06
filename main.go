package main

import (
	"smart_empire/gateway"
	"smart_empire/tg_bot"
)

func main() {
	gateway.Run()
	tg_bot.Run()
}
