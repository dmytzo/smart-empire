package rules

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/devices"
)

func SwitchControl(msg mqtt.Message, sw devices.SwitchType, siren *devices.SirenType, light *devices.LightType) {
	switchType := sw.ParseMsg(msg)
	switch switchType.Action {
	case devices.Btn1Click:
		light.Switch()
	case devices.Btn4Click:
		siren.Switch()
	}
}
