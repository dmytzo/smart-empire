package devices

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/client"
)

type LightReceivedMsg struct {
	State       string `json:"state"`
	Brightness  int64 `json:"brightness"`
	ColorTemp   int64 `json:"color_temp"`
	Linkquality int64 `json:"linkquality"`
}

type LightMsg struct {
	State       string  `json:"state"`
}

type LightType struct {
	Name string
	Topic string
	LastReceivedMsg LightMsg
}

func (ts *LightType) MqttHandler (msg mqtt.Message) {
	var sensorMsg LightReceivedMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	msgToSend := LightMsg{sensorMsg.State}
	ts.LastReceivedMsg = msgToSend
}

func (ts *LightType) GetLatestMsg() LightMsg {
	return ts.LastReceivedMsg
}

func (ts *LightType) GetPublishTopic() string {
	return fmt.Sprintf("%s/set", ts.Topic)
}

func (ts *LightType) Switch() {
	option := "OFF"
	if ts.LastReceivedMsg.State == option {
		option = "ON"
	}
	var payload, _ = json.Marshal(map[string]string{"state": option})
	client.Publish(ts.GetPublishTopic(), payload)
	ts.LastReceivedMsg.State = option
}