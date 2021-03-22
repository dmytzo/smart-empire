package gateway

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/client"
	"smart_empire/gateway/config"
	"smart_empire/gateway/devices"
	"smart_empire/gateway/rules"
)

var DoorSensor = devices.DoorSensorType{
	Topic:           config.Devices.Door.Topic,
	LastReceivedMsg: devices.DoorSensorMsg{},
	EventChan:       make(chan devices.DoorSensorMsg),
}

var Light = devices.LightType{
	Topic:           config.Devices.Light.Topic,
	LastReceivedMsg: devices.LightMsg{},
}

var TemperatureSensor = devices.TemperatureSensorType{
	Topic:           config.Devices.Temperature.Topic,
	LastReceivedMsg: devices.TemperatureSensorMsg{},
}

var Siren = devices.SirenType{
	Topic:           config.Devices.Siren.Topic,
	LastReceivedMsg: devices.SirenMsg{},
	EventChan:       make(chan bool),
}

var Switch = devices.SwitchType{
	Topic:           config.Devices.Switch.Topic,
	LastReceivedMsg: devices.SwitchMsg{},
}

func sensorsHandler(client mqtt.Client, msg mqtt.Message) {
	switch msg.Topic() {
	case DoorSensor.Topic:
		DoorSensor.MqttHandler(msg)
		rules.DoorDef(msg, DoorSensor, &Siren)
	case TemperatureSensor.Topic:
		TemperatureSensor.MqttHandler(msg)
	case Light.Topic:
		Light.MqttHandler(msg)
	case Siren.Topic:
		Siren.MqttHandler(msg)
	case Switch.Topic:
		Switch.MqttHandler(msg)
		rules.SwitchControl(msg, Switch, &Siren, &Light)
	}
}

func setUpSubscriptions(client mqtt.Client) {
	token := client.SubscribeMultiple(
		map[string]byte{
			DoorSensor.Topic:        0,
			TemperatureSensor.Topic: 0,
			Light.Topic:             0,
			Siren.Topic:             0,
			Switch.Topic:            0,
		},
		sensorsHandler,
	)
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Run() {
	setUpSubscriptions(client.MqttClient)
}
