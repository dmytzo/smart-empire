package devices

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
)

const (
	Btn1Click   = "button_1_click"
	Btn1Hold    = "button_1_hold"
	Btn1Release = "button_1_release"

	Btn2Click   = "button_2_click"
	Btn2Hold    = "button_2_hold"
	Btn2Release = "button_2_release"

	Btn3Click   = "button_3_click"
	Btn3Hold    = "button_3_hold"
	Btn3Release = "button_3_release"

	Btn4Click   = "button_4_click"
	Btn4Hold    = "button_4_hold"
	Btn4Release = "button_4_release"
)

type switchMsgReceived struct {
	Action      string `json:"action"`
	Linkquality int64  `json:"linkquality"`
}

type SwitchMsg struct {
	Action string `json:"action"`
}

type SwitchType struct {
	Topic           string
	LastReceivedMsg SwitchMsg
	SwitchEventChan   chan SwitchMsg
}

func (d *SwitchType) ParseMsg(msg mqtt.Message) SwitchMsg {
	var switchMsg switchMsgReceived
	json.Unmarshal(msg.Payload(), &switchMsg)
	return SwitchMsg{switchMsg.Action}
}

func (d *SwitchType) MqttHandler(msg mqtt.Message) {
	msgToSend := d.ParseMsg(msg)
	d.LastReceivedMsg = msgToSend
}

func (d *SwitchType) GetLatestMsg() SwitchMsg {
	return d.LastReceivedMsg
}
