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
	Contact     bool  `json:"contact"`
}

type DoorSensorType struct {
	Name    string
	Topic   string
	LastReceivedMsg DoorSensorMsg
	DoorEventChan chan DoorSensorMsg
}

func (ds *DoorSensorType) ParseMsg(msg mqtt.Message) DoorSensorMsg {
	var sensorMsg doorSensorMsgReceived
	json.Unmarshal(msg.Payload(), &sensorMsg)
	return DoorSensorMsg{sensorMsg.Contact}
}

func (ds *DoorSensorType) MqttHandler(msg mqtt.Message) {
	msgToSend := ds.ParseMsg(msg)
	if ds.LastReceivedMsg == msgToSend {
		return
	}
	ds.LastReceivedMsg = msgToSend
	ds.DoorEventChan <- msgToSend
}

func (ds *DoorSensorType) GetLatestMsg() DoorSensorMsg {
	return ds.LastReceivedMsg
}
