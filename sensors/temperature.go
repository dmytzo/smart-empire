package sensors

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"smart_empire/config"
)

type TemperatureSensorMsg struct {
	Battery int64 		`json:"battery"`
	Voltage int64 		`json:"voltage"`
	Temperature float64 `json:"temperature"`
	Humidity float64 	`json:"humidity"`
	Pressure float64 	`json:"pressure"`
	Linkquality int64 	`json:"linkquality"`
}

type TemperatureSensorType struct {
	Name string
	Topic string
	MsgChan chan TemperatureSensorMsg
}

var TemperatureSensor = TemperatureSensorType{
	Name:    "TemperatureSensor",
	Topic:   config.Cfg.MqttClient.Sensors.Temperature.Topic,
	MsgChan: make(chan TemperatureSensorMsg),
}

func (ts TemperatureSensorType) MqttHandler (msg mqtt.Message) {
	var sensorMsg TemperatureSensorMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	ts.MsgChan <- sensorMsg
}
