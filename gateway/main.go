package gateway

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/client"
	"smart_empire/gateway/config"
	"smart_empire/gateway/devices"
	"smart_empire/gateway/rules"
)



var DoorSensor = devices.DoorSensorType{
	Name:  "DoorSensor",
	Topic: config.Devices.Door.Topic,
	LastReceivedMsg: devices.DoorSensorMsg{},
	DoorEventChan: make(chan devices.DoorSensorMsg),
}

var Light = devices.LightType{
	Name:    "Light",
	Topic:   config.Devices.Light.Topic,
	LastReceivedMsg: devices.LightMsg{},
}

var TemperatureSensor = devices.TemperatureSensorType{
	Name:    "TemperatureSensor",
	Topic:   config.Devices.Temperature.Topic,
	LastReceivedMsg: devices.TemperatureSensorMsg{},
}

var Siren = devices.SirenType{
	Name:    "Siren",
	Topic:   config.Devices.Siren.Topic,
	LastReceivedMsg: devices.SirenMsg{},
	AlarmChan: make(chan bool),
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
	}
}

func setUpSubscriptions(client mqtt.Client) {
	token := client.SubscribeMultiple(
		map[string]byte{
			DoorSensor.Topic:        0,
			TemperatureSensor.Topic: 0,
			Light.Topic: 0,
			Siren.Topic: 0,
		},
		sensorsHandler,
	)
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Run()  {
	setUpSubscriptions(client.MqttClient)
}
