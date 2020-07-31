package sensors

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
)

var DoorSensorChan chan DoorSensorMsg
var DoorSensorTopic = os.Getenv("DOOR_SENSOR_TOPIC")

type DoorSensorMsg struct {
	Battery int64 		`json:"battery"`
	Voltage int64 		`json:"voltage"`
	Contact bool 		`json:"contact"`
	Linkquality int64 	`json:"linkquality"`
}

var DoorSensorHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var sensorMsg DoorSensorMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	DoorSensorChan <- sensorMsg
}

