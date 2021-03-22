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
	Topic string
	LastReceivedMsg LightMsg
}

func (d *LightType) MqttHandler (msg mqtt.Message) {
	var sensorMsg LightReceivedMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	msgToSend := LightMsg{sensorMsg.State}
	d.LastReceivedMsg = msgToSend
}

func (d *LightType) GetLatestMsg() LightMsg {
	return d.LastReceivedMsg
}

func (d *LightType) GetPublishTopic() string {
	return fmt.Sprintf("%s/set", d.Topic)
}

func (d *LightType) Switch() {
	option := "OFF"
	if d.LastReceivedMsg.State == option {
		option = "ON"
	}
	var payload, _ = json.Marshal(map[string]string{"state": option})
	client.Publish(d.GetPublishTopic(), payload)
	d.LastReceivedMsg.State = option
}