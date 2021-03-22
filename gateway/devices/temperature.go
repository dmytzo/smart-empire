package devices

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/config"
	"time"
)

var settings = config.Devices.Temperature.Settings

const (
	EXCELLENT = "EXCELLENT"
	NORMAL    = "NORMAL"
	BAD       = "BAD"
)

type TemperatureSensorReceivedMsg struct {
	Battery     int64   `json:"battery"`
	Voltage     int64   `json:"voltage"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Pressure    float64 `json:"pressure"`
	Linkquality int64   `json:"linkquality"`
}

type TemperatureSensorMsg struct {
	Temperature          float64 `json:"temperature"`
	Humidity             float64 `json:"humidity"`
	Pressure             float64 `json:"pressure"`
	TemperatureIndicator string  `json:"temperature_indicator"`
	HumidityIndicator    string  `json:"humidity_indicator"`
}

type TemperatureSensorType struct {
	Topic           string
	LastReceivedMsg TemperatureSensorMsg
}

func (d *TemperatureSensorType) MqttHandler(msg mqtt.Message) {
	var sensorMsg TemperatureSensorReceivedMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
	msgToSend := TemperatureSensorMsg{
		sensorMsg.Temperature,
		sensorMsg.Humidity,
		sensorMsg.Pressure,
		getTemperatureStatus(sensorMsg.Temperature),
		getHumidityStatus(sensorMsg.Humidity),
	}
	d.LastReceivedMsg = msgToSend
}

func (d *TemperatureSensorType) GetLatestMsg() TemperatureSensorMsg {
	return d.LastReceivedMsg
}

func getTemperatureStatus(temperature float64) string {
	currentMonth := getCurrentMonth()
	monthsCfg := settings.Temperature.HotMonths
	coldMonthsCfg := settings.Temperature.ColdsMonths
	for _, month := range coldMonthsCfg.Months {
		if month == currentMonth {
			monthsCfg = coldMonthsCfg
		}
	}
	normalDelta := settings.Temperature.NormalDelta
	if temperature >= monthsCfg.From && temperature <= monthsCfg.To {
		return EXCELLENT
	}
	normalFrom := monthsCfg.From - normalDelta
	normalTo := monthsCfg.To + normalDelta
	if temperature >= normalFrom && temperature <= normalTo {
		return NORMAL
	}
	return BAD
}

func getHumidityStatus(humidity float64) string {
	normalDelta := settings.Humidity.NormalDelta
	if humidity >= settings.Humidity.From && humidity <= settings.Humidity.To {
		return EXCELLENT
	}
	normalFrom := settings.Humidity.From - normalDelta
	normalTo := settings.Humidity.To + normalDelta
	if humidity >= normalFrom && humidity <= normalTo {
		return NORMAL
	}
	return BAD
}

func getCurrentMonth() int64 {
	month := time.Now().Month()
	return int64(month)
}
