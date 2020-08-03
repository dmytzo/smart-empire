package sensors

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"smart_empire/config"
)

type DoorSensorMsg struct {
	Battery int64 		`json:"battery"`
	Voltage int64 		`json:"voltage"`
	Contact bool 		`json:"contact"`
	Linkquality int64 	`json:"linkquality"`
}

type DoorSensorType struct {
	Name string
	Topic string
	MsgChan chan DoorSensorMsg
}

var DoorSensor = DoorSensorType{
	Name:    "DoorSensor",
	Topic:   config.Cfg.MqttClient.Sensors.Door.Topic,
	MsgChan: make(chan DoorSensorMsg),
}

func (ds DoorSensorType) MqttHandler (client mqtt.Client, msg mqtt.Message) {
	var sensorMsg DoorSensorMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	fmt.Println(sensorMsg.Contact)
	ds.MsgChan <- sensorMsg
}