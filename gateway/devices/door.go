package devices

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
)

type doorSensorMsgReceived struct {
	Battery     int64 `json:"battery"`
	Voltage     int64 `json:"voltage"`
	Contact     bool  `json:"contact"`
	Linkquality int64 `json:"linkquality"`
}

type DoorSensorMsg struct {
	Contact bool `json:"contact"`
}

type DoorSensorType struct {
	Topic           string
	LastReceivedMsg DoorSensorMsg
	EventChan       chan DoorSensorMsg
}

func (d *DoorSensorType) ParseMsg(msg mqtt.Message) DoorSensorMsg {
	var sensorMsg doorSensorMsgReceived
	json.Unmarshal(msg.Payload(), &sensorMsg)
	return DoorSensorMsg{sensorMsg.Contact}
}

func (d *DoorSensorType) MqttHandler(msg mqtt.Message) {
	msgToSend := d.ParseMsg(msg)
	if d.LastReceivedMsg == msgToSend {
		return
	}
	d.LastReceivedMsg = msgToSend
	d.EventChan <- msgToSend
}

func (d *DoorSensorType) GetLatestMsg() DoorSensorMsg {
	return d.LastReceivedMsg
}
