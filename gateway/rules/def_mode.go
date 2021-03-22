package rules

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/devices"
	"time"
)

const (
	ON   = "ON"
	OFF  = "OFF"
	HOME = "HOME"
)

var defMode = OFF

func GetDefMode() string {
	return defMode
}

func SetDefMode(state string) {
	defMode = state
}

func IsHomeDefModeActiveTime() bool {
	now := time.Now().Hour()
	return now <= 7 && now == 23
}

func IsDefModeActivated() bool {
	return defMode == ON
}

func IsHomeDefModeActivated() bool {
	return defMode == HOME
}

func DoorDef(msg mqtt.Message, door devices.DoorSensorType, siren *devices.SirenType) {
	doorMsg := door.ParseMsg(msg)
	if doorMsg.Contact {
		return
	}
	if IsDefModeActivated() || (IsHomeDefModeActivated() && IsHomeDefModeActiveTime()) {
		siren.ActivateAlarm()
	}
}
