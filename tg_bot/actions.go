package tg_bot

import (
	"smart_empire/gateway"
	"smart_empire/gateway/rules"
)

func getCurrentTemperatureMsg() string {
	return getTemperatureEventMessage(gateway.TemperatureSensor.GetLatestMsg())
}

func getCurrentDoorMsg() string {
	return getDoorEventMessage(gateway.DoorSensor.GetLatestMsg())
}

func getCurrentLightMsg() string {
	return getLightEventMessage(gateway.Light.GetLatestMsg())
}

func getCurrentSirenMsg() string {
	return getSirenEventMessage(gateway.Siren.GetLatestMsg())
}

func getCurrentDefModeMsg() string {
	return getDefModeMessage(rules.GetDefMode())
}

func switchLight() string {
	gateway.Light.Switch()
	return getCurrentLightMsg()
}

func switchSiren() string {
	gateway.Siren.AutoSwitch()
	return getCurrentSirenMsg()
}

func switchDefMode(mode string) string {
	rules.SetDefMode(mode)
	return getDefModeMessage(mode)
}
