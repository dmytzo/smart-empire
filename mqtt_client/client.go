package mqtt_client

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"smart_empire/config"
	"smart_empire/sensors"
	"time"
)

var client mqtt.Client
var cfg = config.Cfg.MqttClient

func defaultHandler(client mqtt.Client, msg mqtt.Message) {}

func sensorsHandler(client mqtt.Client, msg mqtt.Message) {
	switch msg.Topic() {
	case sensors.DoorSensor.Topic:
		sensors.DoorSensor.MqttHandler(msg)
	case sensors.TemperatureSensor.Topic:
		sensors.TemperatureSensor.MqttHandler(msg)
	}
}

func Run() {
	opts := mqtt.NewClientOptions().AddBroker(cfg.BrokerUrl).SetClientID(cfg.ClientId)

	opts.SetKeepAlive(120 * time.Second)
	opts.SetDefaultPublishHandler(defaultHandler)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	setUpSubscriptions()
	fmt.Println("mqtt client is connected!")
}

func setUpSubscriptions() {
	token := client.SubscribeMultiple(
		map[string]byte{
			sensors.DoorSensor.Topic: 0,
			sensors.TemperatureSensor.Topic: 0,
		},
		sensorsHandler,
	)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
