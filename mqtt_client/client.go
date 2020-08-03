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
var mqttClientCfg = config.Cfg.MqttClient

var defaultHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {}

func Run() {
	opts := mqtt.NewClientOptions().AddBroker(mqttClientCfg.BrokerUrl).SetClientID(mqttClientCfg.ClientId)

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
	token := client.Subscribe(sensors.DoorSensor.Topic, 0, sensors.DoorSensor.MqttHandler)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
