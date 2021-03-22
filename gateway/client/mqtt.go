package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smart_empire/gateway/config"
	"time"
)

var MqttClient = createMqttClient()

func createMqttClient() mqtt.Client {
	var cfg = config.Clients.Mqtt
	opts := mqtt.NewClientOptions().AddBroker(cfg.BrokerUrl).SetClientID(cfg.ClientId)

	opts.SetKeepAlive(120 * time.Second)
	opts.SetDefaultPublishHandler(func(c mqtt.Client, message mqtt.Message) {})

	var client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func Publish(topic string, payload interface{}) {
	token := MqttClient.Publish(topic, byte(0), false, payload)
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
