package devices

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/client"
)

type SirenReceivedMsg struct {
	Time  int64 `json:"time"`
	Value bool  `json:"value"`
}

type SirenMsg struct {
	Value bool `json:"value"`
}

type SirenType struct {
	Topic           string
	LastReceivedMsg SirenMsg
	EventChan       chan bool
}

func (d *SirenType) MqttHandler(msg mqtt.Message) {
	var sensorMsg SirenReceivedMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	msgToSend := SirenMsg{sensorMsg.Value}
	d.LastReceivedMsg = msgToSend
}

func (d *SirenType) GetLatestMsg() SirenMsg {
	return d.LastReceivedMsg
}

func (d *SirenType) GetPublishTopic() string {
	return fmt.Sprintf("%s/set", d.Topic)
}

func (d *SirenType) Switch() {
	d.setValue(!d.LastReceivedMsg.Value)
}

func (d *SirenType) setValue(option bool) {
	var payload, _ = json.Marshal(map[string]bool{"value": option})
	client.Publish(d.GetPublishTopic(), payload)
	d.LastReceivedMsg.Value = option
}

func (d *SirenType) ActivateAlarm() {
	d.setValue(true)
	d.EventChan <- true
}
