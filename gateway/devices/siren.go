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
	Name            string
	Topic           string
	LastReceivedMsg SirenMsg
	AlarmChan       chan bool
}

func (st *SirenType) MqttHandler(msg mqtt.Message) {
	var sensorMsg SirenReceivedMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	msgToSend := SirenMsg{sensorMsg.Value}
	st.LastReceivedMsg = msgToSend
}

func (st *SirenType) GetLatestMsg() SirenMsg {
	return st.LastReceivedMsg
}

func (st *SirenType) GetPublishTopic() string {
	return fmt.Sprintf("%s/set", st.Topic)
}

func (st *SirenType) AutoSwitch() {
	st.Switch(!st.LastReceivedMsg.Value)
}

func (st *SirenType) Switch(option bool) {
	var payload, _ = json.Marshal(map[string]bool{"value": option})
	client.Publish(st.GetPublishTopic(), payload)
	st.LastReceivedMsg.Value = option
}

func (st *SirenType) ActivateAlarm() {
	st.Switch(true)
	st.AlarmChan <- true
}
